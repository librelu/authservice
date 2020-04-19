package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewJSONEndpoint Create an endpoint that handle JSON request and response
func NewJSONEndpoint(engine *gin.Engine, method string, url string, req interface{}, handler JSONHandler, handlers ...gin.HandlerFunc) {
	handlers = append(handlers, func(c *gin.Context) {
		if err := c.BindJSON(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("convert to struct error: %v", err)})
			panic(fmt.Sprintf("%+v\n", err))
		}
		r, err := handler(c, req)
		if r != nil {
			resultString, err := json.Marshal(r)
			if err != nil {
				c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
				panic(fmt.Sprintf("%+v\n", err))
			}
			resp := make(map[string]interface{})
			json.Unmarshal(resultString, &resp)
			c.JSON(http.StatusAccepted, resp)
			return
		} else if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			panic(fmt.Sprintf("%+v\n", err))
		}
	})
	engine.Handle(method, url, handlers...)
}
