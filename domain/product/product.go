package product

import "time"

type Product struct {
	ProductID    int       `json:"productId" gorm:"primaryKey;autoIncrement" db:"product_id"`
	ProductName  string    `json:"name" db:"name"`
	FormID       int       `json:"formId" db:"form_id"`
	CategoryID   int       `json:"categoryId" db:"category_id"`
	GenreID      int       `json:"genreId" db:"genre_id"`
	AuthorID     int       `json:"authorId" db:"author_id"`
	Stock        int       `json:"stock" db:"stock"`
	ISBN         string    `json:"isbn" db:"isbn"`
	ISBN13       string    `json:"isbn13" db:"isbn13"`
	ReleaseDAte  time.Time `json:"releaseDate" db:"release_date"`
	LanguageID   int       `json:"languageId" db:"language_id"`
	PublisherID  int       `json:"publisherId" db:"publisher_id"`
	NumberOfPage int       `json:"numberOfPage" db:"number_of_page"`
	Description  string    `json:"description" db:"description"`
	IsActive     string    `json:"isActive" db:"is_active"`
}
