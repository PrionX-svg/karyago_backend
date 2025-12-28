package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"hris_backend/internal/request"
	"hris_backend/internal/response"
	"hris_backend/pkg"
)

type TimeRange string

const (
	ThisWeek  TimeRange = "this_week"
	LastWeek  TimeRange = "last_week"
	ThisMonth TimeRange = "this_month"
	LastMonth TimeRange = "last_month"
	ThisYear  TimeRange = "this_year"
	LastYear  TimeRange = "last_year"
)

type ChatbotServices interface {
	Ask(req request.ChatbotRequest, userID uint, role string) (response.ChatbotResponse, error)
}

type chatbotServices struct {
	apiKey    string
	client    *http.Client
	knowledge string
	attSvc    AttendanceService //inject service attendance
}

// Hanya role ini yang boleh lihat data sensitif (rekap karyawan lain)
func (s *chatbotServices) canAccessSensitiveData(role string) bool {
	switch strings.ToLower(role) {
	case "admin", "assistant", "owner":
		return true
	default:
		return false
	}
}

func NewChatbotService(apiKey string, companyFilePath string, attSvc AttendanceService) ChatbotServices {
	knowledge := ""
	if companyFilePath != "" {
		b, err := os.ReadFile(companyFilePath)
		if err == nil {
			knowledge = string(b)
		} else {
			fmt.Printf("⚠️ gagal baca knowledge file: %v\n", err)
		}
	}

	return &chatbotServices{
		apiKey:    apiKey,
		client:    &http.Client{},
		knowledge: knowledge,
		attSvc:    attSvc,
	}
}

func (s *chatbotServices) Ask(req request.ChatbotRequest, userID uint, role string) (response.ChatbotResponse, error) {
	if err := pkg.Validate.Struct(req); err != nil {
		return response.ChatbotResponse{}, err
	}

	model := req.Model
	if model == "" {
		// model = "x-ai/grok-code-fast-1"
		model = "nvidia/nemotron-3-nano-30b-a3b:free"
	}

	//Tambahan: deteksi intent & ambil data attendance
	msg := strings.ToLower(req.Message)
	var preAnswer string

	// --- intent & gating ---
	if strings.Contains(msg, "lembur") {
		if !s.canAccessSensitiveData(role) {
			return response.ChatbotResponse{
				Reply: "Maaf, Anda tidak memiliki izin untuk melihat data lembur karyawan lain.",
			}, nil
		}

		// Default: minggu ini
		timeRange := ThisWeek

		// Parsing rentang waktu
		switch {
		case strings.Contains(msg, "minggu lalu"),
			strings.Contains(msg, "pekan lalu"):
			timeRange = LastWeek

		case strings.Contains(msg, "bulan lalu"),
			strings.Contains(msg, "bulan kemarin"):
			timeRange = LastMonth

		case strings.Contains(msg, "bulan ini"):
			timeRange = ThisMonth

		case strings.Contains(msg, "tahun lalu"),
			strings.Contains(msg, "tahun kemarin"):
			timeRange = LastYear

		case strings.Contains(msg, "tahun ini"):
			timeRange = ThisYear
		}

		// Ambil data lembur sesuai range
		preAnswer = s.getOvertimeSummary(timeRange)

	} else if strings.Contains(msg, "belum clock out") ||
		strings.Contains(msg, "belum absen keluar") ||
		strings.Contains(msg, "belum pulang") {

		if s.canAccessSensitiveData(role) {
			preAnswer = s.getUnclockedOutSummary()
		} else {
			return response.ChatbotResponse{
				Reply: "Maaf, Anda tidak memiliki izin untuk melihat siapa yang belum clock out. Silakan hubungi atasan atau lihat dashboard sesuai peran Anda.",
			}, nil
		}

	} else if strings.Contains(msg, "belum clock in") ||
		strings.Contains(msg, "belum absen masuk") ||
		strings.Contains(msg, "belum hadir") ||
		strings.Contains(msg, "belum presensi") {

		if s.canAccessSensitiveData(role) {
			preAnswer = s.getUnclockedInSummary()
		} else {
			return response.ChatbotResponse{
				Reply: "Maaf, Anda tidak memiliki izin untuk melihat siapa yang belum clock in.",
			}, nil
		}
	}

	messages := []map[string]string{}

	fmt.Println("🔎 preAnswer:", preAnswer)

	if s.knowledge != "" {
		systemContent := "Kamu adalah asisten perusahaan. " +
			"Gunakan informasi berikut tentang perusahaan untuk menjawab pertanyaan: \n\n" +
			s.knowledge

		// kalau ada data attendance tambahan, sisipkan ke konteks
		if preAnswer != "" {
			systemContent += "\n\nBerikut data attendance terkini:\n" + preAnswer
		}

		systemContent += "\n\nJika pertanyaan tidak relevan dengan informasi ini, jawab dengan jawaban umum."

		messages = append(messages, map[string]string{
			"role":    "system",
			"content": systemContent,
		})
	}

	messages = append(messages, map[string]string{
		"role":    "user",
		"content": req.Message,
	})

	payload := map[string]interface{}{
		"model":    model,
		"messages": messages,
	}

	body, _ := json.Marshal(payload)

	httpReq, err := http.NewRequest(
		"POST",
		"https://openrouter.ai/api/v1/chat/completions",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return response.ChatbotResponse{}, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return response.ChatbotResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(resp.Body)
		return response.ChatbotResponse{}, errors.New(string(b))
	}

	var result struct {
		Choices []struct {
			Message struct {
				Role    string `json:"role"`
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return response.ChatbotResponse{}, err
	}

	if len(result.Choices) == 0 {
		return response.ChatbotResponse{}, errors.New("no reply from chatbot")
	}

	return response.ChatbotResponse{
		Reply: result.Choices[0].Message.Content,
	}, nil

}

func resolveTimeRange(r TimeRange) (time.Time, time.Time) {
	now := time.Now()
	loc := now.Location()

	switch r {

	case ThisWeek:
		from := now.AddDate(0, 0, -7)
		return from, now

	case LastWeek:
		to := now.AddDate(0, 0, -7)
		from := to.AddDate(0, 0, -7)
		return from, to

	case ThisMonth:
		from := time.Date(
			now.Year(), now.Month(), 1,
			0, 0, 0, 0, loc,
		)
		return from, now

	case LastMonth:
		firstThisMonth := time.Date(
			now.Year(), now.Month(), 1,
			0, 0, 0, 0, loc,
		)
		to := firstThisMonth.Add(-time.Second)
		from := time.Date(
			to.Year(), to.Month(), 1,
			0, 0, 0, 0, loc,
		)
		return from, to

	case ThisYear:
		from := time.Date(
			now.Year(), time.January, 1,
			0, 0, 0, 0, loc,
		)
		return from, now

	case LastYear:
		from := time.Date(
			now.Year()-1, time.January, 1,
			0, 0, 0, 0, loc,
		)
		to := time.Date(
			now.Year()-1, time.December, 31,
			23, 59, 59, 0, loc,
		)
		return from, to
	}

	// fallback aman
	return now.AddDate(0, 0, -7), now
}

// Ambil ringkasan lembur minggu ini
func (s *chatbotServices) getOvertimeSummary(r TimeRange) string {
	from, to := resolveTimeRange(r)

	companyUUID, err := s.attSvc.GetDefaultCompanyUUID()
	if err != nil {
		return "Tidak dapat menentukan perusahaan aktif."
	}

	list, err := s.attSvc.ListAllEmployeeAttendance(companyUUID, from, to)

	var sb strings.Builder

	//Judul dinamis sesuai rentang waktu
	switch r {
	case ThisWeek:
		sb.WriteString("Daftar karyawan lembur minggu ini:\n")
	case LastWeek:
		sb.WriteString("Daftar karyawan lembur minggu lalu:\n")
	case ThisMonth:
		sb.WriteString("Daftar karyawan lembur bulan ini:\n")
	case LastMonth:
		sb.WriteString("Daftar karyawan lembur bulan lalu:\n")
	case ThisYear:
		sb.WriteString("Daftar karyawan lembur tahun ini:\n")
	case LastYear:
		sb.WriteString("Daftar karyawan lembur tahun lalu:\n")
	default:
		sb.WriteString("Daftar karyawan lembur:\n")
	}

	count := 0

	for _, a := range list {

		var overtime float64
		var isOT bool

		//Kalau overtime sudah dihitung & disimpan
		if a.IsOvertime && a.OvertimeHours != nil && *a.OvertimeHours > 0 {
			overtime = *a.OvertimeHours
			isOT = true

			//FALLBACK untuk data lama
		} else if a.TotalWorkHours != nil && *a.TotalWorkHours > 8 {
			overtime = *a.TotalWorkHours - 8
			isOT = true
		}

		if !isOT {
			continue
		}

		name := "-"
		if a.Employee.User != nil {
			name = fmt.Sprintf(
				"%s %s",
				a.Employee.User.FirstName,
				a.Employee.User.LastName,
			)
		}

		sb.WriteString(fmt.Sprintf(
			"- %s (%s): %.1f jam lembur\n",
			name,
			a.WorkDate.Format("02 Jan 2006"),
			overtime,
		))

		count++
	}

	if count == 0 {
		return "Tidak ada karyawan yang melakukan lembur pada periode tersebut."
	}

	fmt.Println("📦 Data lembur dikirim ke OpenRouter:")
	fmt.Println(sb.String())

	return sb.String()
}

// Belum Absen Masuk (Clock In)
func (s *chatbotServices) getUnclockedInSummary() string {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	// Ambil range hari ini (00:00–23:59 WIB)
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, loc)

	companyUUID, err := s.attSvc.GetDefaultCompanyUUID()
	if err != nil {
		fmt.Println("❌ Gagal ambil company UUID:", err)
		return "Tidak dapat menentukan perusahaan aktif."
	}

	// Ambil data attendance seluruh karyawan di hari ini
	list, err := s.attSvc.ListAllEmployeeAttendance(companyUUID, startOfDay, endOfDay)
	if err != nil {
		fmt.Println("❌ Error getUnclockedInSummary:", err)
		return "Terjadi kesalahan saat mengambil data absensi."
	}
	if len(list) == 0 {
		return "Sepertinya belum ada karyawan yang melakukan clock in hari ini ⏰"
	}

	var sb strings.Builder
	sb.WriteString("Karyawan yang belum melakukan clock in hari ini:\n")

	count := 0
	for _, a := range list {
		if a.ClockInAt == nil {
			name := "-"
			if a.Employee.User != nil {
				name = fmt.Sprintf("%s %s", a.Employee.User.FirstName, a.Employee.User.LastName)
			}
			sb.WriteString(fmt.Sprintf("- %s belum melakukan clock in.\n", name))
			count++
		}
	}

	if count == 0 {
		return "Semua karyawan sudah melakukan clock in hari ini ✅"
	}

	return sb.String()
}

// Belum Clock Out
func (s *chatbotServices) getUnclockedOutSummary() string {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, loc)

	companyUUID, err := s.attSvc.GetDefaultCompanyUUID()
	if err != nil {
		fmt.Println("❌ Gagal ambil company UUID:", err)
		return "Tidak dapat menentukan perusahaan aktif."
	}

	list, err := s.attSvc.ListAllEmployeeAttendance(companyUUID, startOfDay, endOfDay)
	if err != nil {
		fmt.Println("❌ Error getUnclockedOutSummary:", err)
		return "Terjadi kesalahan saat mengambil data absensi."
	}
	if len(list) == 0 {
		return "Belum ada data absensi hari ini."
	}

	var sb strings.Builder
	sb.WriteString("Karyawan yang belum melakukan clock out hari ini:\n")

	count := 0
	for _, a := range list {
		if a.ClockInAt != nil && a.ClockOutAt == nil {
			name := "-"
			if a.Employee.User != nil {
				name = fmt.Sprintf("%s %s", a.Employee.User.FirstName, a.Employee.User.LastName)
			}
			inTime := a.ClockInAt.Format("15:04")
			sb.WriteString(fmt.Sprintf("- %s (Clock In: %s)\n", name, inTime))
			count++
		}
	}

	if count == 0 {
		return "Semua karyawan sudah melakukan clock out hari ini 🕓"
	}

	return sb.String()
}
