package product

import (
	"net/http"
	"strconv"

	cg "github.com/buka-kitab/backend/constants/general"
	"github.com/buka-kitab/backend/domain/general"
	dm "github.com/buka-kitab/backend/domain/product"
	"github.com/buka-kitab/backend/handlers"
	"github.com/buka-kitab/backend/usecase"
	um "github.com/buka-kitab/backend/usecase/product"
	"github.com/buka-kitab/backend/utils"
	"github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v4"
)

type FormHandler struct {
	Usecase um.FormUsecaseItf
	conf    *general.SectionService
	log     *logrus.Logger
}

func newFormHandler(uc usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) FormHandler {
	return FormHandler{
		Usecase: uc.Product.Form,
		conf:    conf,
		log:     logger,
	}
}

func (dh FormHandler) GetListForm(res http.ResponseWriter, req *http.Request) {
	respData := &handlers.ResponseData{
		Status: cg.Fail,
	}

	var tableFilter dm.FormFilter
	var err error

	paginationData := general.GetPagination()

	// Check province name value
	if req.FormValue("name") != "" {
		tableFilter.Name = null.StringFrom(req.FormValue("name"))
	}

	// Check sort value
	if req.FormValue("sort") != "" {
		paginationData.Sort = req.FormValue(("sort"))
	}

	// Check page value. If exist, convert to int
	if req.FormValue("page") != "" {
		paginationData.Page, err = strconv.Atoi(req.FormValue("page"))
		if err != nil {
			respData.Message = cg.HandlerErrorRequestDataFormatInvalid
			handlers.WriteResponse(res, respData, http.StatusBadRequest)
			return
		}
	}

	// Check orderby value.
	paginationData.OrderBy = null.StringFrom("form_id")
	if req.FormValue("order-by") != "" {
		paginationData.OrderBy.String = req.FormValue("order-by")
	}

	// Check isGetAll value.
	if req.FormValue("is-get-all") != "" {
		paginationData.IsGetAll = utils.GetBool(req.FormValue("is-get-all"))
	}

	// Check limit value. If exists, convert to int
	if req.FormValue("limit") != "" {
		paginationData.Limit, err = strconv.Atoi(req.FormValue("limit"))
		if err != nil {
			respData.Message = cg.HandlerErrorRequestDataFormatInvalid
			handlers.WriteResponse(res, respData, http.StatusBadRequest)
			return
		}
	}

	// Convert page to offset
	paginationData.SetOffset()

	data, paginationData, _, err := dh.Usecase.GetListForm(paginationData, tableFilter)
	if err != nil {
		respData.Message = "fail to get list form"
		handlers.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &handlers.ResponseData{
		Status:  cg.Success,
		Message: "success get list form",
		Detail: general.ResponseData{
			Data:       data,
			Pagination: paginationData,
		},
	}

	handlers.WriteResponse(res, respData, http.StatusOK)
	return
}
