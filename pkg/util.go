package pkg

import (
	"os"
	"strconv"
	"time"
)

func DerefString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func AppLocation() *time.Location {
	tz := os.Getenv("DB_TIMEZONE")
	if tz == "" {
		tz = "Asia/Jakarta" // fallback
	}
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return time.UTC
	}
	return loc
}


// AttendanceCutoffHours mengembalikan jumlah jam untuk batas "past due".
// ENV: ATTENDANCE_EDIT_AFTER_HOURS (default 24).
func AttendanceCutoffHours() int {
	def := 24
	if v := os.Getenv("ATTENDANCE_EDIT_AFTER_HOURS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			return n
		}
	}
	return def
}
