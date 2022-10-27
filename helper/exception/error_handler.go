package exception

import (
	"github.com/gofiber/fiber/v2"
)

var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
	var (
		msg  string
		code int
		d    struct{}
	)

	code = fiber.StatusInternalServerError
	msg = "INTERNAL SERVER ERROR"
	if e, ok := err.(*fiber.Error); ok {
		if e.Code == fiber.StatusBadRequest {
			msg = "BAD REQUEST"
			code = e.Code
		} else if e.Code == fiber.StatusMethodNotAllowed {
			msg = "METHOD NOT ALLOWED"
			code = e.Code
		}
	}
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	return c.Status(code).JSON(fiber.Map{
		"status":  msg,
		"success": false,
		"message": err.Error(),
		"data":    d,
	})
}

// var HandlerErrorNew = func(c *fiber.Ctx, err error) error {
// 	if notFoundError(c, err) {
// 		return
// 	}

// 	internalServerError(c, err)
// }

// func notFoundError(c *fiber.Ctx, err interface{}) bool {
// 	exception, ok := err.(NotFoundError)
// 	if ok {
// 		webResponse := response.WebResponse{
// 			Code:   fiber.StatusNotFound,
// 			Status: "NOT FOUND",
// 			Data:   exception.Error,
// 		}

// 		c.Status(fiber.StatusNotFound).JSON(webResponse)
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func internalServerError(c *fiber.Ctx, err interface{}) {
// 	webResponse := response.WebResponse{
// 		Code:   fiber.StatusInternalServerError,
// 		Status: "INTERNAL SERVER ERROR BARU",
// 		Data:   err,
// 	}

// 	c.Status(fiber.StatusInternalServerError).JSON(webResponse)
// }
