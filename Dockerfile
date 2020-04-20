FROM golang:1.13

RUN apt-get update || exit 0
RUN apt-get upgrade -y
RUN mkdir -p /go/src/github.com/authsvc
COPY . /go/src/github.com/authsvc
WORKDIR /go/src/github.com/authsvc
RUN  go mod tidy
RUN go test ./...
RUN go install github.com/authsvc
EXPOSE 8080