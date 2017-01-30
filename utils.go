package main

import (
	"compress/gzip"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Contains -
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Must -
func Must(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// GetDataFromFile - get data from a file
func GetDataFromFile(fileName string) []byte {
	if _, err := os.Stat(fileName); err == nil {
		data, e := ioutil.ReadFile(fileName)
		Must(e)
		return data
	}

	return []byte(`[]`)
}

// GetMongoMajorVersion -
func GetMongoMajorVersion(doc map[string]interface{}) int {
	if doc["version"].(string) == "" {
		return 0
	}

	n := strings.Index(doc["version"].(string), ".")
	if n < 0 {
		return 0
	}

	v, _ := strconv.ParseInt(doc["version"].(string)[:n], 10, 32)
	return int(v)
}

// GUNZip -
func GUNZip(c *gin.Context) {
	enc := c.Request.Header.Get("Content-Encoding")
	if enc == "gzip" {
		defer c.Request.Body.Close()
		r, _ := gzip.NewReader(c.Request.Body)
		c.Request.Body = r
	}
	c.Next()
}
