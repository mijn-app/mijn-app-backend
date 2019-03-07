package main

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	uuid "github.com/gofrs/uuid"
	"github.com/mijn-app/mijn-app-backend/app"
	"github.com/mijn-app/mijn-app-backend/models"
	"time"
)

// JWTController implements the jwt resource.
type JWTController struct {
	*goa.Controller
}

// NewJWTController creates a jwt controller.
func NewJWTController(service *goa.Service) *JWTController {
	return &JWTController{Controller: service.NewController("JWTController")}
}

// Refresh runs the refresh action.
func (c *JWTController) Refresh(ctx *app.RefreshJWTContext) error {
	// JWTController_Refresh: start_implement

	uid, _, _ := getAccountIDAndScopesFromJWT(ctx)
	token := jwt.ContextJWT(ctx)
	if token == nil {
		return ErrUnauthorized("Missing JWT token")
	}
	// Use the claims to authorize
	claims := token.Claims.(jwtgo.MapClaims)
	var jwtfromdb models.JWT
	if err := jwtDB.Db.Scopes().Table(jwtDB.TableName()).Where("unique_id = ?", claims["jti"]).Find(&jwtfromdb).Error; err != nil {
		return ErrUnauthorized("Forcing failure: JWT not found in DB")
	}
	in1Month := time.Now().Add(time.Hour * 24 * 7 * time.Duration(4)).Unix()
	jwtUniqueID, _ := uuid.FromString(claims["jti"].(string))
	signedToken := createJWTToken(&jwtUniqueID, in1Month, uid, claims["scopes"], "")
	ctx.ResponseData.Header().Set("Authorization", "Bearer "+signedToken)
	ctx.ResponseData.Header().Set("Access-Control-Expose-Headers", "Authorization")
	//Fill app state for user
	appState, err := fillAppStateForUserID(uid)
	if err != nil {
		return ErrUnauthorized(err.Error())
	}
	return ctx.OK(appState)

	// JWTController_Refresh: end_implement
}

// Signin runs the signin action.
func (c *JWTController) Signin(ctx *app.SigninJWTContext) error {
	// JWTController_Signin: start_implement

	platform := ctx.RequestData.Request.Header.Get("X-Platform")
	username, _, _ := ctx.RequestData.Request.BasicAuth()
	var account models.User
	err := userDB.Db.Scopes().Table(userDB.TableName()).Where("email = ?", username).Find(&account).Error

	if err != nil {
		return ctx.Unauthorized()
	}

	newScopes := getScopes(ctx, account.Role)

	// Generate auth token
	in1Month := time.Now().Add(time.Hour * 24 * 7 * time.Duration(4)).Unix()
	signedToken := createJWTToken(nil, in1Month, account.ID, newScopes, platform)
	// Set auth header for client retrieval
	ctx.ResponseData.Header().Set("Authorization", "Bearer "+signedToken)
	ctx.ResponseData.Header().Set("Access-Control-Expose-Headers", "Authorization")
	//Fill app state for user
	appState, err := fillAppStateForUser(account)
	if err != nil {
		return ErrUnauthorized(err.Error())
	}
	return ctx.OK(appState)

	// JWTController_Signin: end_implement
}
