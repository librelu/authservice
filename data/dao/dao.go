package dao

import (
	"github.com/authsvc/data/postgres"
)

// NewHandler new a dao handler
func NewHandler(pc *postgres.Client) (handler Handler, err error) {
	return &dao{
		postgres: pc,
	}, nil
}
