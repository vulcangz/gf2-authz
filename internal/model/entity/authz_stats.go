package entity

import "time"

// Stats is the golang structure for table authz_stats.
type Stats struct {
	ID                  string    `json:"id"                  gorm:"primarykey"`               //
	Date                time.Time `json:"date"                gorm:"date"`                     //
	ChecksAllowedNumber int64     `json:"checks_allowed_number" gorm:"checks_allowed_number" ` //
	ChecksDeniedNumber  int64     `json:"checks_denied_number"  gorm:"checks_denied_number"  ` //
}

func (Stats) TableName() string {
	return "authz_stats"
}
