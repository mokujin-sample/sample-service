package objects

import (
	"net/http"
	"strconv"

	svc "sample-service"

	"github.com/gin-gonic/gin"
)

func NewRoutesFactory(group *gin.RouterGroup) func(service svc.ObjectService) {
	objectsRoutesFactory := func(service svc.ObjectService) {

		group.GET("/:objectId", func(c *gin.Context) {
			id := c.Param("objectId")
			i, err := strconv.Atoi(id)
			if err != nil {
				c.Error(err)
				return
			}
			result, err := service.Get(uint64(i))
			if err != nil {
				c.Error(err)
				return
			}

			c.JSON(http.StatusOK, *toResponseModel(&result))
		})

		group.GET("/statistics", func(c *gin.Context) {
			var query struct {
				IP     []string `form:"ip" binding:"dive,omitempty,cidrv4"`
				Period string   `form:"period,default=3h" binding:"oneof=15m 30m 1h 3h 1d 1w"`
			}

			if err := c.ShouldBindQuery(&query); err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
				return
			}

			data, err := service.GetData(query.IP, query.Period)
			if err != nil {
				c.Error(err)
				return
			}

			c.JSON(http.StatusOK, data)
		})

		group.GET("/under-object", func(c *gin.Context) {
			var query struct {
				AS []uint32 `form:"as" binding:"dive,gt=0"`
				IP []string `form:"ip" binding:"dive,omitempty,cidrv4"`
			}

			if err := c.ShouldBindQuery(&query); err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
				return
			}

			result, err := service.UnderObject(query.AS, query.IP)
			if err != nil {
				c.Error(err)
				return
			}
			c.JSON(http.StatusOK, gin.H{"status": result})
		})

		group.GET("/dictionary/object", func(c *gin.Context) {
			data := service.GetObjectDictionary()
			c.JSON(http.StatusOK, data)
		})
	}

	return objectsRoutesFactory
}
