package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const ver string = "0.2.0"

var version = new(bool)
var logfile = new(string)

func main() {
	version = flag.Bool("version", false, "a bool")
	logfile = flag.String("logfile", "", "a string")
	flag.Parse()

	if *version {
		fmt.Println("gin-api - v.", ver)
		os.Exit(0)
	}

	if *logfile != "" {
		fmt.Println("Use log file", *logfile)
		f, err := os.OpenFile(*logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		log.SetOutput(f)
		gin.DefaultWriter = io.MultiWriter(f)
		log.SetPrefix(time.Now().UTC().Format("2006-01-02T15:04:05Z") + " - ")
		log.Println("gin-api started")
	}

	// start API server
	gin.SetMode(gin.ReleaseMode)
	Server().Run(":3000")
}
