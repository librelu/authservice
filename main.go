package main

import (
	"log"
	"net/http"

	"github.com/authsvc/config"
	"github.com/authsvc/data/dao"
	"github.com/authsvc/data/postgres"
	"github.com/authsvc/service"
	goauthHandler "github.com/authsvc/service/googleoauth"
	"github.com/authsvc/service/healthcheck"
	"github.com/authsvc/service/login"
	"github.com/authsvc/service/register"
	"github.com/authsvc/thirdparty/googleoauth"
	"github.com/authsvc/thirdparty/smtp"
	endpointutils "github.com/authsvc/utils/endpoints"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set up log status
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Init config
	c, err := config.NewConfig()
	if err != nil {
		panic("can't init config")
	}

	// Init postgres
	postgresClient, err := postgres.NewDBClient(
		c.GetString("db.host", ""),
		c.GetString("db.port", ""),
		c.GetString("db.user", ""),
		c.GetString("db.password", ""),
		c.GetString("db.database", ""),
		c.GetString("db.sslmode", "false"),
	)
	if err != nil {
		panic(err)
	}
	defer postgresClient.DB.Close()

	// Init Google Oauth
	goauthClient := googleoauth.NewClient(
		c.GetString("google_oauth.client_id", ""),
		c.GetString("google_oauth.client_secret", ""),
		c.GetString("google_oauth.redirect_url", ""),
		c.GetStringList("google_oauth.scopes", nil),
	)

	// Init dao
	daoHandler, err := dao.NewHandler(postgresClient)
	if err != nil {
		panic(err)
	}
	// Init SMTP
	smtpHandler := smtp.NewSMTPClient(
		c.GetString("smtp.identity", ""),
		c.GetString("smtp.username", ""),
		c.GetString("smtp.password", ""),
		c.GetString("smtp.host", ""),
		c.GetString("smtp.port", ""),
	)

	// Init endpotins
	endpoints := service.Endpoints{
		service.Endpoint{
			Method:  http.MethodGet,
			URL:     "/ping",
			Handler: healthcheck.Handler(),
		},
		service.Endpoint{
			Method:  http.MethodPost,
			URL:     "/login",
			Handler: login.Handler(),
		},
		service.Endpoint{
			Method:  http.MethodPost,
			URL:     "/register",
			Handler: register.Handler(daoHandler, smtpHandler),
			Request: new(register.Request),
		},
		service.Endpoint{
			Method:  http.MethodGet,
			URL:     "/googleoauth",
			Handler: goauthHandler.Handler(goauthClient, daoHandler, smtpHandler),
		},
		service.Endpoint{
			Method:  http.MethodGet,
			URL:     "/get_googleoauth_url",
			Handler: goauthHandler.GetURLHandler(goauthClient),
		}
	}
	r := gin.Default()
	for _, e := range endpoints {
		switch e.Handler.(type) {
		case endpointutils.JSONHandler:
			if h, ok := e.Handler.(endpointutils.JSONHandler); ok {
				endpointutils.NewJSONEndpoint(r, e.Method, e.URL, e.Request, h)
			}
		default:
			if handler, ok := e.Handler.(gin.HandlerFunc); ok {
				r.Handle(e.Method, e.URL, handler)
			} else {
				log.Printf("The endpoint: %s, method: %s, is skipping init process", e.URL, e.Method)
			}
		}
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
