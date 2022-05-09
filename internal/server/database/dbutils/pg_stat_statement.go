package dbutils

import (
	"strings"

	"github.com/taalhach/slowpsql/internal/server/common"
	"github.com/taalhach/slowpsql/internal/server/database"
	"github.com/taalhach/slowpsql/internal/server/models"
	"github.com/taalhach/slowpsql/pkg/forms"
)

//FindStatements finds statements, supports pagination
// by default sorts rows on mean_time but sort can applied to any columns as mentioned in following map
func FindStatements(form *forms.BasicList) ([]*models.PgStatStatement, int64, error) {
	var (
		err   error
		total int64
		items []*models.PgStatStatement
	)
	// prepare query
	query := database.Db.Model(&models.PgStatStatement{}).Joins("LEFT JOIN pg_database AS pdb ON pdb.oid = pg_stat_statements.dbid")
	columns := map[string]string{
		"query":     "query",
		"mean_time": "mean_time",
		"min_time":  "min_time",
		"max_time":  "max_time",
		"database":  "pdb.datname",
	}

	// attach default sort
	if form.SortBy == "" {
		form.SortBy = "mean_time"
	}

	if form.SortOrder == "" {
		form.SortOrder = "DESC"
	}

	// prepare list params
	listParams := common.BasicList{
		Limit:   form.Limit,
		Page:    form.Page,
		Filters: form.Filters,
		Query:   query,
		Columns: columns,
	}

	// set sort column
	listParams.Sort = &common.Sort{
		By:   form.SortBy,
		Desc: strings.EqualFold(form.SortOrder, "DESC"),
	}

	// first skip so that we can get count of all first
	listParams.SkipPagination = true
	listParams.Query, err = listParams.PrepareSql()
	if err != nil {
		return nil, 0, err
	}

	// count total items
	if err = listParams.Query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// apply pagination by adding limit and offset
	listParams.Query = listParams.Paginate()

	const selectColumns = "query, min_time/1000 AS min_time_secs, max_time/1000 AS max_time_secs, mean_time/1000 AS mean_time_secs, datname AS database"
	if err := listParams.Query.Select(selectColumns).Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}
