package facebookoauth

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

// NewClient new an oauth client
func NewClient(clientID, clientSecret, redirectURL string, scopes []string) Handler {
	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
		Endpoint:     facebook.Endpoint,
	}
	return &client{
		config: config,
	}
}

// GetOauthURL getting oauth url for processing facebook oauth
func (c *client) GetOauthURL() string {
	return c.config.AuthCodeURL("state")
}

func (c *client) GetUserProfileByCode(ctx context.Context, code string) (userinfo *UserInfo, err error) {
	tok, err := c.config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	client := c.config.Client(ctx, tok)
	resp, err := client.Get(authTokenURL + tok.AccessToken)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	v, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	userinfo = new(UserInfo)
	log.Println(string(v))
	err = json.Unmarshal(v, userinfo)
	if err != nil {
		return nil, err
	}
	return userinfo, nil
}
