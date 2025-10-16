package models

// OrderItem — товар в заказе
type OrderItem struct {
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"` // ID позиции / order item ID

	OrderID int64 `gorm:"index;not null" json:"order_id"`   // ID заказа / order ID
	Order   Order `gorm:"foreignKey:OrderID;references:ID"` // связь / order

	ProductID int64   `gorm:"index;not null" json:"product_id"`   // ID товара / product ID
	Product   Product `gorm:"foreignKey:ProductID;references:ID"` // связь / product

	Quantity     float64 `gorm:"type:numeric(14,3);not null" json:"quantity"`       // количество / quantity
	PricePerUnit float64 `gorm:"type:numeric(12,2);not null" json:"price_per_unit"` // цена за единицу / unit price
	Discount     float64 `gorm:"type:numeric(12,2);default:0" json:"discount"`      // скидка / discount
	VatRate      float64 `gorm:"type:numeric(5,2);default:0" json:"vat_rate"`       // ставка НДС / VAT rate
}
