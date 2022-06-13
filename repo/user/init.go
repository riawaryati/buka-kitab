package user

import (
	"github.com/buka-kitab/backend/infra"
	"github.com/sirupsen/logrus"
)

type UserRepo struct {
	User UserDataRepoItf
}

func NewMasterRepo(db *infra.DatabaseList, logger *logrus.Logger) UserRepo {
	return UserRepo{
		User: newUserDataRepo(db),
	}
}
