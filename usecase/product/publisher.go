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

type PublisherUsecaseItf interface {
	GetListPublisher(pagination gen.PaginationData, filter domain.PublisherFilter) ([]domain.Publisher, gen.PaginationData, string, error)
}

type PublisherUsecase struct {
	Repo   product.PublisherRepoItf
	DBList *infra.DatabaseList
	Log    *logrus.Logger
}

func newPublisherUsecase(r repo.Repo, logger *logrus.Logger, dbList *infra.DatabaseList) PublisherUsecase {
	return PublisherUsecase{
		Repo:   r.Product.Publisher,
		Log:    logger,
		DBList: dbList,
	}
}

func (pu PublisherUsecase) GetListPublisher(pagination gen.PaginationData, filter domain.PublisherFilter) ([]domain.Publisher, gen.PaginationData, string, error) {
	data, err := pu.Repo.GetListPublisher(pagination, filter)
	if err != nil {
		pu.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("GetListPublisher | fail to get publisher list from repo")
		return data, pagination, "", err
	}

	count, page, err := pu.Repo.GetTotalDataPublisher(pagination, filter)
	if err != nil {
		pu.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("error get total data publisher from repo")
		return data, pagination, "", err
	}

	pagination.TotalData = int(count)
	pagination.TotalPage = int(page)

	return data, pagination, general.SourceFromDB, nil
}
