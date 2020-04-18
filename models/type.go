package models

import (
	"database/sql"
)

//User user model refer to db table
type User struct {
	UUID     sql.NullString `gorm:"not null"`
	Username sql.NullString `gorm:"not null"`
	Email    sql.NullString `gorm:"not null"`
	Password []byte         `gorm:"not null"`
	CreateAt sql.NullTime   `gorm:"not null"`
	UpdateAt sql.NullTime   `gorm:"not null"`
}
