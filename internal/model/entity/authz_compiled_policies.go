package entity

import "time"

// CompiledPolicy is the golang structure for table authz_compiled_policies.
type CompiledPolicy struct {
	PolicyID      string    `json:"policy_id"      gorm:"index"`      //
	PrincipalID   string    `json:"principal_id"   gorm:"index"`      //
	ResourceKind  string    `json:"resource_kind"  gorm:"index"`      //
	ResourceValue string    `json:"resource_value" gorm:"index"`      //
	ActionID      string    `json:"action_id"      gorm:"index"`      //
	Version       int64     `json:"version"       gorm:"index"`       //
	CreatedAt     time.Time `json:"created_at"     gorm:"created_at"` //
	UpdatedAt     time.Time `json:"updated_at"     gorm:"updated_at"` //
}

func (CompiledPolicy) TableName() string {
	return "authz_compiled_policies"
}
