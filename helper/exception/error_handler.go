package exception

import (
	"fiber-rest-api/model/web/response"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var ErrorHandler = func(c *fiber.Ctx, err error) {
	if notFoundError(c, err) {
		webResponse := response.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
		}
		c.Status(http.StatusNotFound).JSON(webResponse)
		return
	}
	return
}

func notFoundError(c *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	log.Println("Error", exception)
	log.Println("STATUS", ok)

	if ok {
		return true
	} else {
		return false
	}

}
