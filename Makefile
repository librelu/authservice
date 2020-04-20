migration:
	go run $(GOPATH)/src/github.com/authsvc/data/migrations/main.go
up:
	docker-compose build && docker-compose up
build:
	docker build -t authsvc .
deploy:
  heroku container:push authsvc -a amazingtalker && heroku container:release authsvc -a amazingtalker