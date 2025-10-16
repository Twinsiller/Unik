package models

import "time"

// Purchase — закупка у поставщика
type Purchase struct {
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"` // ID закупки / purchase ID

	SupplierID int64    `gorm:"index" json:"supplier_id"`            // ID поставщика / supplier ID
	Supplier   Supplier `gorm:"foreignKey:SupplierID;references:ID"` // связь / supplier

	CreatedAt     time.Time `gorm:"type:timestamptz;default:now()" json:"created_at"` // дата закупки / purchase date
	TotalAmount   float64   `gorm:"type:numeric(14,2);default:0" json:"total_amount"` // сумма закупки / total amount
	InvoiceNumber *string   `gorm:"type:text" json:"invoice_number"`                  // номер накладной / invoice number

	Items []PurchaseItem `gorm:"foreignKey:PurchaseID;references:ID"` // список товаров в закупке / purchase items
}
