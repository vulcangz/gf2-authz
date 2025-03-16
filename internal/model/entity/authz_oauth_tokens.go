package entity

import "time"

// Token is the golang structure for table authz_oauth_tokens.
type Token struct {
	ID        uint64    `json:"id"        gorm:"primarykey"`
	Code      string    `json:"code"      gorm:"type:varchar(512)"`
	Access    string    `json:"access"    gorm:"type:varchar(512)"`
	Refresh   string    `json:"refresh"   gorm:"type:varchar(512)"`
	Data      string    `json:"data"      gorm:"type:text"`
	ExpiredAt int64     `json:"expired_at" gorm:"expired_at" `
	CreatedAt time.Time `json:"created_at" gorm:"created_at" `
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at" `
}

func (Token) TableName() string {
	return "authz_oauth_tokens"
}
