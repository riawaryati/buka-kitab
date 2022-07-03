package product

import (
	"fmt"
	"strings"

	dg "github.com/buka-kitab/backend/domain/general"
	dm "github.com/buka-kitab/backend/domain/product"
	"github.com/buka-kitab/backend/infra"
)

type PublisherRepo struct {
	DBList *infra.DatabaseList
}

func newPublisherRepo(dbList *infra.DatabaseList) PublisherRepo {
	return PublisherRepo{
		DBList: dbList,
	}
}

const (
	pqSelectPublisher = `
	SELECT
		publisher_id,
		name,
		is_active
	FROM
		publisher`

	pqCountPublisher = `
	SELECT
		COUNT(1) as count
	FROM
		publisher`

	pqWhere = `
	WHERE`

	pqFilterPublisherID = `
		publisher_id = ?`

	pqFilterName = `
		lower(name) LIKE ?`

	pqFilterIsActive = `
		is_active = ?`

	pqLimitOffset = `
	LIMIT ?
	OFFSET ?`

	pqOrderBy = `
	ORDER BY`
)

type PublisherRepoItf interface {
	GetByID(publisherID int64) (dm.Publisher, error)
	GetByName(name string) (dm.Publisher, error)
	GetListPublisher(pagination dg.PaginationData, filter dm.PublisherFilter) ([]dm.Publisher, error)
	GetTotalDataPublisher(pagination dg.PaginationData, filter dm.PublisherFilter) (int64, int64, error)
}

func (cr PublisherRepo) GetByID(publisherID int64) (dm.Publisher, error) {
	var res dm.Publisher

	q := fmt.Sprintf("%s%s%s", pqSelectPublisher, pqWhere, pqFilterPublisherID)
	query, args, err := cr.DBList.Backend.Read.In(q, publisherID)
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

func (cr PublisherRepo) GetByName(name string) (dm.Publisher, error) {
	var res dm.Publisher

	q := fmt.Sprintf("%s%s%s", pqSelectPublisher, pqWhere, pqFilterName)
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

func (cr PublisherRepo) GetListPublisher(pagination dg.PaginationData, filter dm.PublisherFilter) ([]dm.Publisher, error) {
	var result []dm.Publisher
	param := make([]interface{}, 0)
	var fl []string

	if filter.Name.Valid {
		fl = append(fl, pqFilterName)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.Name.String))+"%")
	}

	if filter.IsActive.Valid {
		fl = append(fl, pqFilterIsActive)
		param = append(param, filter.IsActive.Bool)
	}

	q := pqSelectPublisher

	if len(fl) > 0 {
		q += pqWhere + strings.Join(fl, " AND ")
	}

	// Add orderby value.
	q += " " + pqOrderBy + " " + pagination.OrderBy.String + " " + strings.ToLower(pagination.Sort)

	if !pagination.IsGetAll {
		// Add limit & page to param.
		q += pqLimitOffset
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

func (cr PublisherRepo) GetTotalDataPublisher(pagination dg.PaginationData, filter dm.PublisherFilter) (int64, int64, error) {
	var result int64
	param := make([]interface{}, 0)
	var fl []string

	if filter.Name.Valid {
		fl = append(fl, pqFilterName)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.Name.String))+"%")
	}

	if filter.IsActive.Valid {
		fl = append(fl, pqFilterIsActive)
		param = append(param, filter.IsActive.Bool)
	}

	q := pqCountPublisher

	if len(fl) > 0 {
		q += pqWhere + strings.Join(fl, " AND ")
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
