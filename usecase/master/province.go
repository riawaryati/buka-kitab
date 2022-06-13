package master

import (
	"github.com/buka-kitab/backend/constants/general"
	gen "github.com/buka-kitab/backend/domain/general"
	domain "github.com/buka-kitab/backend/domain/master"
	"github.com/buka-kitab/backend/infra"
	"github.com/buka-kitab/backend/repo"
	"github.com/buka-kitab/backend/repo/master"
	"github.com/buka-kitab/backend/utils"
	"github.com/sirupsen/logrus"
)

type ProvinceUsecaseItf interface {
	GetListProvince(pagination gen.PaginationData, filter domain.ProvinceFilter) ([]domain.Province, gen.PaginationData, string, error)
}

type ProvinceUsecase struct {
	Repo   master.ProvinceRepoItf
	DBList *infra.DatabaseList
	Log    *logrus.Logger
}

func newProvinceUsecase(r repo.Repo, logger *logrus.Logger, dbList *infra.DatabaseList) ProvinceUsecase {
	return ProvinceUsecase{
		Repo:   r.Master.Province,
		Log:    logger,
		DBList: dbList,
	}
}

func (pu ProvinceUsecase) GetListProvince(pagination gen.PaginationData, filter domain.ProvinceFilter) ([]domain.Province, gen.PaginationData, string, error) {
	data, err := pu.Repo.GetListProvince(pagination, filter)
	if err != nil {
		pu.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("GetListProvince | fail to get province list from repo")
		return data, pagination, "", err
	}

	count, page, err := pu.Repo.GetTotalDataProvince(pagination, filter)
	if err != nil {
		pu.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("error get total data province from repo")
		return data, pagination, "", err
	}

	pagination.TotalData = int(count)
	pagination.TotalPage = int(page)

	return data, pagination, general.SourceFromDB, nil
}
