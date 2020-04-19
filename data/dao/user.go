package dao

import (
	"database/sql"
	"time"

	"github.com/authsvc/models"
	"github.com/google/uuid"
)

func (d *dao) GetUser(input *models.User) (output *models.User, err error) {
	return nil, nil
}

func (d *dao) CreateUser(user, email string, password []byte) (output bool, err error) {
	uuid := uuid.New().String()
	timeNow := time.Now()
	// find user or create a new row
	db := d.postgres.DB.Where(
		models.User{
			Username: sql.NullString{
				String: user,
				Valid:  true,
			},
			Email: sql.NullString{
				String: email,
				Valid:  true,
			},
		},
	).First(new(models.User))

	if db.Error != nil {
		if !db.RecordNotFound() {
			return false, err
		}
	}

	if db.RecordNotFound() {
		if err := d.postgres.DB.Create(models.User{
			UUID: sql.NullString{
				String: uuid,
				Valid:  true,
			},
			Username: sql.NullString{
				String: user,
				Valid:  true,
			},
			Email: sql.NullString{
				String: email,
				Valid:  true,
			},
			Password: password,
			CreateAt: sql.NullTime{
				Time:  timeNow,
				Valid: true,
			},
			UpdateAt: sql.NullTime{
				Time:  timeNow,
				Valid: true,
			},
		}).Error; err != nil {
			return false, err
		}
	}
	return true, nil
}
