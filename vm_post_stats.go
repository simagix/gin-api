package main

import (
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

// VMPostStats - post vm stats of a type to server
//	POST /obim/osca/:hostname/commands/:type
func VMPostStats(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in VMPostStats", r)
			c.JSON(500, gin.H{"message": r})
		}
	}()

	hostname := c.Param("hostname")
	cmdType := c.Param("type")
	defer c.Request.Body.Close()
	_, err := ioutil.ReadAll(c.Request.Body)
	Must(err)
	// mydb := c.MustGet("mydb").(*mgo.Database)
	c.JSON(201, gin.H{"ok": 1, "hostname": hostname, "command": cmdType})
}
