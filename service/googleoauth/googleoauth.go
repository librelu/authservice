package googleoauth

import (
	"net/http"

	"github.com/authsvc/data/dao"
	"github.com/authsvc/thirdparty/googleoauth"
	"github.com/authsvc/thirdparty/smtp"
	"github.com/authsvc/utils/jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Handler get resp after credential, and create a user if existed
func Handler(googleoauth googleoauth.Handler, daoHandler dao.Handler, smtpHandler smtp.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		code, ok := c.GetQuery("code")
		if !ok {
			c.JSON(http.StatusBadGateway, gin.H{"error": "can't get code from google"})
			return
		}

		userinfo, err := googleoauth.GetUserProfileByToken(c, code)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		username := userinfo.Username
		email := userinfo.Email
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(code), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		if user, _ := daoHandler.GetUserByUsername(username); user != nil {
			token, err := jwt.ClaimJWTByUserInfo(username, email, passwordHash)
			if err != nil {
				c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			}
			c.JSON(http.StatusAccepted, gin.H{"token": token})
			return
		}

		if ok, err := daoHandler.CreateUser(username, email, passwordHash); !ok || err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		token, err := jwt.ClaimJWTByUserInfo(username, email, passwordHash)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		// Give Coupon to User
		coupon, err := daoHandler.GetCouponByName("WelcomeCoupon")
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		user, err := daoHandler.GetUserByUsername(username)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		if err := daoHandler.GiveCouponToUser(coupon, user); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		err = smtpHandler.SendWelcomeEmail(email, username, coupon.Name.String)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{"token": token})
	}
}
