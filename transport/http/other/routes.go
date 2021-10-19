package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRoutesFactory() func(group *gin.RouterGroup) {
	otherRoutesFactory := func(group *gin.RouterGroup) {

		group.GET("/version", func(c *gin.Context) {
			c.String(http.StatusOK, "0.10.1")
		})

		group.GET("/health", func(c *gin.Context) {
			c.String(http.StatusOK, "ok")
		})
	}

	return otherRoutesFactory
}
