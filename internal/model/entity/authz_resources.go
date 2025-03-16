package entity

import "time"

// Resource is the golang structure for table authz_resources.
type Resource struct {
	ID         string     `json:"id"        gorm:"primarykey"`  //
	Kind       string     `json:"kind"      gorm:"kind"`        //
	Value      string     `json:"value"     gorm:"value"`       //
	IsLocked   bool       `json:"is_locked"  gorm:"is_locked"`  //
	CreatedAt  time.Time  `json:"created_at" gorm:"created_at"` //
	UpdatedAt  time.Time  `json:"updated_at" gorm:"updated_at"` //
	Attributes Attributes `json:"attributes,omitempty" gorm:"many2many:authz_resources_attributes;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Resource) TableName() string {
	return "authz_resources"
}

type ResourceQueryOption func(*ResourceQueryOptions)

type ResourceQueryOptions struct {
	ResourceIDs []string
}

func WithResourceIDs(resourceIDs []string) ResourceQueryOption {
	return func(o *ResourceQueryOptions) {
		o.ResourceIDs = resourceIDs
	}
}

type ResourceMatchingAttribute struct {
	ResourceKind   string
	ResourceValue  string
	AttributeValue string
}
