package service

// Endpoint general endpoint description
type Endpoint struct {
	Method   string
	URL      string
	Handler  interface{}
	Request  interface{}
	Response interface{}
}

// Endpoints endpoint collection
type Endpoints []Endpoint
