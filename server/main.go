package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()
	default_router := gin.Default()
	default_router.LoadHTMLGlob("templates/*")
	default_router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	default_router.GET("/dockerfile/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "lazo.html", gin.H{})
	})
	default_router.POST("/dockerfile/create", func(c *gin.Context) {

		c.HTML(http.StatusOK, "lazo.html", gin.H{})
	})
	default_router.Run(":8080")
}
