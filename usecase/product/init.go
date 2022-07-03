package product

import (
	"github.com/buka-kitab/backend/domain/general"
	"github.com/buka-kitab/backend/infra"
	"github.com/buka-kitab/backend/repo"
	"github.com/sirupsen/logrus"
)

type ProductUsecase struct {
	Author    AuthorUsecaseItf
	Category  CategoryUsecaseItf
	Form      FormUsecaseItf
	Genre     GenreUsecaseItf
	Language  LanguageUsecaseItf
	Publisher PublisherUsecaseItf
}

func NewUsecase(repo repo.Repo, conf *general.SectionService, dbList *infra.DatabaseList, logger *logrus.Logger) ProductUsecase {
	return ProductUsecase{
		Author:    newAuthorUsecase(repo, logger, dbList),
		Category:  newCategoryUsecase(repo, logger, dbList),
		Form:      newFormUsecase(repo, logger, dbList),
		Genre:     newGenreUsecase(repo, logger, dbList),
		Language:  newLanguageUsecase(repo, logger, dbList),
		Publisher: newPublisherUsecase(repo, logger, dbList),
	}
}
