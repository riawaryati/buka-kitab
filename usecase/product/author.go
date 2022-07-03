package product

import (
	"github.com/buka-kitab/backend/constants/general"
	gen "github.com/buka-kitab/backend/domain/general"
	domain "github.com/buka-kitab/backend/domain/product"
	"github.com/buka-kitab/backend/infra"
	"github.com/buka-kitab/backend/repo"
	"github.com/buka-kitab/backend/repo/product"
	"github.com/buka-kitab/backend/utils"
	"github.com/sirupsen/logrus"
)

type AuthorUsecaseItf interface {
	GetListAuthor(pagination gen.PaginationData, filter domain.AuthorFilter) ([]domain.Author, gen.PaginationData, string, error)
}

type AuthorUsecase struct {
	Repo   product.AuthorRepoItf
	DBList *infra.DatabaseList
	Log    *logrus.Logger
}

func newAuthorUsecase(r repo.Repo, logger *logrus.Logger, dbList *infra.DatabaseList) AuthorUsecase {
	return AuthorUsecase{
		Repo:   r.Product.Author,
		Log:    logger,
		DBList: dbList,
	}
}

func (pu AuthorUsecase) GetListAuthor(pagination gen.PaginationData, filter domain.AuthorFilter) ([]domain.Author, gen.PaginationData, string, error) {
	data, err := pu.Repo.GetListAuthor(pagination, filter)
	if err != nil {
		pu.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("GetListAuthor | fail to get author list from repo")
		return data, pagination, "", err
	}

	count, page, err := pu.Repo.GetTotalDataAuthor(pagination, filter)
	if err != nil {
		pu.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("error get total data author from repo")
		return data, pagination, "", err
	}

	pagination.TotalData = int(count)
	pagination.TotalPage = int(page)

	return data, pagination, general.SourceFromDB, nil
}
