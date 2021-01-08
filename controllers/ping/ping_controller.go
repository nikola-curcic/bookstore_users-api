package ping

import (
	 "github.com/gin-gonic/gin"
	// "net/http"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
