package entity

import (
	"time"
)

// Action is the golang structure for table authz_actions.
type Action struct {
	ID        string    `json:"id"         gorm:"id"         ` //
	CreatedAt time.Time `json:"created_at" gorm:"created_at" ` //
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at" ` //
}

func (Action) TableName() string {
	return "authz_actions"
}
