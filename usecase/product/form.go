package product

import (
	"github.com/buka-kitab/backend/constants/general"
	gen "github.com/buka-kitab/backend/domain/general"
	domain "github.com/buka-kitab/backend/domain/product"
	"github.com/buka-kitab/backend/infra"
	"github.com/buka-kitab/backend/repo"
	"github.com/buka-kitab/backend/repo/product"
	"github.com/buka-kitab/backend/utils"
	"github.com/sirupsen/logrus"
)

type FormUsecaseItf interface {
	GetListForm(pagination gen.PaginationData, filter domain.FormFilter) ([]domain.Form, gen.PaginationData, string, error)
}

type FormUsecase struct {
	Repo   product.FormRepoItf
	DBList *infra.DatabaseList
	Log    *logrus.Logger
}

func newFormUsecase(r repo.Repo, logger *logrus.Logger, dbList *infra.DatabaseList) FormUsecase {
	return FormUsecase{
		Repo:   r.Product.Form,
		Log:    logger,
		DBList: dbList,
	}
}

func (pu FormUsecase) GetListForm(pagination gen.PaginationData, filter domain.FormFilter) ([]domain.Form, gen.PaginationData, string, error) {
	data, err := pu.Repo.GetListForm(pagination, filter)
	if err != nil {
		pu.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("GetListForm | fail to get form list from repo")
		return data, pagination, "", err
	}

	count, page, err := pu.Repo.GetTotalDataForm(pagination, filter)
	if err != nil {
		pu.Log.WithField("filter", utils.StructToString(filter)).WithError(err).Error("error get total data form from repo")
		return data, pagination, "", err
	}

	pagination.TotalData = int(count)
	pagination.TotalPage = int(page)

	return data, pagination, general.SourceFromDB, nil
}
