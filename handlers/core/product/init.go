package product

import (
	"github.com/buka-kitab/backend/domain/general"
	"github.com/buka-kitab/backend/usecase"
	"github.com/sirupsen/logrus"
)

type ProductHandler struct {
	Author    AuthorHandler
	Category  CategoryHandler
	Form      FormHandler
	Language  LanguageHandler
	Publisher PublisherHandler
}

func NewHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) ProductHandler {
	return ProductHandler{
		Author:    newAuthorHandler(uc, conf, logger),
		Category:  newCategoryHandler(uc, conf, logger),
		Form:      newFormHandler(uc, conf, logger),
		Language:  newLanguageHandler(uc, conf, logger),
		Publisher: newPublisherHandler(uc, conf, logger),
	}
}
