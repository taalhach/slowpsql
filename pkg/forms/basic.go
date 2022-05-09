package forms

import "github.com/taalhach/slowpsql/pkg/items"

type BasicResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

type BasicList struct {
	Limit     int      `json:"limit"`
	Page      int      `json:"page"`
	Filters   []string `json:"filters"`
	SortBy    string   `json:"sort_by"`
	SortOrder string   `json:"sort_order"`
}

func (bl *BasicList) AttachDefaults() {
	if bl.Limit == 0 {
		bl.Limit = items.DefaultPaginationLimit
	}

	if bl.Page == 0 {
		bl.Page = 1
	}
}
