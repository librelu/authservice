package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler check server is alives
func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{
			"message": "pong",
		})
	}
}
