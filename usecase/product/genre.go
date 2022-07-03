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

type GenreUsecaseItf interface {
	GetListGenre(pagination gen.PaginationData, filter domain.GenreFilter) ([]domain.Genre, gen.PaginationData, string, error)
}

type GenreUsecase struct {
	Repo   product.GenreRepoItf
	DBList *infra.DatabaseList
	Log    *logrus.Logger
}

func newGenreUsecase(r repo.Repo, logger *logrus.Logger, dbList *infra.DatabaseList) GenreUsecase {
	return GenreUsecase{
		Repo:   r.Product.Genre,
		Log:    logger,
		DBList: dbList,
	}
}

func (pu GenreUsecase) GetListGenre(pagination gen.PaginationData, filter domain.GenreFilter) ([]domain.Genre, gen.PaginationData, string, error) {
	data, err := pu.Repo.GetListGenre(pagination, filter)
	if err != nil {
		pu.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("GetListGenre | fail to get genre list from repo")
		return data, pagination, "", err
	}

	count, page, err := pu.Repo.GetTotalDataGenre(pagination, filter)
	if err != nil {
		pu.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("error get total data genre from repo")
		return data, pagination, "", err
	}

	pagination.TotalData = int(count)
	pagination.TotalPage = int(page)

	return data, pagination, general.SourceFromDB, nil
}
