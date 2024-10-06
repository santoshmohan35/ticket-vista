package router

import (
	"github.com/gin-gonic/gin"
	"github.com/santoshmohan35/tickets-data-api/internal/configuration"
	"net/http"
)

func SetupRouter(config configuration.Config) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
	return router
}
