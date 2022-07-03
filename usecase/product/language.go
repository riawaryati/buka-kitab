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

type LanguageUsecaseItf interface {
	GetListLanguage(pagination gen.PaginationData, filter domain.LanguageFilter) ([]domain.Language, gen.PaginationData, string, error)
}

type LanguageUsecase struct {
	Repo   product.LanguageRepoItf
	DBList *infra.DatabaseList
	Log    *logrus.Logger
}

func newLanguageUsecase(r repo.Repo, logger *logrus.Logger, dbList *infra.DatabaseList) LanguageUsecase {
	return LanguageUsecase{
		Repo:   r.Product.Language,
		Log:    logger,
		DBList: dbList,
	}
}

func (pu LanguageUsecase) GetListLanguage(pagination gen.PaginationData, filter domain.LanguageFilter) ([]domain.Language, gen.PaginationData, string, error) {
	data, err := pu.Repo.GetListLanguage(pagination, filter)
	if err != nil {
		pu.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("GetListLanguage | fail to get language list from repo")
		return data, pagination, "", err
	}

	count, page, err := pu.Repo.GetTotalDataLanguage(pagination, filter)
	if err != nil {
		pu.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("error get total data language from repo")
		return data, pagination, "", err
	}

	pagination.TotalData = int(count)
	pagination.TotalPage = int(page)

	return data, pagination, general.SourceFromDB, nil
}
