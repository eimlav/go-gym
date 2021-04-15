package router

import (
	"net/http"

	v1 "github.com/eimlav/go-gym/api/router/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() http.Handler {
	router := gin.Default()

	apiRouter := router.Group("/api")
	{
		v1Router := apiRouter.Group("/v1")
		{
			v1Router.POST("/classes", v1.HandleClassesPOST)
		}
	}

	return router
}
