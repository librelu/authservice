package googleoauth

import (
	"context"

	"golang.org/x/oauth2"
)

const authTokenURL = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

// Handler Oauth2 handler
type Handler interface {
	GetOauthURL() string
	GetUserProfileByToken(ctx context.Context, token string) (userinfo *UserInfo, err error)
}

type client struct {
	config *oauth2.Config
}

//UserInfo userinfo description form google
type UserInfo struct {
	Email string
	ID    string
}
