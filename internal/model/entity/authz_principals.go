package entity

import (
	"fmt"
	"time"
)

// Principal is the golang structure for table authz_principals.
type Principal struct {
	ID        string    `json:"id"        gorm:"primarykey"`  //
	IsLocked  bool      `json:"is_locked"  gorm:"is_locked"`  //
	CreatedAt time.Time `json:"created_at" gorm:"created_at"` //
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"` //

	Roles      []*Role    `json:"roles,omitempty" gorm:"many2many:authz_principals_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Attributes Attributes `json:"attributes,omitempty" gorm:"many2many:authz_principals_attributes;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Principal) TableName() string {
	return "authz_principals"
}

func ClientPrincipal(identifier string) string {
	return fmt.Sprintf("%s-sa-%s", appName, identifier)
}

func UserPrincipal(identifier string) string {
	return fmt.Sprintf("%s-user-%s", appName, identifier)
}

type PrincipalMatchingAttribute struct {
	PrincipalID    string
	AttributeValue string
}
