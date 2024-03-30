package models

type Locality struct {
	ID       uint     `gorm:"primaryKey" json:"id"`
	Name     string   `json:"name"`
	Entities []Entity `gorm:"foreignKey:LocalityId;references:ID" json:",omitempty"`
}
