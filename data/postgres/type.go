package postgres

import "github.com/jinzhu/gorm"

const postgres = "postgres"

// Client client handler by gorm
type Client struct {
	DB *gorm.DB
}
