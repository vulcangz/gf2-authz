package entity

import "time"

// Role is the golang structure for table authz_roles.
type Role struct {
	ID        string    `json:"id"        gorm:"primarykey"`  //
	CreatedAt time.Time `json:"created_at" gorm:"created_at"` //
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"` //

	Policies   []*Policy    `json:"policies,omitempty" gorm:"many2many:authz_roles_policies;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Principals []*Principal `json:"-" gorm:"many2many:authz_principals_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Role) TableName() string {
	return "authz_roles"
}
