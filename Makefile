migration:
	go run $(GOPATH)/src/github.com/authsvc/data/migrations/main.go
up:
	docker-compose build && docker-compose up