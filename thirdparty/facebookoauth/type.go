package facebookoauth

import (
	"context"

	"golang.org/x/oauth2"
)

const authTokenURL = "https://graph.facebook.com/me?fields=id,name,email&access_token="

// Handler Oauth2 handler
type Handler interface {
	GetOauthURL() string
	GetUserProfileByCode(ctx context.Context, token string) (userinfo *UserInfo, err error)
}

type client struct {
	config *oauth2.Config
}

//UserInfo userinfo description form facebook
type UserInfo struct {
	Email string
	ID    string
}
