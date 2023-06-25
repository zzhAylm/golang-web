package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	log.Fatal(autotls.Run(r, "localhost:8080"))
}
