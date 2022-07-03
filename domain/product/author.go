package product

import "gopkg.in/guregu/null.v4"

type Author struct {
	AuthorID   int    `json:"authorId" gorm:"primaryKey;autoIncrement" db:"author_id"`
	Name       string `json:"name" db:"name"`
	WebAddress string `json:"webAddress" db:"web_address"`
	About      string `json:"about" db:"about"`
	IsActive   string `json:"isActive" db:"is_active"`
}

type AuthorFilter struct {
	Name       null.String
	WebAddress null.String
	About      null.String
	IsActive   null.Bool
}
