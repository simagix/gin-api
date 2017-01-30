package main

import (
	"os"

	"github.com/gin-gonic/gin"

	mgo "gopkg.in/mgo.v2"
)

// MongoSession -
var MongoSession *mgo.Session

func init() {
	if os.Getenv("MYDB_URI") == "" {
		panic("MYDB_URI not defined")
	}
	session, err := mgo.Dial(os.Getenv("MYDB_URI"))
	if err != nil {
		panic(err.Error())
	}
	MongoSession = session
}

// DataStore -
func DataStore(c *gin.Context) {
	s := MongoSession.Clone()
	defer s.Close()
	c.Set("mydb", s.DB("MYDB"))
	c.Next()
}
