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

// UserCoupons coupon user index table
type UserCoupons struct {
	UUID     sql.NullString `gorm:"not null"`
	CouponID sql.NullString `gorm:"not null",gorm:"foreignkey:UUID"`
	UserID   sql.NullString `gorm:"not null",gorm:"foreignkey:UUID"`
	CreateAt sql.NullTime   `gorm:"not null"`
	UpdateAt sql.NullTime   `gorm:"not null"`
}

// Coupon coupon info
type Coupon struct {
	UUID     sql.NullString `gorm:"not null"`
	Name     sql.NullString `gorm:"not null"`
	CreateAt sql.NullTime   `gorm:"not null"`
	UpdateAt sql.NullTime   `gorm:"not null"`
}
