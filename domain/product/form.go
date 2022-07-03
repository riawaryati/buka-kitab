package product

import "gopkg.in/guregu/null.v4"

type Form struct {
	FormID   int    `json:"formId" gorm:"primaryKey;autoIncrement" db:"form_id"`
	Name     string `json:"name" db:"name"`
	IsActive string `json:"isActive" db:"is_active"`
}

type FormFilter struct {
	Name     null.String
	IsActive null.Bool
}
