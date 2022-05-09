package server

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/taalhach/slowpsql/pkg/forms"
)

var customErrHandler = func(c *fiber.Ctx, err error) error {
	var resp forms.BasicResponse
	respCode := 500

	// generally internal error should not only be visible in development mode
	if he, ok := err.(*fiber.Error); ok {
		respCode = he.Code
		resp.Message = fmt.Sprintf("%v", he.Message)
	} else {
		resp.Message = http.StatusText(respCode)
	}

	// 4 KB stack
	stack := make([]byte, 4<<10)
	length := runtime.Stack(stack, false)
	fmt.Printf("[RECOVER From Exception]: %v %s\n", err, stack[:length])

	return c.Status(respCode).JSON(err.Error())
}
