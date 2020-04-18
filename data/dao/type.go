package dao

import (
	"github.com/authsvc/data/postgres"
	"github.com/authsvc/models"
)

type dao struct {
	postgres *postgres.Client
}

// Handler dao handler to access database
type Handler interface {
	GetUser(input *models.User) (output *models.User, err error)
	CreateUser(user, email string, password []byte) (output bool, err error)
}
