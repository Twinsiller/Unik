package models

import "gorm.io/datatypes"

// Supplier — поставщик, у которого закупаем товары
type Supplier struct {
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"` // ID поставщика / supplier ID

	Name        string         `gorm:"type:text;not null" json:"name"` // название компании / supplier name
	ContactInfo datatypes.JSON `gorm:"type:jsonb" json:"contact_info"` // контакты / contact info
	Preferred   bool           `gorm:"default:false" json:"preferred"` // пометка «основной поставщик» / preferred supplier

	Purchases []Purchase `gorm:"foreignKey:SupplierID;references:ID"` // закупки у этого поставщика / purchases
}
