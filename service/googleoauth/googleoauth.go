package googleoauth

import (
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
			c.JSON(http.StatusBadGateway, gin.H{"error": errors.Errorf("can't get code from google")})
			return
		}

		unescapeCode, err := url.QueryUnescape(code)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": errors.Errorf("can't unescape code: %v", err)})
		}

		userinfo, err := googleoauth.GetUserProfileByToken(c, unescapeCode)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": errors.Errorf(err.Error()).Error()})
			return
		}

		username := userinfo.ID
		email := userinfo.Email
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(unescapeCode), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": errors.Errorf(err.Error()).Error()})
			return
		}

		if user, _ := daoHandler.GetUserByEmail(email); user != nil {
			token, err := jwt.ClaimJWTByUserInfo(username, email, passwordHash)
			if err != nil {
				c.JSON(http.StatusBadGateway, gin.H{"error": errors.Errorf(err.Error()).Error()})
			}
			c.JSON(http.StatusAccepted, gin.H{"token": token})
			return
		}

		if ok, err := daoHandler.CreateUser(username, email, passwordHash); !ok || err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": errors.Errorf(err.Error()).Error()})
			return
		}

		token, err := jwt.ClaimJWTByUserInfo(username, email, passwordHash)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": errors.Errorf(err.Error()).Error()})
			return
		}

		// Give Coupon to User
		coupon, err := daoHandler.GetCouponByName("WelcomeCoupon")
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": errors.Errorf(err.Error()).Error()})
			return
		}

		user, err := daoHandler.GetUserByEmail(email)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": errors.Errorf(err.Error()).Error()})
			return
		}

		if err := daoHandler.GiveCouponToUser(coupon, user); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": errors.Errorf(err.Error()).Error()})
			return
		}

		err = smtpHandler.SendWelcomeEmail(email, username, coupon.Name.String)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": errors.Errorf(err.Error()).Error()})
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
