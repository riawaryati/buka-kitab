package repo

import (
	"github.com/buka-kitab/backend/infra"
	m "github.com/buka-kitab/backend/repo/master"
	"github.com/buka-kitab/backend/repo/order"
	"github.com/buka-kitab/backend/repo/user"
	"github.com/sirupsen/logrus"
)

type Repo struct {
	Master m.MasterRepo
	User   user.UserRepo
	Order  order.OrderRepo
}

func NewRepo(db *infra.DatabaseList, logger *logrus.Logger) Repo {
	return Repo{
		Master: m.NewMasterRepo(db, logger),
		User:   user.NewMasterRepo(db, logger),
		Order:  order.NewMasterRepo(db, logger),
	}
}
