package handlers

import (
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	greet := "pong"
	c.JSON(200, gin.H{
		"message": greet,
	})
	return
}
