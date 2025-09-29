package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"hris_backend/internal/services"
	"hris_backend/pkg"
)

type AttendanceHandler interface {
	// actions
	ClockIn(c *fiber.Ctx) error
	ClockOut(c *fiber.Ctx) error
	ToggleHomeOffice(c *fiber.Ctx) error
	SaveNotes(c *fiber.Ctx) error

	// queries
	GetByDate(c *fiber.Ctx) error
	ListRange(c *fiber.Ctx) error
}

type attendanceHandler struct {
	svc services.AttendanceService
}

func NewAttendanceHandler(svc services.AttendanceService) AttendanceHandler {
	return &attendanceHandler{svc}
}

func parseDate(s string) (time.Time, error) {
	// format: YYYY-MM-DD (mengikuti WorkDate DATE)
	return time.Parse("2006-01-02", s)
}

func (h *attendanceHandler) ClockIn(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok || userID == 0 {
		return pkg.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	type Req struct {
		CompanyUUID *string `json:"company_uuid,omitempty"` // opsional jika multi-company
		WorkDate    string  `json:"work_date"`              // "2025-09-26" (local date)
		At          *string `json:"at,omitempty"`           // optional override datetime ISO; default now
	}
	var req Req
	if err := c.BodyParser(&req); err != nil || req.WorkDate == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "invalid payload")
	}

	workDate, err := parseDate(req.WorkDate)
	if err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "invalid work_date")
	}

	at := time.Now().UTC()
	if req.At != nil && *req.At != "" {
		if atParsed, e := time.Parse(time.RFC3339, *req.At); e == nil {
			at = atParsed
		}
	}

	att, err := h.svc.ClockIn(userID, req.CompanyUUID, workDate, at)
	if err != nil {
		return pkg.Error(c, 400, err.Error())
	}

	return pkg.Success(c, att, "clock-in success")
}

func (h *attendanceHandler) ClockOut(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok || userID == 0 {
		return pkg.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	type Req struct {
		CompanyUUID *string `json:"company_uuid,omitempty"`
		WorkDate    string  `json:"work_date"`
		At          *string `json:"at,omitempty"`
	}
	var req Req
	if err := c.BodyParser(&req); err != nil || req.WorkDate == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "invalid payload")
	}
	workDate, err := time.Parse("2006-01-02", req.WorkDate)
	if err != nil {
		return pkg.Error(c, 400, "invalid work_date")
	}

	at := time.Now().UTC()
	if req.At != nil && *req.At != "" {
		if t, e := time.Parse(time.RFC3339, *req.At); e == nil {
			at = t
		}
	}

	att, err := h.svc.ClockOut(userID, req.CompanyUUID, workDate, at)

	if err != nil {
		return pkg.Error(c, 400, err.Error())
	}
	return pkg.Success(c, att, "clock-out success")
}

func (h *attendanceHandler) ToggleHomeOffice(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok || userID == 0 {
		return pkg.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	type Req struct {
		CompanyUUID *string `json:"company_uuid,omitempty"`
		WorkDate    string  `json:"work_date"`      // tanggal kerja
		IsHome      bool    `json:"is_home_office"` // true=Home Office, false=In Office
	}
	var req Req
	if err := c.BodyParser(&req); err != nil || req.WorkDate == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "invalid payload")
	}
	workDate, err := parseDate(req.WorkDate)
	if err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "invalid work_date")
	}

	att, err := h.svc.ToggleHomeOffice(userID, req.CompanyUUID, workDate, req.IsHome)
	if err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return pkg.Success(c, att, "home/office updated")
}

func (h *attendanceHandler) SaveNotes(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok || userID == 0 {
		return pkg.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	type Req struct {
		CompanyUUID *string `json:"company_uuid,omitempty"`
		WorkDate    string  `json:"work_date"`
		Notes       *string `json:"notes"` // boleh kosong => null
	}
	var req Req
	if err := c.BodyParser(&req); err != nil || req.WorkDate == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "invalid payload")
	}
	workDate, err := parseDate(req.WorkDate)
	if err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "invalid work_date")
	}

	att, err := h.svc.SaveNotes(userID, req.CompanyUUID, workDate, req.Notes)
	if err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return pkg.Success(c, att, "notes saved")
}

/* ===================== Queries ===================== */

func (h *attendanceHandler) GetByDate(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok || userID == 0 {
		return pkg.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	workDateStr := c.Query("work_date")
	if workDateStr == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "work_date is required")
	}
	workDate, err := parseDate(workDateStr)
	if err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "invalid work_date")
	}

	var companyUUID *string
	if v := c.Query("company_uuid"); v != "" {
		companyUUID = &v
	}

	att, err := h.svc.GetByDate(userID, companyUUID, workDate)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return pkg.Success(c, att, "ok")
}

func (h *attendanceHandler) ListRange(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok || userID == 0 {
		return pkg.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	fromStr := c.Query("from")
	toStr := c.Query("to")
	if fromStr == "" || toStr == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "from and to are required (YYYY-MM-DD)")
	}

	from, err1 := parseDate(fromStr)
	to, err2 := parseDate(toStr)
	if err1 != nil || err2 != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "invalid date format")
	}

	var companyUUID *string
	if v := c.Query("company_uuid"); v != "" {
		companyUUID = &v
	}

	list, err := h.svc.ListRange(userID, companyUUID, from, to)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return pkg.Success(c, list, "ok")
}
