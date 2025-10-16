package models

import "time"

// Product — товар, который продаётся в магазине
type Product struct {
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"` // ID товара / product ID

	Name string `gorm:"type:text;not null" json:"name"` // название товара / product name

	ManufacturerID *int64        `gorm:"index" json:"manufacturer_id"`            // ID производителя / manufacturer ID
	Manufacturer   *Manufacturer `gorm:"foreignKey:ManufacturerID;references:ID"` // связь с производителем / manufacturer

	CategoryID *int64    `gorm:"index" json:"category_id"`            // ID категории / category ID
	Category   *Category `gorm:"foreignKey:CategoryID;references:ID"` // связь с категорией / category

	Barcode      *string   `gorm:"type:text;unique" json:"barcode"`                  // штрихкод / barcode
	DefaultPrice float64   `gorm:"type:numeric(12,2);not null" json:"default_price"` // базовая цена / default price
	MinStock     float64   `gorm:"type:numeric(12,2);default:0" json:"min_stock"`    // минимальный запас / minimum stock
	CreatedAt    time.Time `gorm:"type:timestamptz;default:now()" json:"created_at"` // дата добавления / creation date

	OrderItems     []OrderItem     `gorm:"foreignKey:ProductID;references:ID"` // продажи / order items
	PurchaseItems  []PurchaseItem  `gorm:"foreignKey:ProductID;references:ID"` // закупки / purchase items
	Stocks         []Stock         `gorm:"foreignKey:ProductID;references:ID"` // остатки на складе / stock entries
	StockMovements []StockMovement `gorm:"foreignKey:ProductID;references:ID"` // движение по складу / stock movements
}
