package models
import (
	"time"
)

type AttendanceHistory struct {
    ID           uint       `json:"id" gorm:"primaryKey;autoIncrement"`
    AttendanceID uint       `json:"attendance_id" gorm:"not null;index"`
    Action       string     `json:"action" gorm:"type:enum('CLOCK_IN','CLOCK_OUT','TOGGLE_HOME','NOTES_UPDATE','EDIT_APPROVED','EDIT_REJECTED','ADMIN_UPDATE');not null"`
    OldValueJSON *string    `json:"old_value_json" gorm:"type:json"`
    NewValueJSON *string    `json:"new_value_json" gorm:"type:json"`
    ActedBy      *uint      `json:"acted_by"`
    ActedAt      time.Time  `json:"acted_at" gorm:"autoCreateTime"`
}
