package models

import "time"

// Order — чек продажи (заказ клиента)
type Order struct {
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"` // ID заказа / order ID

	CashierID int64 `gorm:"index;not null" json:"cashier_id"`   // ID кассира / cashier ID
	Cashier   User  `gorm:"foreignKey:CashierID;references:ID"` // связь / cashier

	CreatedAt     time.Time `gorm:"type:timestamptz;default:now()" json:"created_at"` // дата заказа / order date
	TotalAmount   float64   `gorm:"type:numeric(14,2);default:0" json:"total_amount"` // сумма заказа / total amount
	PaymentMethod string    `gorm:"type:text" json:"payment_method"`                  // способ оплаты / payment method
	ShiftID       *string   `gorm:"type:text;default:unknown" json:"shift_id"`        // смена кассира / shift ID

	ItemList []OrderItem `gorm:"foreignKey:OrderID;references:ID" json:"item_list"` // товары в заказе / order items
}

type CreateOrder struct {
	CashierID int64 `json:"cashier_id" binding:"required"` // ID кассира / cashier ID

	TotalAmount   float64 `json:"total_amount" binding:"required"`   // сумма заказа / total amount
	PaymentMethod string  `json:"payment_method" binding:"required"` // способ оплаты / payment method
	ShiftID       *string `json:"shift_id"`                          // смена кассира / shift ID
}

type UpdateOrder struct {
	CashierID *int64 `json:"cashier_id"` // ID кассира / cashier ID

	CreatedAt     *time.Time `json:"created_at"`     // дата заказа / order date
	TotalAmount   *float64   `json:"total_amount"`   // сумма заказа / total amount
	PaymentMethod *string    `json:"payment_method"` // способ оплаты / payment method
	ShiftID       *string    `json:"shift_id"`       // смена кассира / shift ID
}
