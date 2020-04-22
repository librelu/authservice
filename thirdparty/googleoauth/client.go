package googleoauth

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// NewClient new an oauth client
func NewClient(clientID, clientSecret, redirectURL string, scopes []string) Handler {
	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
		Endpoint:     google.Endpoint,
	}
	return &client{
		config: config,
	}
}

// GetOauthURL getting oauth url for processing google oauth
func (c *client) GetOauthURL() string {
	return c.config.AuthCodeURL("state")
}

func (c *client) GetUserProfileByToken(ctx context.Context, code string) (userinfo *UserInfo, err error) {
	tok, err := c.config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	log.Println("debug code:", code)
	log.Println("debug tok:", tok)
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
	log.Println("debug resp:", string(v))

	err = json.Unmarshal(v, userinfo)
	if err != nil {
		return nil, err
	}
	return userinfo, nil
}
