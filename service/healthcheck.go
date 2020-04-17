package service

import "github.com/gin-gonic/gin"

// HealthCheckHandler check server is alives
func HealthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
