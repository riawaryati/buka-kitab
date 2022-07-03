package product

import (
	"fmt"
	"strings"

	dg "github.com/buka-kitab/backend/domain/general"
	dm "github.com/buka-kitab/backend/domain/product"
	"github.com/buka-kitab/backend/infra"
)

type LanguageRepo struct {
	DBList *infra.DatabaseList
}

func newLanguageRepo(dbList *infra.DatabaseList) LanguageRepo {
	return LanguageRepo{
		DBList: dbList,
	}
}

const (
	lqSelectLanguage = `
	SELECT
		language_id,
		name,
		is_active
	FROM
		language`

	lqCountLanguage = `
	SELECT
		COUNT(1) as count
	FROM
		language`

	lqWhere = `
	WHERE`

	lqFilterLanguageID = `
		language_id = ?`

	lqFilterName = `
		lower(name) LIKE ?`

	lqFilterIsActive = `
		is_active = ?`

	lqLimitOffset = `
	LIMIT ?
	OFFSET ?`

	lqOrderBy = `
	ORDER BY`
)

type LanguageRepoItf interface {
	GetByID(languageID int64) (dm.Language, error)
	GetByName(name string) (dm.Language, error)
	GetListLanguage(pagination dg.PaginationData, filter dm.LanguageFilter) ([]dm.Language, error)
	GetTotalDataLanguage(pagination dg.PaginationData, filter dm.LanguageFilter) (int64, int64, error)
}

func (cr LanguageRepo) GetByID(languageID int64) (dm.Language, error) {
	var res dm.Language

	q := fmt.Sprintf("%s%s%s", lqSelectLanguage, lqWhere, lqFilterLanguageID)
	query, args, err := cr.DBList.Backend.Read.In(q, languageID)
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

func (cr LanguageRepo) GetByName(name string) (dm.Language, error) {
	var res dm.Language

	q := fmt.Sprintf("%s%s%s", lqSelectLanguage, lqWhere, lqFilterName)
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

func (cr LanguageRepo) GetListLanguage(pagination dg.PaginationData, filter dm.LanguageFilter) ([]dm.Language, error) {
	var result []dm.Language
	param := make([]interface{}, 0)
	var fl []string

	if filter.Name.Valid {
		fl = append(fl, lqFilterName)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.Name.String))+"%")
	}

	if filter.IsActive.Valid {
		fl = append(fl, lqFilterIsActive)
		param = append(param, filter.IsActive.Bool)
	}

	q := lqSelectLanguage

	if len(fl) > 0 {
		q += lqWhere + strings.Join(fl, " AND ")
	}

	// Add orderby value.
	q += " " + lqOrderBy + " " + pagination.OrderBy.String + " " + strings.ToLower(pagination.Sort)

	if !pagination.IsGetAll {
		// Add limit & page to param.
		q += lqLimitOffset
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

func (cr LanguageRepo) GetTotalDataLanguage(pagination dg.PaginationData, filter dm.LanguageFilter) (int64, int64, error) {
	var result int64
	param := make([]interface{}, 0)
	var fl []string

	if filter.Name.Valid {
		fl = append(fl, lqFilterName)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.Name.String))+"%")
	}

	if filter.IsActive.Valid {
		fl = append(fl, lqFilterIsActive)
		param = append(param, filter.IsActive.Bool)
	}

	q := lqCountLanguage

	if len(fl) > 0 {
		q += lqWhere + strings.Join(fl, " AND ")
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
