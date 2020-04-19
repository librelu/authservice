package register

import (
	"regexp"

	"github.com/authsvc/data/dao"
	"github.com/authsvc/thirdparty/smtp"
	"github.com/authsvc/utils/endpoints"
	"github.com/authsvc/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// Handler user register process
func Handler(dbHandler dao.Handler, smtpHandler smtp.Handler) (jsonHandler endpoints.JSONHandler) {
	return func(c *gin.Context, req interface{}) (resp interface{}, err error) {
		r, ok := req.(*Request)
		if !ok {
			return nil, errors.Errorf("can't convert request to struct")
		}

		password := r.Password
		if ok, err := validatePassword(password); !ok || err != nil {
			return nil, err
		}

		email := r.Email
		if ok, err := validateEmail(email); !ok || err != nil {
			return nil, err
		}

		username := r.Username
		if len(username) == 0 {
			return nil, errors.Errorf("Username is require in request body")
		}

		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, errors.Errorf("crypt password failed, error: %s", err)
		}

		if ok, err := dbHandler.CreateUser(username, email, passwordHash); !ok || err != nil {
			return nil, errors.Errorf("create user failed, error: %s", err)
		}

		// Give Coupon to User
		coupon, err := dbHandler.GetCouponByName("WelcomeCoupon")
		if err != nil {
			return nil, errors.Errorf("can't get coupon: %s", err)
		}

		user, err := dbHandler.GetUserByUsername(username)
		if err != nil {
			return nil, errors.Errorf("can't get user: %s", err)
		}

		if err := dbHandler.GiveCouponToUser(coupon, user); err != nil {
			return nil, errors.Errorf("can't give coupon: %s", err)
		}

		token, err := jwt.ClaimJWTByUserInfo(username, email, passwordHash)
		if err != nil {
			return nil, errors.Errorf("can't get token")
		}

		err = smtpHandler.SendWelcomeEmail(email, username, coupon.Name.String)
		if err != nil {
			return nil, errors.Errorf("failed to send welcome email: %s", err)
		}

		return Response{
			Token: token,
		}, nil
	}
}

func validateEmail(email string) (result bool, err error) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(email) {
		return false, errors.Errorf("incorrect email ex: 1234@domain.com current input:%s", email)
	}
	return true, nil
}

func validatePassword(password string) (result bool, err error) {
	if len(password) < 8 {
		return false, errors.Errorf("the password should contains at lease 8 charactor")
	}
	return true, nil
}
