package models

import "time"

// StockMovement — история изменений остатков (аудит)
type StockMovement struct {
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"` // ID движения / movement ID

	StockID *int64 `gorm:"index" json:"stock_id"`            // ID записи склада / stock ID
	Stock   *Stock `gorm:"foreignKey:StockID;references:ID"` // связь / stock

	ProductID int64   `gorm:"index;not null" json:"product_id"`   // ID товара / product ID
	Product   Product `gorm:"foreignKey:ProductID;references:ID"` // связь / product

	Change float64 `gorm:"type:numeric(14,3);not null" json:"change"` // изменение количества (+/-) / change
	Reason string  `gorm:"type:text" json:"reason"`                   // причина (sale, purchase, spoilage) / reason

	RelatedOrderID *int64 `gorm:"index" json:"related_order_id"` // связанный заказ / related order
	Order          *Order `gorm:"foreignKey:RelatedOrderID;references:ID"`

	RelatedPurchaseID *int64    `gorm:"index" json:"related_purchase_id"` // связанная закупка / related purchase
	Purchase          *Purchase `gorm:"foreignKey:RelatedPurchaseID;references:ID"`

	CreatedAt time.Time `gorm:"type:timestamptz;default:now()" json:"created_at"` // дата изменения / change date

	ActorID *int64 `gorm:"index" json:"actor_id"`            // кто сделал действие (сотрудник) / actor ID
	Actor   *User  `gorm:"foreignKey:ActorID;references:ID"` // связь / actor
}
