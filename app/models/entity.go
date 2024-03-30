package models

type Entity struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Nif  int    `json:"nif"`
	Name string `json:"name"`

	LocalityId uint     `json:"locality_id,omitempty"`
	Locality   Locality `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"locality,omitempty"`
}
