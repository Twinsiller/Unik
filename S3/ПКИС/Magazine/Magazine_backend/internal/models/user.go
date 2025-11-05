package models

import "time"

// User — сотрудник магазина (кассир, админ, менеджер)
type User struct {
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"` // уникальный ID сотрудника / unique user ID

	Nickname     string    `gorm:"type:text;not null" json:"nickname"`             // никнейм сотрудника / employee nickname
	HashPassword string    `gorm:"type:text;not null" json:"hashpassword"`         // (-)!!!!!!!!! хэш-пароль сотрудника / employee hashpassword
	Name         string    `gorm:"type:text;not null" json:"name"`                 // имя сотрудника / employee name
	Role         string    `gorm:"type:text;not null" json:"role"`                 // роль (cashier, manager, admin) / role
	Email        *string   `gorm:"type:text;unique" json:"email"`                  // почта / email
	HiredAt      time.Time `gorm:"type:timestamptz;default:now()" json:"hired_at"` // дата найма / hire date

	OrderList []Order `gorm:"foreignKey:CashierID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"order_list"` //gorm:"foreignKey:CashierID;references:ID; заказы, проведенные кассиром / orders created by this cashier
}

type CreateUser struct {
	Nickname string  `json:"nickname" binding:"required"`
	Password string  `json:"password" binding:"required"`
	Name     string  `json:"name" binding:"required"`
	Role     string  `json:"role" binding:"required"`
	Email    *string `json:"email"`
}

type UpdateUser struct {
	Nickname     *string `json:"nickname"`
	HashPassword *string `json:"hashpassword"`
	Name         *string `json:"name"`
	Role         *string `json:"role"`
	Email        *string `json:"email"`
}

type LoginUser struct {
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
}
