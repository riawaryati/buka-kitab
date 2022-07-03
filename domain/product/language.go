package product

import "gopkg.in/guregu/null.v4"

type Language struct {
	LanguageID int    `json:"languageId" gorm:"primaryKey;autoIncrement" db:"language_id"`
	Name       string `json:"name" db:"name"`
	IsActive   string `json:"isActive" db:"is_active"`
}

type LanguageFilter struct {
	Name     null.String
	IsActive null.Bool
}
