package entity

import "time"

// Client is the golang structure for table authz_clients.
type Client struct {
	ID        string    `json:"client_id"        gorm:"id"         `
	Secret    string    `json:"client_secret"    gorm:"secret"     `
	Name      string    `json:"name"      gorm:"name"       `
	Domain    string    `json:"domain"    gorm:"domain"     `
	Data      string    `json:"data"      gorm:"data"       `
	CreatedAt time.Time `json:"created_at" gorm:"created_at" `
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at" `
}

func (Client) TableName() string {
	return "authz_clients"
}
