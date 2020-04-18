package endpoints

import "github.com/gin-gonic/gin"

// JSONHandler generic JSON handler for all JSON endpoints
type JSONHandler func(c *gin.Context, req interface{}) (resp interface{}, err error)
