package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test/:handson-boy", get)
	r.POST("/test", post)
	r.PUT("/test", put)
	r.DELETE("/test", delete)
	r.Run(":8080")
}

func get(c *gin.Context) {
	var a string
	a = c.Param("handson-boy")
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "sucessful action.", "data": a})

}

func put(c *gin.Context) {
	c.JSON(200, gin.H{"message": "put"})
}

func post(c *gin.Context) {
	c.JSON(200, gin.H{"message": "post"})
}

func delete(c *gin.Context) {
	c.JSON(200, gin.H{"status": 200, "message": "delete"})
}
