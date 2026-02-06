package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("internal/templates/*")
	router.Static("/css", "./internal/css")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main",
		})
	})

	err := router.Run()
	if err != nil {
		log.Printf("Error starting Webserver")
	}
}
