package product

import (
	"fmt"
	"strings"

	dg "github.com/buka-kitab/backend/domain/general"
	dm "github.com/buka-kitab/backend/domain/product"
	"github.com/buka-kitab/backend/infra"
)

type FormRepo struct {
	DBList *infra.DatabaseList
}

func newFormRepo(dbList *infra.DatabaseList) FormRepo {
	return FormRepo{
		DBList: dbList,
	}
}

const (
	fqSelectForm = `
	SELECT
		form_id,
		name,
		is_active
	FROM
		form`

	fqCountForm = `
	SELECT
		COUNT(1) as count
	FROM
		form`

	fqWhere = `
	WHERE`

	fqFilterFormID = `
		form_id = ?`

	fqFilterName = `
		lower(name) LIKE ?`

	fqFilterIsActive = `
		is_active = ?`

	fqLimitOffset = `
	LIMIT ?
	OFFSET ?`

	fqOrderBy = `
	ORDER BY`
)

type FormRepoItf interface {
	GetByID(formID int64) (dm.Form, error)
	GetByName(name string) (dm.Form, error)
	GetListForm(pagination dg.PaginationData, filter dm.FormFilter) ([]dm.Form, error)
	GetTotalDataForm(pagination dg.PaginationData, filter dm.FormFilter) (int64, int64, error)
}

func (cr FormRepo) GetByID(formID int64) (dm.Form, error) {
	var res dm.Form

	q := fmt.Sprintf("%s%s%s", fqSelectForm, fqWhere, fqFilterFormID)
	query, args, err := cr.DBList.Backend.Read.In(q, formID)
	if err != nil {
		return res, err
	}

	query = cr.DBList.Backend.Read.Rebind(query)
	err = cr.DBList.Backend.Read.Get(&res, query, args...)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (cr FormRepo) GetByName(name string) (dm.Form, error) {
	var res dm.Form

	q := fmt.Sprintf("%s%s%s", fqSelectForm, fqWhere, fqFilterName)
	query, args, err := cr.DBList.Backend.Read.In(q, strings.ToLower(name))
	if err != nil {
		return res, err
	}

	query = cr.DBList.Backend.Read.Rebind(query)
	err = cr.DBList.Backend.Read.Get(&res, query, args...)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (cr FormRepo) GetListForm(pagination dg.PaginationData, filter dm.FormFilter) ([]dm.Form, error) {
	var result []dm.Form
	param := make([]interface{}, 0)
	var fl []string

	if filter.Name.Valid {
		fl = append(fl, fqFilterName)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.Name.String))+"%")
	}

	if filter.IsActive.Valid {
		fl = append(fl, fqFilterIsActive)
		param = append(param, filter.IsActive.Bool)
	}

	q := fqSelectForm

	if len(fl) > 0 {
		q += fqWhere + strings.Join(fl, " AND ")
	}

	// Add orderby value.
	q += " " + fqOrderBy + " " + pagination.OrderBy.String + " " + strings.ToLower(pagination.Sort)

	if !pagination.IsGetAll {
		// Add limit & page to param.
		q += fqLimitOffset
		param = append(param, pagination.Limit)
		param = append(param, pagination.Offset)
	}

	query, args, err := cr.DBList.Backend.Read.In(q, param...)
	if err != nil {
		return result, err
	}

	query = cr.DBList.Backend.Read.Rebind(query)
	err = cr.DBList.Backend.Read.Select(&result, query, args...)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (cr FormRepo) GetTotalDataForm(pagination dg.PaginationData, filter dm.FormFilter) (int64, int64, error) {
	var result int64
	param := make([]interface{}, 0)
	var fl []string

	if filter.Name.Valid {
		fl = append(fl, fqFilterName)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.Name.String))+"%")
	}

	if filter.IsActive.Valid {
		fl = append(fl, fqFilterIsActive)
		param = append(param, filter.IsActive.Bool)
	}

	q := fqCountForm

	if len(fl) > 0 {
		q += fqWhere + strings.Join(fl, " AND ")
	}

	query, args, err := cr.DBList.Backend.Read.In(q, param...)
	if err != nil {
		return result, 0, err
	}

	//Run query to get total data
	query = cr.DBList.Backend.Read.Rebind(query)
	err = cr.DBList.Backend.Read.Get(&result, query, args...)
	if err != nil {
		return result, 0, err
	}

	//Calculate Total Page
	totalPage := result / int64(pagination.Limit)
	if result%int64(pagination.Limit) > 0 {
		totalPage++
	}

	return result, totalPage, nil
}
