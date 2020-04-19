package main

import (
	"database/sql"
	"time"

	"github.com/authsvc/config"
	"github.com/authsvc/data/postgres"
	"github.com/authsvc/models"
	"github.com/google/uuid"
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

	postgresClient.DB.AutoMigrate(&models.User{}, &models.UserCoupons{}, &models.Coupon{})
	// Create Init Coupon
	postgresClient.DB.Create(
		&models.Coupon{
			UUID: sql.NullString{
				String: uuid.New().String(),
				Valid:  true,
			},
			Name: sql.NullString{
				String: "WelcomeCoupon",
				Valid:  true,
			},
			CreateAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			UpdateAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		},
	)

}
