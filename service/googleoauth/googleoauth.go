package googleoauth

import (
	"log"
	"net/http"
	"net/url"

	"github.com/authsvc/data/dao"
	"github.com/authsvc/thirdparty/googleoauth"
	"github.com/authsvc/thirdparty/smtp"
	"github.com/authsvc/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// Handler get resp after credential, and create a user if existed
func Handler(googleoauth googleoauth.Handler, daoHandler dao.Handler, smtpHandler smtp.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		code, ok := c.GetQuery("code")
		if !ok {
			err := errors.Errorf("can't get code from google")
			responseError(c, err)
			return
		}

		unescapeCode, err := url.QueryUnescape(code)
		if err != nil {
			err := errors.Errorf("can't unescape code: %v", err)
			responseError(c, err)
			return
		}

		userinfo, err := googleoauth.GetUserProfileByCode(c, unescapeCode)
		if err != nil {
			err := errors.Errorf("can't get user profie code: %v", err)
			responseError(c, err)
			return
		}

		username := userinfo.ID
		email := userinfo.Email
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(unescapeCode), bcrypt.DefaultCost)
		if err != nil {
			err := errors.Errorf("can't get get password hash: %v", err)
			responseError(c, err)
			return
		}

		if user, _ := daoHandler.GetUserByEmail(email); user != nil {
			token, err := jwt.ClaimJWTByUserInfo(username, email, passwordHash)
			if err != nil {
				err := errors.Errorf("can't claim JWT: %v", err)
				responseError(c, err)
				return
			}
			c.JSON(http.StatusAccepted, gin.H{"token": token})
			return
		}

		if ok, err := daoHandler.CreateUser(username, email, passwordHash); !ok || err != nil {
			err := errors.Errorf("can't create user: %v", err)
			responseError(c, err)
			return
		}

		token, err := jwt.ClaimJWTByUserInfo(username, email, passwordHash)
		if err != nil {
			err := errors.Errorf("can't claim JWT: %v", err)
			responseError(c, err)
			return
		}

		// Give Coupon to User
		coupon, err := daoHandler.GetCouponByName("WelcomeCoupon")
		if err != nil {
			err := errors.Errorf("can't get coupon: %v", err)
			responseError(c, err)
			return
		}

		user, err := daoHandler.GetUserByEmail(email)
		if err != nil {
			err := errors.Errorf("can't get user by email: %v", err)
			responseError(c, err)
			return
		}

		if err := daoHandler.GiveCouponToUser(coupon, user); err != nil {
			err := errors.Errorf("can't give coupon to user: %v", err)
			responseError(c, err)
			return
		}

		err = smtpHandler.SendWelcomeEmail(email, username, coupon.Name.String)
		if err != nil {
			err := errors.Errorf("can't send welcome email: %v", err)
			responseError(c, err)
			return
		}

		c.JSON(http.StatusAccepted, gin.H{"token": token})
	}
}

// GetURLHandler get oauth url
func GetURLHandler(googleoauth googleoauth.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{"url": googleoauth.GetOauthURL()})
	}
}

func responseError(c *gin.Context, err error) {
	log.Println(err)
	c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	return
}
