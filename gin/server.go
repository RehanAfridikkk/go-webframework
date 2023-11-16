package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Text string `json:"text"`
	ID int 
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, gin World!"})
		
	})

	r.POST("/post", func(c *gin.Context) {
		var message Message
		if err := c.BindJSON(&message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"received": message.Text})
	})

	r.Run(":1336")
}
