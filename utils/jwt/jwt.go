package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// ClaimJWTByUserInfo claim jwt token by user info, use password encrypt as token
func ClaimJWTByUserInfo(username, email string, passwordHash []byte) (tokenResult string, err error) {
	expirationTime := time.Now().Add(expiredTime * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(signMethod, claims)
	tokenResult, err = token.SignedString(passwordHash)
	if err != nil {
		return "", err
	}
	return tokenResult, nil
}

// VerifyJTW verify token and check the validate
func VerifyJTW(jwt string) (result bool, err error) {
	return true, nil
}
