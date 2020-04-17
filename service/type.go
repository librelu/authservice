package service

import "github.com/gin-gonic/gin"

// Endpoint general endpoint description
type Endpoint struct {
	Method   string
	URL      string
	Handler func(c *gin.Context)
}

// Endpoints endpoint collection
type Endpoints []Endpoint
