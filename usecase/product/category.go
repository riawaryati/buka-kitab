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

type CategoryUsecaseItf interface {
	GetListCategory(pagination gen.PaginationData, filter domain.CategoryFilter) ([]domain.Category, gen.PaginationData, string, error)
}

type CategoryUsecase struct {
	Repo   product.CategoryRepoItf
	DBList *infra.DatabaseList
	Log    *logrus.Logger
}

func newCategoryUsecase(r repo.Repo, logger *logrus.Logger, dbList *infra.DatabaseList) CategoryUsecase {
	return CategoryUsecase{
		Repo:   r.Product.Category,
		Log:    logger,
		DBList: dbList,
	}
}

func (pu CategoryUsecase) GetListCategory(pagination gen.PaginationData, filter domain.CategoryFilter) ([]domain.Category, gen.PaginationData, string, error) {
	data, err := pu.Repo.GetListCategory(pagination, filter)
	if err != nil {
		pu.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("GetListCategory | fail to get category list from repo")
		return data, pagination, "", err
	}

	count, page, err := pu.Repo.GetTotalDataCategory(pagination, filter)
	if err != nil {
		pu.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("error get total data category from repo")
		return data, pagination, "", err
	}

	pagination.TotalData = int(count)
	pagination.TotalPage = int(page)

	return data, pagination, general.SourceFromDB, nil
}
