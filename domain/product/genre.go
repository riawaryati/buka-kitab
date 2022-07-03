package product

import "gopkg.in/guregu/null.v4"

type Genre struct {
	GenreID  int    `json:"genreId" gorm:"primaryKey;autoIncrement" db:"genre_id"`
	Name     string `json:"name" db:"name"`
	IsActive string `json:"isActive" db:"is_active"`
}

type GenreFilter struct {
	Name     null.String
	IsActive null.Bool
}
