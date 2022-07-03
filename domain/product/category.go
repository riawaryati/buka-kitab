package product

import "gopkg.in/guregu/null.v4"

type Category struct {
	CategoryID int    `json:"categoryId" gorm:"primaryKey;autoIncrement" db:"category_id"`
	Name       string `json:"name" db:"name"`
	IsActive   string `json:"isActive" db:"is_active"`
}

type CategoryFilter struct {
	Name     null.String
	IsActive null.Bool
}
