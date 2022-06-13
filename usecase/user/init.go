package user

import (
	"github.com/buka-kitab/backend/domain/general"
	"github.com/buka-kitab/backend/infra"
	"github.com/buka-kitab/backend/repo"
	"github.com/sirupsen/logrus"
)

type UserUsecase struct {
	User UserDataUsecaseItf
}

func NewUsecase(repo repo.Repo, conf *general.SectionService, dbList *infra.DatabaseList, logger *logrus.Logger) UserUsecase {
	return UserUsecase{
		User: newUserDataUsecase(repo, conf, logger, dbList),
	}
}
