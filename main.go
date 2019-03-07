package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/vrischmann/envconfig"

	"github.com/mijn-app/mijn-app-backend/app"
	"github.com/mijn-app/mijn-app-backend/utils/log"
	"github.com/mijn-app/mijn-app-backend/utils/sendinblue"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/joho/godotenv"
	"github.com/mijn-app/mijn-app-backend/utils/goalogadapter"
)

var (
	// ErrUnauthorized is the error returned for unauthorized requests.
	ErrUnauthorized = goa.NewErrorClass("unauthorized", 401)
	// BadRequestCustom custom bad request with custom message
	BadRequestCustom = goa.NewErrorClass("BadRequest", 400)
	ErrConflict      = goa.NewErrorClass("Conflict", 409)
	conf             Config
)

//Config stores the config to run the energy zero backend
type Config struct {
	App struct {
		// APP_URL is the scheme, hostname (+ port for dev purposes when e-mailing)
		URL             string `envconfig:"default=http://localhost:8081"`
		EnvironmentName string `envconfig:"ENVIRONMENT_NAME"`
	}
	NLX struct {
		URL                         string `envconfig:"NLX_OUTWAY_URL"`
		SolviteersContractsApiPath  string `envconfig:"SOLVITEERS_CONTRACTS_API_PATH"`
		SolviteersAVGLoggingApiPath string `envconfig:"SOLVITEERS_AVG_LOGGING_API_PATH"`
		SolviteersKey               string `envconfig:"SOLVITEERS_KEY"`
	}
	SendInBlue struct {
		SendInBlueAPIToken string `envconfig:"SEND_IN_BLUE_API_TOKEN"`
	}
	ItsMeOAuth struct {
		ClientID         string `envconfig:"OAUTH_ITSME_CLIENT_ID"`
		RedirectUrl      string `envconfig:"OAUTH_REDIRECT_URL"`
		LoginServiceCode string `envconfig:"OAUTH_ITSME_LOGIN_SERVICE_CODE"`
		DiscoveryUrl     string `envconfig:"OAUTH_ITSME_DISCOVERY_URL"`
	}
}

func init() {
	_, err := os.Stat("../variables.env")
	if !os.IsNotExist(err) {
		err := godotenv.Load("../variables.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	if err := envconfig.InitWithOptions(&conf, envconfig.Options{AllOptional: true}); err != nil {
		log.Fatal(err)
	}

	// In live environments all environment variables are mandatory
	if conf.App.EnvironmentName == "staging" || conf.App.EnvironmentName == "production" {
		if err := envconfig.InitWithOptions(&conf, envconfig.Options{AllOptional: false}); err != nil {
			log.Fatal("Environment variable check failed, error: ", err.Error())
		}
	}

	initDatabase(true, false)

	sendinblue.Init(conf.SendInBlue.SendInBlueAPIToken)

}

func main() {

	// Create service
	service := goa.New("Mijn-app")
	logger := log.GetLogger()

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	service.Use(jsMiddleware)

	jwtMiddleware, err := NewJWTMiddleware()
	if err != nil {
		log.Fatal(err)
	}
	app.UseJWTMiddleware(service, jwtMiddleware)
	app.UseSigninBasicAuthMiddleware(service, NewBasicAuthMiddleware())

	adapter := goalogadapter.New(logger)
	service.WithLogger(adapter)

	// Mount "JWT" controller
	c0 := NewJWTController(service)
	app.MountJWTController(service, c0)

	// Mount "Health" controller
	c1 := NewHealthController(service)
	app.MountHealthController(service, c1)

	// Mount "User" controller
	c2 := NewUserController(service)
	app.MountUserController(service, c2)

	// Mount "Contract" controller
	c3 := NewContractController(service)
	app.MountContractController(service, c3)

	// Mount "Avglog" controller
	c4 := NewAvglogController(service)
	app.MountAvglogController(service, c4)

	// Mount "Mail" controller
	c5 := NewMailController(service)
	app.MountMailController(service, c5)

	c6 := NewAddressController(service)
	app.MountAddressController(service, c6)

	c7 := NewOrderController(service)
	app.MountOrderController(service, c7)

	c8 := NewOauthController(service)
	app.MountOauthController(service, c8)

	// Start service
	if err := service.ListenAndServe(":80"); err != nil {
		service.LogError("startup", "err", err)
	}

}

func jsMiddleware(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		reqURL := string(req.URL.Path)
		if strings.Contains(reqURL, "js") {
			req.Header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
			req.Header.Set("Pragma", "no-cache")
			req.Header.Set("Pragma", "0")
			a := rw.Header()
			a.Set("Cache-Control", "no-cache, no-store, must-revalidate")
			a.Set("Pragma", "no-cache")
			a.Set("Expires", "0")
		}
		return h(ctx, rw, req)
	}
}

// exitOnFailure prints a fatal error message and exits the process with status 1.
func exitOnFailure(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "[CRIT] %s", err.Error())
	os.Exit(1)
}
