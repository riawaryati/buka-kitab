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

type SubDistrictUsecaseItf interface {
	GetListSubDistrict(pagination gen.PaginationData, filter domain.SubDistrictFilter) ([]domain.SubDistrict, gen.PaginationData, string, error)
}

type SubDistrictUsecase struct {
	Repo   master.SubDistrictRepoItf
	DBList *infra.DatabaseList
	Log    *logrus.Logger
}

func newSubDistrictUsecase(r repo.Repo, logger *logrus.Logger, dbList *infra.DatabaseList) SubDistrictUsecase {
	return SubDistrictUsecase{
		Repo:   r.Master.SubDistrict,
		Log:    logger,
		DBList: dbList,
	}
}

func (sdu SubDistrictUsecase) GetListSubDistrict(pagination gen.PaginationData, filter domain.SubDistrictFilter) ([]domain.SubDistrict, gen.PaginationData, string, error) {
	data, err := sdu.Repo.GetListSubDistrict(pagination, filter)
	if err != nil {
		sdu.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("GetListSubDistrict | fail to get sub district list from repo")
		return data, pagination, "", err
	}

	count, page, err := sdu.Repo.GetTotalDataSubDistrict(pagination, filter)
	if err != nil {
		sdu.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("error get total data sub district from repo")
		return data, pagination, "", err
	}

	pagination.TotalData = int(count)
	pagination.TotalPage = int(page)

	return data, pagination, general.SourceFromDB, nil
}
