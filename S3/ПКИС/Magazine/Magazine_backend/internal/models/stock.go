package models

import "time"

// Stock — остатки товара на складе
type Stock struct {
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"` // ID записи / stock entry ID

	ProductID int64   `gorm:"index;not null" json:"product_id"`   // ID товара / product ID
	Product   Product `gorm:"foreignKey:ProductID;references:ID"` // связь / product

	Quantity      float64    `gorm:"type:numeric(14,3);not null" json:"quantity"`       // количество / quantity
	BatchNumber   *string    `gorm:"type:text" json:"batch_number"`                     // номер партии / batch number
	PurchasePrice *float64   `gorm:"type:numeric(12,4)" json:"purchase_price"`          // цена закупки / purchase price
	ReceivedAt    time.Time  `gorm:"type:timestamptz;default:now()" json:"received_at"` // дата приёмки / received date
	ExpiryDate    *time.Time `gorm:"type:date" json:"expiry_date"`                      // срок годности / expiry date

	Movements []StockMovement `gorm:"foreignKey:StockID;references:ID"` // движение по этой партии / stock movements
}

type CreateStock struct {
	ProductID     int64      `json:"product_id"`     // ID товара / product ID
	Quantity      float64    `json:"quantity"`       // количество / quantity
	BatchNumber   *string    `json:"batch_number"`   // номер партии / batch number
	PurchasePrice *float64   `json:"purchase_price"` // цена закупки / purchase price
	ReceivedAt    time.Time  `json:"received_at"`    // дата приёмки / received date
	ExpiryDate    *time.Time `json:"expiry_date"`    // срок годности / expiry date
}

type UpdateStock struct {
	ProductID     *int64     `json:"product_id"`     // ID товара / product ID
	Quantity      *float64   `json:"quantity"`       // количество / quantity
	BatchNumber   *string    `json:"batch_number"`   // номер партии / batch number
	PurchasePrice *float64   `json:"purchase_price"` // цена закупки / purchase price
	ReceivedAt    *time.Time `json:"received_at"`    // дата приёмки / received date
	ExpiryDate    *time.Time `json:"expiry_date"`    // срок годности / expiry date

}
