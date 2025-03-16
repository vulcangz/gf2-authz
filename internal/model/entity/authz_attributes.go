package entity

type Attributes []*Attribute

func (a Attributes) GetAttribute(key string) string {
	for _, attribute := range a {
		if attribute.Key == key {
			return attribute.Value
		}
	}

	return ""
}

// Attribute is the golang structure for table authz_attributes.
type Attribute struct {
	ID    int64  `json:"-" gorm:"id"`
	Key   string `json:"key" gorm:"column:key_name"`
	Value string `json:"value" gorm:"value"`
}

func (Attribute) TableName() string {
	return "authz_attributes"
}
