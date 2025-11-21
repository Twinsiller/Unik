package models

import "time"

// PurchaseItem — конкретный товар в закупке
type PurchaseItem struct {
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"` // ID позиции / purchase item ID

	PurchaseID int64    `gorm:"index;not null" json:"purchase_id"`   // ID закупки / purchase ID
	Purchase   Purchase `gorm:"foreignKey:PurchaseID;references:ID"` // связь / purchase

	ProductID int64   `gorm:"index;not null" json:"product_id"`   // ID товара / product ID
	Product   Product `gorm:"foreignKey:ProductID;references:ID"` // связь / product

	Quantity      float64    `gorm:"type:numeric(14,3);not null" json:"quantity"`       // количество / quantity
	PurchasePrice float64    `gorm:"type:numeric(12,4);not null" json:"purchase_price"` // цена закупки / purchase price
	ExpiryDate    *time.Time `gorm:"type:date" json:"expiry_date"`                      // срок годности / expiry date
	BatchNumber   *string    `gorm:"type:text" json:"batch_number"`                     // номер партии / batch number
}

type CreatePurchaseItem struct {
	PurchaseID int64 `json:"purchase_id"` // ID закупки / purchase ID
	ProductID  int64 `json:"product_id"`  // ID товара / product ID

	Quantity      float64    `json:"quantity"`       // количество / quantity
	PurchasePrice float64    `json:"purchase_price"` // цена закупки / purchase price
	ExpiryDate    *time.Time `json:"expiry_date"`    // срок годности / expiry date
	BatchNumber   *string    `json:"batch_number"`   // номер партии / batch number
}

type UpdatePurchaseItem struct {
	PurchaseID int64 `json:"purchase_id"` // ID закупки / purchase ID
	ProductID  int64 `json:"product_id"`  // ID товара / product ID

	Quantity      float64    `json:"quantity"`       // количество / quantity
	PurchasePrice float64    `json:"purchase_price"` // цена закупки / purchase price
	ExpiryDate    *time.Time `json:"expiry_date"`    // срок годности / expiry date
	BatchNumber   *string    `json:"batch_number"`   // номер партии / batch number
}
