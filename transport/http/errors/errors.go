package router

import (
	"net/http"

	errors "sample-service/implementation/errors"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	c.Next()

	errs := c.Errors

	if len(errs) > 0 {
		err, ok := errs[0].Err.(*errors.AppError)
		if ok {
			switch err.Type {
			case errors.NotFound:
				c.JSON(http.StatusNotFound, err.Error())
				return
			case errors.ValidationError:
				c.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
}
