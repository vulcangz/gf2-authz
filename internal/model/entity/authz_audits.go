package entity

import "time"

// Audit is the golang structure for table authz_audit.
type Audit struct {
	ID            int64     `json:"id"             gorm:"id"             ` //
	Date          time.Time `json:"date"           gorm:"date"           ` //
	Principal     string    `json:"principal"      gorm:"principal"      ` //
	ResourceKind  string    `json:"resource_kind"  gorm:"resource_kind"  ` //
	ResourceValue string    `json:"resource_value" gorm:"resource_value" ` //
	Action        string    `json:"action"         gorm:"action"         ` //
	IsAllowed     int       `json:"is_allowed"     gorm:"is_allowed"     ` //
	PolicyId      string    `json:"policy_id"      gorm:"policy_id"      ` //
}

func (Audit) TableName() string {
	return "authz_audits"
}
