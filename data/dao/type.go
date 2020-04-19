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
	GetUserByUsername(username string) (output *models.User, err error)
	CreateUser(user, email string, password []byte) (output bool, err error)
	CreateCoupon(name string) error
	GetCouponByName(name string) (output *models.Coupon, err error)
	GiveCouponToUser(coupon *models.Coupon, user *models.User) error
}
