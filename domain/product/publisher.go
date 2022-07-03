package product

import "gopkg.in/guregu/null.v4"

type Publisher struct {
	PublisherID int    `json:"publisherId" gorm:"primaryKey;autoIncrement" db:"publisher_id"`
	Name        string `json:"name" db:"name"`
	IsActive    string `json:"isActive" db:"is_active"`
}

type PublisherFilter struct {
	Name     null.String
	IsActive null.Bool
}
