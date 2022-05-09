package common

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

const (
	eq        = "eq"
	neq       = "neq"
	startWith = "starts_with"
	endsWith  = "ends_with"
)

type BasicListRet struct {
	Page  int   `json:"page"`
	Pages int   `json:"pages"`
	Total int64 `json:"total"`
}

type Sort struct {
	By   string
	Desc bool
}

type BasicList struct {
	Query   *gorm.DB
	Columns map[string]string
	Sort    *Sort

	// pagination related
	Limit          int
	Page           int
	SkipPagination bool
	Filters        []string

	// sort related
	SkipOrder bool
}

//PrepareSql molds query and applies filters and pagination
func (form BasicList) PrepareSql() (*gorm.DB, error) {
	if !form.SkipOrder {
		if form.Sort != nil {
			column, ok := form.Columns[form.Sort.By]
			if !ok {
				return nil, fmt.Errorf("column mapping for %v not found", column)
			}

			sort := column
			if form.Sort.Desc {
				sort = fmt.Sprintf("%s DESC", column)
			}

			form.Query = form.Query.Order(sort)
		}
	}

	for _, filter := range form.Filters {
		subParts := strings.Split(filter, ":")
		if len(subParts) == 3 {
			col := subParts[0]
			operation := subParts[1]
			val := subParts[2]

			column, ok := form.Columns[col]
			if !ok {
				return nil, fmt.Errorf("column mapping for %v not found", column)
			}

			// TODO: add types(text, boolean, time, numbers) and more operation
			switch operation {
			case eq:
				form.Query = form.Query.Where(fmt.Sprintf("%s = ?", column), val)
			case neq:
				form.Query = form.Query.Where(fmt.Sprintf("%s <> ?", column), val)
			case startWith:
				form.Query = form.Query.Where(fmt.Sprintf("%s ILIKE ?", column), fmt.Sprintf("%s%%", val))
			case endsWith:
				form.Query = form.Query.Where(fmt.Sprintf("%s ILIKE ?", column), fmt.Sprintf("%%%s", val))
			}
		}
	}

	if !form.SkipPagination {
		form.Query = form.Paginate()
	}

	return form.Query, nil
}

//Paginate applies pagination if skipped in PrepareSql method
func (form BasicList) Paginate() *gorm.DB {
	offset := (form.Page - 1) * form.Limit
	form.Query = form.Query.Limit(form.Limit).Offset(offset)

	return form.Query
}

//ApplySort applies sorting if skipped in PrepareSql
func (form BasicList) ApplySort() (*gorm.DB, error) {
	if form.Sort != nil {
		column, ok := form.Columns[form.Sort.By]
		if !ok {
			return nil, fmt.Errorf("column mapping for %v not found", column)
		}

		sort := column
		if form.Sort.Desc {
			sort = fmt.Sprintf("%s DESC", column)
		}

		form.Query = form.Query.Order(sort)
	}

	return form.Query, nil
}
