package jwt

import "github.com/dgrijalva/jwt-go"

// default expired time is 5 min
const (
	expiredTime = 5
)

var (
	signMethod = jwt.SigningMethodHS256
)

// Claims jwt token message contianed
type Claims struct {
	Username string
	jwt.StandardClaims
}
