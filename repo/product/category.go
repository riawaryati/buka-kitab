package product

import (
	"fmt"
	"strings"

	dg "github.com/buka-kitab/backend/domain/general"
	dm "github.com/buka-kitab/backend/domain/product"
	"github.com/buka-kitab/backend/infra"
)

type CategoryRepo struct {
	DBList *infra.DatabaseList
}

func newCategoryRepo(dbList *infra.DatabaseList) CategoryRepo {
	return CategoryRepo{
		DBList: dbList,
	}
}

const (
	cqSelectCategory = `
	SELECT
		category_id,
		name,
		is_active
	FROM
		categories`

	cqCountCategory = `
	SELECT
		COUNT(1) as count
	FROM
		categories`

	cqWhere = `
	WHERE`

	cqFilterCategoryID = `
		category_id = ?`

	cqFilterName = `
		lower(name) LIKE ?`

	cqFilterIsActive = `
		is_active = ?`

	cqLimitOffset = `
	LIMIT ?
	OFFSET ?`

	cqOrderBy = `
	ORDER BY`
)

type CategoryRepoItf interface {
	GetByID(categoryID int64) (dm.Category, error)
	GetByName(name string) (dm.Category, error)
	GetListCategory(pagination dg.PaginationData, filter dm.CategoryFilter) ([]dm.Category, error)
	GetTotalDataCategory(pagination dg.PaginationData, filter dm.CategoryFilter) (int64, int64, error)
}

func (cr CategoryRepo) GetByID(categoryID int64) (dm.Category, error) {
	var res dm.Category

	q := fmt.Sprintf("%s%s%s", cqSelectCategory, cqWhere, cqFilterCategoryID)
	query, args, err := cr.DBList.Backend.Read.In(q, categoryID)
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

func (cr CategoryRepo) GetByName(name string) (dm.Category, error) {
	var res dm.Category

	q := fmt.Sprintf("%s%s%s", cqSelectCategory, cqWhere, cqFilterName)
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

func (cr CategoryRepo) GetListCategory(pagination dg.PaginationData, filter dm.CategoryFilter) ([]dm.Category, error) {
	var result []dm.Category
	param := make([]interface{}, 0)
	var fl []string

	if filter.Name.Valid {
		fl = append(fl, cqFilterName)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.Name.String))+"%")
	}

	if filter.IsActive.Valid {
		fl = append(fl, cqFilterIsActive)
		param = append(param, filter.IsActive.Bool)
	}

	q := cqSelectCategory

	if len(fl) > 0 {
		q += cqWhere + strings.Join(fl, " AND ")
	}

	// Add orderby value.
	q += " " + cqOrderBy + " " + pagination.OrderBy.String + " " + strings.ToLower(pagination.Sort)

	if !pagination.IsGetAll {
		// Add limit & page to param.
		q += cqLimitOffset
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

func (cr CategoryRepo) GetTotalDataCategory(pagination dg.PaginationData, filter dm.CategoryFilter) (int64, int64, error) {
	var result int64
	param := make([]interface{}, 0)
	var fl []string

	if filter.Name.Valid {
		fl = append(fl, cqFilterName)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.Name.String))+"%")
	}

	if filter.IsActive.Valid {
		fl = append(fl, cqFilterIsActive)
		param = append(param, filter.IsActive.Bool)
	}

	q := cqCountCategory

	if len(fl) > 0 {
		q += cqWhere + strings.Join(fl, " AND ")
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
