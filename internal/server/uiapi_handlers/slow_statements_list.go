package uiapi_handlers

import (
	"math"

	"github.com/taalhach/slowpsql/internal/server/database/dbutils"
	"github.com/taalhach/slowpsql/pkg/forms"

	"github.com/taalhach/slowpsql/internal/server/common"
	"github.com/taalhach/slowpsql/internal/server/models"

	"github.com/gofiber/fiber/v2"
)

type SlowestStatementResponse struct {
	common.BasicListRet
	Items []*models.PgStatStatement `json:"items"`
}

//SlowestStatementsList finds slowest queries
func SlowestStatementsList(c *fiber.Ctx) error {
	form := new(forms.BasicList)
	if err := c.QueryParser(form); err != nil {
		return err
	}

	// load default params
	form.AttachDefaults()
	// find statements
	items, total, err := dbutils.FindStatements(form)
	if err != nil {
		return err
	}

	ret := SlowestStatementResponse{
		BasicListRet: common.BasicListRet{
			Page:  form.Page,
			Pages: int(math.Ceil(float64(total) / float64(form.Limit))),
			Total: total,
		},
		Items: items,
	}
	return c.JSON(ret)
}
