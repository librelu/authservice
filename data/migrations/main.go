package main

import (
	"github.com/authsvc/config"
	"github.com/authsvc/data/postgres"
	"github.com/authsvc/models"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		panic("can't init config")
	}

	postgresClient, err := postgres.NewDBClient(
		c.GetString("db.host", ""),
		c.GetString("db.port", ""),
		c.GetString("db.user", ""),
		c.GetString("db.password", ""),
		c.GetString("db.database", ""),
		c.GetString("db.sslmode", "false"),
	)
	if err != nil {
		panic(err)
	}
	defer postgresClient.DB.Close()

	postgresClient.DB.AutoMigrate(&models.User{})
}
