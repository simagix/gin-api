package main

import (
	"github.com/gin-gonic/gin"
)

// Server - spin an Gin API server
func Server() *gin.Engine {
	router := gin.Default()
	router.Use(DataStore)
	router.Use(GUNZip)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Gin API server",
		})
	})
	router.GET("/ping", Ping)

	osca := router.Group("/simagix")
	{
		osca.POST("/:hostname/commands/:type", VMPostStats)
	}
	return router
}
