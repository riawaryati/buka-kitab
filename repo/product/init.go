package product

import (
	"github.com/buka-kitab/backend/infra"
	"github.com/sirupsen/logrus"
)

type ProductRepo struct {
	Author    AuthorRepoItf
	Category  CategoryRepoItf
	Form      FormRepoItf
	Genre     GenreRepoItf
	Language  LanguageRepoItf
	Publisher PublisherRepoItf
}

func NewMasterRepo(db *infra.DatabaseList, logger *logrus.Logger) ProductRepo {
	return ProductRepo{
		Author:    newAuthorRepo(db),
		Category:  newCategoryRepo(db),
		Form:      newFormRepo(db),
		Genre:     newGenreRepo(db),
		Language:  newLanguageRepo(db),
		Publisher: newPublisherRepo(db),
	}
}
