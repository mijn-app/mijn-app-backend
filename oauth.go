package main

import (
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"golang.org/x/oauth2"
	"github.com/mijn-app/mijn-app-backend/app"
	"github.com/mijn-app/mijn-app-backend/utils/oauth"
	"net/url"
)

// OauthController implements the oauth resource.
type OauthController struct {
	*goa.Controller
}

// NewOauthController creates a oauth controller.
func NewOauthController(service *goa.Service) *OauthController {
	return &OauthController{Controller: service.NewController("OauthController")}
}

// Callback runs the callback action.
func (c *OauthController) Callback(ctx *app.CallbackOauthContext) error {
	// OauthController_Callback: start_implement

	var state oauth.State

	// Decode the url-encoded state from oauth
	rawState, err := url.QueryUnescape(ctx.State)
	if err != nil {
		goa.LogError(ctx, "unable to unescape oauth state", "error", err.Error())
		return buildErrorResponse(ctx)
	}

	// Unmarshal the state attached from the oauth action
	err = json.Unmarshal([]byte(rawState), &state)
	if err != nil {
		goa.LogError(ctx, "unable to unmarshal oauth state", "error", err.Error())
		return buildErrorResponse(ctx)
	}

	// At any failure point, redirect back to the app
	redirect := fmt.Sprintf("%s:/?server-error=true", conf.App.URL)

	// Check if the OAuthProvider is valid
	if !isOAuthProviderAllowed(state.OAuthProvider) {
		goa.LogError(ctx, "Invalid OAuth provider for OAuth init", "error", err.Error())
		return buildFoundResponse(ctx, redirect)
	}

	if ctx.ErrorCode != nil || ctx.Code == nil {
		return buildFoundResponse(ctx, redirect)
	}
	code := *ctx.Code

	redirect = fmt.Sprintf("%s/oauth-%s?code=%s&state=%s", conf.App.URL, string(state.OAuthProvider), code, ctx.State)

	return buildFoundResponse(ctx, redirect)

	// OauthController_Callback: end_implement
}

// Handle runs the handle action.
func (c *OauthController) Handle(ctx *app.HandleOauthContext) error {
	// OauthController_Handle: start_implement

	var provider oauth.Provider
	err := provider.ScanFromString(ctx.Provider)
	if err != nil || !isOAuthProviderAllowed(provider) {
		goa.LogError(ctx, "Invalid OAuth provider for OAuth handle", "error", err.Error())
		return ctx.BadRequest(goa.ErrBadRequest("Invalid OAuth provider"))
	}
	err = initOauthForProvider(provider)
	if err != nil {
		goa.LogError(ctx, "Could not init OAuth", "error", err.Error())
		return ctx.InternalServerError()
	}

	isTokenValid, err := oAuthStateDB.ValidateStateToken(ctx.StateToken)
	if err != nil {
		goa.LogError(ctx, "Error validating state token for OAuth handle", "error", err.Error())
		return ctx.InternalServerError()
	}
	if !isTokenValid {
		return ctx.Unauthorized()
	}

	config := getConfigForProvider(provider)

	clientAssertion, err := oauth.BuildItsMeClientAssertion()
	if err != nil {
		return ctx.InternalServerError()
	}

	token, err := config.Exchange(ctx, ctx.Code,
		oauth2.SetAuthURLParam("client_assertion_type", "urn:ietf:params:oauth:client-assertion-type:jwt-bearer"),
		oauth2.SetAuthURLParam("client_assertion", clientAssertion),
	)
	if err != nil {
		goa.LogError(ctx, "Error validating OAuth token for OAuth handle", "error", err.Error())
		return ctx.Unauthorized()
	}

	claims, err := oauth.VerifyItsMeIDToken(token)
	if err != nil {
		goa.LogError(ctx, "Error verifying ID token for OAuth handle", "error", err.Error())
		return ctx.Unauthorized()
	}

	itsMeUser, err := oauth.GetItsMeUser(token)

	if claims.Subject != itsMeUser.Subject {
		goa.LogError(ctx, "Subject claims differ between user info and token!")
		return ctx.Unauthorized()
	}

	user, err := userDB.GetUserByItsMeSubjectNumber(ctx, claims.Subject)
	if err == gorm.ErrRecordNotFound {
		err = addItsMeSubjectNumberToUser(ctx, itsMeUser.EID.GovernmentIdentifier, itsMeUser.Subject)
		user, err = userDB.GetUserByItsMeSubjectNumber(ctx, claims.Subject)
	}
	if err != nil {
		goa.LogError(ctx, "Error retrieving db user for OAuth handle", "error", err.Error())
	}

	signedToken, err := createHandleResponseObjects(ctx, *user)
	if err != nil {
		goa.LogError(ctx, "Error creating jwt token")
		return ctx.InternalServerError()
	}

	// Set auth header for client retrieval
	ctx.ResponseData.Header().Set("Authorization", "Bearer "+signedToken)
	ctx.ResponseData.Header().Set("Access-Control-Expose-Headers", "Authorization")

	appState := AppState{
		User: *user,
	}

	res := user.UserToUser()
	res.Address = user.Addres.AddressToAddress()
	return ctx.OK(appState.GetAppStateMedia())

	// OauthController_Handle: end_implement
}

// Init runs the init action.
func (c *OauthController) Init(ctx *app.InitOauthContext) error {
	// OauthController_Init: start_implement

	var provider oauth.Provider
	err := provider.ScanFromString(ctx.Provider)
	if err != nil || !isOAuthProviderAllowed(provider) {
		goa.LogError(ctx, "Invalid OAuth provider for OAuth init", "error", err.Error())
		return ctx.BadRequest(goa.ErrBadRequest("Invalid OAuth provider"))
	}
	err = initOauthForProvider(provider)
	if err != nil {
		goa.LogError(ctx, "Could not init OAuth", "error", err.Error())
		return ctx.InternalServerError()
	}

	stateToken, err := oAuthStateDB.GenerateStateToken(ctx)
	stateString, err := oauth.BuildStateString(stateToken, ctx.App, provider, ctx.AppName)
	if err != nil {
		goa.LogError(ctx, "Error creating state token for OAuth init", "error", err.Error())
		return ctx.InternalServerError()
	}

	switch provider {
	case oauth.ProviderItsMe:
		ctx.ResponseData.Header().Set("Location", oauth.GetItsMeLoginURL(stateString))
	}

	ctx.ResponseData.Header().Set("Access-Control-Expose-Headers", "Location")
	return ctx.NoContent()

	// OauthController_Init: end_implement
}
