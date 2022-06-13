package order

import (
	"github.com/buka-kitab/backend/domain/general"
	"github.com/buka-kitab/backend/infra"
	"github.com/buka-kitab/backend/repo"
	"github.com/sirupsen/logrus"
)

type OrderUsecase struct {
	Order OrderDataUsecaseItf
}

func NewUsecase(repo repo.Repo, conf *general.SectionService, dbList *infra.DatabaseList, logger *logrus.Logger) OrderUsecase {
	return OrderUsecase{
		Order: newOrderDataUsecase(repo, conf, logger, dbList),
	}
}
