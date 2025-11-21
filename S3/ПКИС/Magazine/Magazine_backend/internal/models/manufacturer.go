package models

import "gorm.io/datatypes"

// Manufacturer — компания-производитель товара
type Manufacturer struct {
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"` // ID производителя / manufacturer ID

	Name        string         `gorm:"type:text;not null" json:"name"` // название компании / company name
	Country     *string        `gorm:"type:text" json:"country"`       // страна производителя / country
	ContactInfo datatypes.JSON `gorm:"type:jsonb" json:"contact_info"` // контакты (телефон, сайт) / contact info

	Products []Product `gorm:"foreignKey:ManufacturerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // список товаров / products by manufacturer
}

type CreateManufacturer struct {
	Name        string         `json:"name"`
	Country     *string        `json:"country"`
	ContactInfo datatypes.JSON `json:"contact_info"`
}

type UpdateManufacturer struct {
	Name        *string         `json:"name"`
	Country     *string         `json:"country"`
	ContactInfo *datatypes.JSON `json:"contact_info"`
}
