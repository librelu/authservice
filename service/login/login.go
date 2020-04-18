package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler user login process
func Handler() (h gin.HandlerFunc) {
	return func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{
			"message": "pong",
		})
	}
}
