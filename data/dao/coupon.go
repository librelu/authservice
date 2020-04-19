package dao

import (
	"database/sql"
	"time"

	"github.com/authsvc/models"
	"github.com/google/uuid"
)

// CreateCoupon create coupon return nil error if success
func (d *dao) CreateCoupon(name string) error {
	return nil
}

// GetCouponByName input a uni name to get coupon
func (d *dao) GetCouponByName(name string) (output *models.Coupon, err error) {
	coupon := &models.Coupon{}
	if db := d.postgres.DB.Find(coupon, models.Coupon{Name: sql.NullString{
		String: name,
		Valid:  true,
	}}); db.Error != nil {
		if !db.RecordNotFound() {
			return nil, db.Error
		}
	}
	return coupon, nil
}

// GiveCouponToUser give one coupon to user
func (d *dao) GiveCouponToUser(coupon *models.Coupon, user *models.User) error {
	timeNow := time.Now()
	userCoupons := &models.UserCoupons{
		UUID: sql.NullString{
			String: uuid.New().String(),
			Valid:  true,
		},
		CouponID: sql.NullString{
			String: coupon.UUID.String,
			Valid:  true,
		},
		UserID: sql.NullString{
			String: user.UUID.String,
			Valid:  true,
		},
		CreateAt: sql.NullTime{
			Time:  timeNow,
			Valid: true,
		},
		UpdateAt: sql.NullTime{
			Time:  timeNow,
			Valid: true,
		},
	}

	if db := d.postgres.DB.Create(userCoupons); db.Error != nil {
		return db.Error
	}
	return nil
}
