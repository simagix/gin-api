package main

import "github.com/gin-gonic/gin"

// Ping - ping server
//  GET /ping
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
