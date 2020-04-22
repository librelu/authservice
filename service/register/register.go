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
func Handler(daoHandler dao.Handler, smtpHandler smtp.Handler) (jsonHandler endpoints.JSONHandler) {
	return func(c *gin.Context, req interface{}) (resp interface{}, err error) {
		r, ok := req.(*Request)
		if !ok {
			return nil, errors.Errorf("can't convert request to struct")
		}

		password := r.Password
		if err := validatePassword(password); err != nil {
			return nil, err
		}

		email := r.Email
		if err := validateEmail(email); err != nil {
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

		if user, _ := daoHandler.GetUserByUsername(username); user != nil {
			return Response{
				message: "user already register",
			}, nil
		}

		if ok, err := daoHandler.CreateUser(username, email, passwordHash); !ok || err != nil {
			return nil, errors.Errorf("create user failed, error: %s", err)
		}

		// Give Coupon to User
		coupon, err := daoHandler.GetCouponByName("WelcomeCoupon")
		if err != nil {
			return nil, errors.Errorf("can't get coupon: %s", err)
		}

		user, err := daoHandler.GetUserByUsername(username)
		if err != nil {
			return nil, errors.Errorf("can't get user: %s", err)
		}

		if err := daoHandler.GiveCouponToUser(coupon, user); err != nil {
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

func validateEmail(email string) (err error) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(email) {
		return errors.Errorf("incorrect email ex: 1234@domain.com current input:%s", email)
	}
	return nil
}

func validatePassword(password string) (err error) {
	if len(password) < 8 {
		return errors.Errorf("the password should contains at lease 8 charactor")
	}
	return nil
}
