package usecase

import (
	"github.com/buka-kitab/backend/domain/general"
	"github.com/buka-kitab/backend/infra"
	"github.com/buka-kitab/backend/repo"
	"github.com/buka-kitab/backend/usecase/master"
	"github.com/buka-kitab/backend/usecase/order"
	"github.com/buka-kitab/backend/usecase/user"
	"github.com/sirupsen/logrus"
)

type Usecase struct {
	Master master.MasterUsecase
	User   user.UserUsecase
	Order  order.OrderUsecase
}

func NewUsecase(repo repo.Repo, conf *general.SectionService, dbList *infra.DatabaseList, logger *logrus.Logger) Usecase {
	return Usecase{
		Master: master.NewUsecase(repo, conf, dbList, logger),
		User:   user.NewUsecase(repo, conf, dbList, logger),
		Order:  order.NewUsecase(repo, conf, dbList, logger),
	}
}
