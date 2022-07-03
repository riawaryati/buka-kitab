package product

import (
	"fmt"
	"strings"

	dg "github.com/buka-kitab/backend/domain/general"
	dm "github.com/buka-kitab/backend/domain/product"
	"github.com/buka-kitab/backend/infra"
)

type AuthorRepo struct {
	DBList *infra.DatabaseList
}

func newAuthorRepo(dbList *infra.DatabaseList) AuthorRepo {
	return AuthorRepo{
		DBList: dbList,
	}
}

const (
	aqSelectAuthor = `
	SELECT
		author_id,
		name,
		web_address,
		about,
		is_active
	FROM
		authors`

	aqCountAuthor = `
	SELECT
		COUNT(1) as count
	FROM
		authors`

	aqWhere = `
	WHERE`

	aqFilterAuthorID = `
		author_id = ?`

	aqFilterName = `
		lower(name) LIKE ?`

	aqFilterWebAddress = `
		lower(web_address) LIKE ?`

	aqFilterAbout = `
		lower(about) LIKE ?`

	aqFilterIsActive = `
		is_active = ?`

	aqLimitOffset = `
	LIMIT ?
	OFFSET ?`

	aqOrderBy = `
	ORDER BY`
)

type AuthorRepoItf interface {
	GetByID(authorID int64) (dm.Author, error)
	GetByName(name string) (dm.Author, error)
	GetListAuthor(pagination dg.PaginationData, filter dm.AuthorFilter) ([]dm.Author, error)
	GetTotalDataAuthor(pagination dg.PaginationData, filter dm.AuthorFilter) (int64, int64, error)
}

func (cr AuthorRepo) GetByID(authorID int64) (dm.Author, error) {
	var res dm.Author

	q := fmt.Sprintf("%s%s%s", aqSelectAuthor, aqWhere, aqFilterAuthorID)
	query, args, err := cr.DBList.Backend.Read.In(q, authorID)
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

func (cr AuthorRepo) GetByName(name string) (dm.Author, error) {
	var res dm.Author

	q := fmt.Sprintf("%s%s%s", aqSelectAuthor, aqWhere, aqFilterName)
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

func (cr AuthorRepo) GetListAuthor(pagination dg.PaginationData, filter dm.AuthorFilter) ([]dm.Author, error) {
	var result []dm.Author
	param := make([]interface{}, 0)
	var fl []string

	if filter.Name.Valid {
		fl = append(fl, aqFilterName)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.Name.String))+"%")
	}

	if filter.WebAddress.Valid {
		fl = append(fl, aqFilterWebAddress)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.WebAddress.String))+"%")
	}

	if filter.About.Valid {
		fl = append(fl, aqFilterAbout)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.About.String))+"%")
	}

	if filter.IsActive.Valid {
		fl = append(fl, aqFilterIsActive)
		param = append(param, filter.IsActive.Bool)
	}

	q := aqSelectAuthor

	if len(fl) > 0 {
		q += aqWhere + strings.Join(fl, " AND ")
	}

	// Add orderby value.
	q += " " + aqOrderBy + " " + pagination.OrderBy.String + " " + strings.ToLower(pagination.Sort)

	if !pagination.IsGetAll {
		// Add limit & page to param.
		q += aqLimitOffset
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

func (cr AuthorRepo) GetTotalDataAuthor(pagination dg.PaginationData, filter dm.AuthorFilter) (int64, int64, error) {
	var result int64
	param := make([]interface{}, 0)
	var fl []string

	if filter.Name.Valid {
		fl = append(fl, aqFilterName)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.Name.String))+"%")
	}

	if filter.WebAddress.Valid {
		fl = append(fl, aqFilterWebAddress)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.WebAddress.String))+"%")
	}

	if filter.About.Valid {
		fl = append(fl, aqFilterAbout)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.About.String))+"%")
	}

	if filter.IsActive.Valid {
		fl = append(fl, aqFilterIsActive)
		param = append(param, filter.IsActive.Bool)
	}

	q := aqCountAuthor

	if len(fl) > 0 {
		q += aqWhere + strings.Join(fl, " AND ")
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
