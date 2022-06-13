package user

import (
	"github.com/buka-kitab/backend/domain/general"
	"github.com/buka-kitab/backend/usecase"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	User UserDataHandler
}

func NewHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) UserHandler {
	return UserHandler{
		User: newUserHandler(uc, conf, logger),
	}
}
