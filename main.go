package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/login/service"
)

func main() {
	endpoints := service.Endpoints{
		service.Endpoint{
			Method:  http.MethodGet,
			URL:     "/ping",
			Handler: service.HealthCheckHandler,
		},
	}

	r := gin.Default()
	for _, e := range endpoints {
		r.Handle(e.Method, e.URL, e.Handler)
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
