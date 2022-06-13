package order

import (
	"github.com/buka-kitab/backend/domain/general"
	"github.com/buka-kitab/backend/usecase"
	"github.com/sirupsen/logrus"
)

type OrderHandler struct {
	Order OrderDataHandler
}

func NewHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) OrderHandler {
	return OrderHandler{
		Order: newOrderHandler(uc, conf, logger),
	}
}
