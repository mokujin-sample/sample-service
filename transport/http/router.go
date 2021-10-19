package router

import (
	"net/http"

	svc "sample-service"

	errors "sample-service/transport/http/errors"
	objectRoutes "sample-service/transport/http/objects"
	otherRoutes "sample-service/transport/http/other"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewHTTPHandler(objectService svc.ObjectService) http.Handler {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	router.Use(errors.Handler)

	otherGroup := router.Group("/other")
	otherRoutes.NewRoutesFactory()(otherGroup)

	api := router.Group("/api")
	objectsGroup := api.Group("/objects")
	objectRoutes.NewRoutesFactory(objectsGroup)(objectService)
	return router
}
