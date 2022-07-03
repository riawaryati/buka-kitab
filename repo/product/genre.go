package product

import (
	"fmt"
	"strings"

	dg "github.com/buka-kitab/backend/domain/general"
	dm "github.com/buka-kitab/backend/domain/product"
	"github.com/buka-kitab/backend/infra"
)

type GenreRepo struct {
	DBList *infra.DatabaseList
}

func newGenreRepo(dbList *infra.DatabaseList) GenreRepo {
	return GenreRepo{
		DBList: dbList,
	}
}

const (
	gqSelectGenre = `
	SELECT
		genre_id,
		name,
		is_active
	FROM
		genres`

	gqCountGenre = `
	SELECT
		COUNT(1) as count
	FROM
		genres`

	gqWhere = `
	WHERE`

	gqFilterGenreID = `
		genre_id = ?`

	gqFilterName = `
		lower(name) LIKE ?`

	gqFilterIsActive = `
		is_active = ?`

	gqLimitOffset = `
	LIMIT ?
	OFFSET ?`

	gqOrderBy = `
	ORDER BY`
)

type GenreRepoItf interface {
	GetByID(genreID int64) (dm.Genre, error)
	GetByName(name string) (dm.Genre, error)
	GetListGenre(pagination dg.PaginationData, filter dm.GenreFilter) ([]dm.Genre, error)
	GetTotalDataGenre(pagination dg.PaginationData, filter dm.GenreFilter) (int64, int64, error)
}

func (cr GenreRepo) GetByID(genreID int64) (dm.Genre, error) {
	var res dm.Genre

	q := fmt.Sprintf("%s%s%s", gqSelectGenre, gqWhere, gqFilterGenreID)
	query, args, err := cr.DBList.Backend.Read.In(q, genreID)
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

func (cr GenreRepo) GetByName(name string) (dm.Genre, error) {
	var res dm.Genre

	q := fmt.Sprintf("%s%s%s", gqSelectGenre, gqWhere, gqFilterName)
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

func (cr GenreRepo) GetListGenre(pagination dg.PaginationData, filter dm.GenreFilter) ([]dm.Genre, error) {
	var result []dm.Genre
	param := make([]interface{}, 0)
	var fl []string

	if filter.Name.Valid {
		fl = append(fl, gqFilterName)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.Name.String))+"%")
	}

	if filter.IsActive.Valid {
		fl = append(fl, gqFilterIsActive)
		param = append(param, filter.IsActive.Bool)
	}

	q := gqSelectGenre

	if len(fl) > 0 {
		q += gqWhere + strings.Join(fl, " AND ")
	}

	// Add orderby value.
	q += " " + gqOrderBy + " " + pagination.OrderBy.String + " " + strings.ToLower(pagination.Sort)

	if !pagination.IsGetAll {
		// Add limit & page to param.
		q += gqLimitOffset
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

func (cr GenreRepo) GetTotalDataGenre(pagination dg.PaginationData, filter dm.GenreFilter) (int64, int64, error) {
	var result int64
	param := make([]interface{}, 0)
	var fl []string

	if filter.Name.Valid {
		fl = append(fl, gqFilterName)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.Name.String))+"%")
	}

	if filter.IsActive.Valid {
		fl = append(fl, gqFilterIsActive)
		param = append(param, filter.IsActive.Bool)
	}

	q := gqCountGenre

	if len(fl) > 0 {
		q += gqWhere + strings.Join(fl, " AND ")
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
