package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Scoping",
		})
	})
	router.GET("/task/", getAllTasks)
	router.GET("/task/:id", getTask)
	router.POST("/task/create", createTask)
	router.PUT("/task/:id", updateTask)
	router.DELETE("/task/:id", deleteTask)

	router.Run(":8080")
}
