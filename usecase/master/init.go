package master

import (
	"github.com/buka-kitab/backend/domain/general"
	"github.com/buka-kitab/backend/infra"
	"github.com/buka-kitab/backend/repo"
	"github.com/sirupsen/logrus"
)

type MasterUsecase struct {
	Country     CountryUsecaseItf
	Province    ProvinceUsecaseItf
	City        CityUsecaseItf
	District    DistrictUsecaseItf
	SubDistrict SubDistrictUsecaseItf
}

func NewUsecase(repo repo.Repo, conf *general.SectionService, dbList *infra.DatabaseList, logger *logrus.Logger) MasterUsecase {
	return MasterUsecase{
		Country:     newCountryUsecase(repo, logger, dbList),
		Province:    newProvinceUsecase(repo, logger, dbList),
		City:        newCityUsecase(repo, logger, dbList),
		District:    newDistrictUsecase(repo, logger, dbList),
		SubDistrict: newSubDistrictUsecase(repo, logger, dbList),
	}
}
