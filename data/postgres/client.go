package postgres

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// NewDBClient new db client by gorm
func NewDBClient(host, port, user, password, database, sslmode string) (client *Client, err error) {
	config := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, user, database, password, sslmode,
	)
	db, err := gorm.Open(postgres, config)
	if err != nil {
		return nil, err
	}
	return &Client{
		DB: db,
	}, err
}
