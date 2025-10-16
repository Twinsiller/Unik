package models

// Category — категория товара (например, напитки, хлеб, мясо)
type Category struct {
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"` // ID категории / category ID

	Name string `gorm:"type:text;unique;not null" json:"name"` // название категории / category name

	Products []Product `gorm:"foreignKey:CategoryID;references:ID"` // товары этой категории / products in category
}
