package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mijn-app/mijn-app-backend/app"
	"github.com/mijn-app/mijn-app-backend/models"

	"net/http"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	uuid "github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	jwtutils "github.com/mijn-app/mijn-app-backend/utils/jwt"
	"github.com/mijn-app/mijn-app-backend/utils/log"
	"golang.org/x/crypto/bcrypt"
)

type AppState struct {
	models.User
}

/*
	Middleware functions
*/

// NewJWTMiddleware creates a middleware that checks for the presence of a JWT Authorization header
// and validates its content. A real app would probably use goa's JWT security middleware instead.
func NewJWTMiddleware() (goa.Middleware, error) {
	err := jwtutils.LoadKeys()
	if err != nil {
		panic(err)
	}
	pub := jwtutils.PublicKey()
	return jwt.New(pub, jwtutils.ForceFail(), app.NewJWTSecurity()), nil
}

// NewBasicAuthMiddleware creates a middleware that checks for the presence of a basic auth header
// and validates its content.
func NewBasicAuthMiddleware() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Retrieve and log basic auth info
			user, pass, ok := req.BasicAuth()
			if !ok || user == "" || pass == "" {
				log.Warningf("Failed basic auth, user: %s", user)
				return ErrUnauthorized("Failed basic authentication")
			}

			// Normal basic auth
			var account models.User
			err := userDB.Db.Scopes().Table(userDB.TableName()).Where("email = ? and password IS NOT NULL and password != ''", user).Find(&account).Error

			if err != nil || err == gorm.ErrRecordNotFound {
				log.Warningf("User trying to sign in not found in DB, error: %s", err.Error())
				return ErrUnauthorized("Authorization failed")
			}

			err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(pass))
			if err != nil {
				return ErrUnauthorized("Authorization failed")
			}

			// Proceed
			rw.Header().Set("Access-Control-Expose-Headers", "Authorization")
			return h(ctx, rw, req)
		}
	}
}

func createJWTToken(uniqueID *uuid.UUID, inXMinutes int64, userID uuid.UUID, scope interface{}, platform string) string {
	token := jwtgo.New(jwtgo.SigningMethodRS512)
	var insert bool
	var uniqueJWTID uuid.UUID
	if uniqueID != nil {
		uniqueJWTID = *uniqueID
	} else {
		uniqueJWTID = uuid.Must(uuid.NewV4())
		insert = true
	}

	token.Claims = jwtgo.MapClaims{
		"iss":    "System",          // who creates the token and signs it
		"aud":    "User",            // to whom the token is intended to be sent
		"exp":    inXMinutes,        // time when the token will expire (x minutes from now)
		"jti":    uniqueJWTID,       // a unique identifier for the token
		"iat":    time.Now().Unix(), // when the token was issued/created (now)
		"nbf":    2,                 // time before which the token is not yet valid (2 minutes ago)
		"sub":    userID,            // the subject/principal is whom the token is about
		"scopes": scope,             // token scope - not a standard claim
	}

	signedToken, err := token.SignedString(jwtutils.PrivateKey())
	if err != nil {
		log.Errorf("Failed to sign token, error: %s", err.Error())
	}

	// Create JWTModel for DB
	var jwtmodel models.JWT
	jwtmodel.UserID = userID
	jwtmodel.UniqueID = uniqueJWTID.String()

	if insert {
		jwtmodel.Platform = platform

		err = jwtDB.Db.Create(&jwtmodel).Error
		if err != nil {
			log.Errorf("Error creating jwt model: %v", err.Error())
		}
	} else {
		var obj models.JWT
		err := jwtDB.Db.Scopes().Table(jwtDB.TableName()).Where("unique_id = ?", uniqueJWTID).Find(&obj).Error
		if err != nil {
			return "JWT not found in DB"
		}
		err = jwtDB.Db.Model(obj).Updates(jwtmodel).Error
	}
	return signedToken
}

func isDashboard(scopes []string) bool {
	return jwtutils.HasScope("api:DASHBOARD", scopes)
}

func isAdmin(scopes []string) bool {
	return jwtutils.HasScope("api:ADMIN", scopes)
}

func isUser(scopes []string) bool {
	return jwtutils.HasScope("api:USER", scopes)
}

// JWT handler, check if token is in DB, can be expired
func getTokenFromRequest(ctx context.Context, reqData *http.Request) (*jwtgo.Token, error) {
	bearer := reqData.Header.Get("X-auth")
	if bearer == "" || bearer == "undefined" {
		goa.LogError(ctx, "No JWT token provided")
		return nil, errors.New("No JWT token provided")
	}
	token, err := jwtgo.ParseWithClaims(
		bearer[7:],
		jwtgo.MapClaims{},
		func(token *jwtgo.Token) (interface{}, error) {
			// Make sure token's signature wasn't changed
			if _, ok := token.Method.(*jwtgo.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected siging method")
			}
			// Unpack key from PEM encoded PKCS8
			return jwtutils.PublicKey(), nil
		},
	)
	if err != nil {
		goa.LogError(ctx, "Token provided not valid", "error", err.Error())
		return nil, errors.New("Token not valid")
	}
	return token, nil
}

// JWT handler, check if token is in DB, can be expired
func getAccountIDAndScopesFromJWTNoSecurity(reqData *goa.RequestData) (uuid.UUID, []string, error) {
	bearer := reqData.Request.Header.Get("X-auth")
	var userID uuid.UUID
	if bearer == "" || bearer == "undefined" {
		return userID, nil, errors.New("No bearer specified")
	}
	token, err := jwtgo.ParseWithClaims(
		bearer[7:],
		jwtgo.MapClaims{},
		func(token *jwtgo.Token) (interface{}, error) {
			// Make sure token's signature wasn't changed
			if _, ok := token.Method.(*jwtgo.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected siging method")
			}
			// Unpack key from PEM encoded PKCS8
			return jwtutils.PublicKey(), nil
		},
	)
	if err != nil {
		log.Error("Did not receive token claims")
		return userID, nil, fmt.Errorf("Did not receive token claims: %s", err) // internal error
	}
	// Check if jwt is in DB
	claims := token.Claims.(jwtgo.MapClaims)

	var jwtfromdb models.JWT
	errValidationFailed := goa.NewErrorClass("validation_failed", 401)
	errDB := jwtDB.Db.Scopes().Table(jwtDB.TableName()).
		Where("unique_id = ?", claims["jti"]).
		Find(&jwtfromdb).Error
	if errDB != nil {
		return userID, nil, errValidationFailed("forcing failure for mobile session not in DB anymore")
	}
	userID, scopes := getIDAndScopesFromClaims(claims)
	return userID, scopes, err
}

func getAccountIDAndScopesFromJWT(ctx context.Context) (uuid.UUID, []string, error) {
	errValidationFailed := goa.NewErrorClass("validation_failed", 401)
	token := jwt.ContextJWT(ctx)
	if token == nil {
		return uuid.Must(uuid.NewV4()), nil, errValidationFailed("forcing failure because token is missing")
	}
	userID, scopes := getIDAndScopesFromClaims(token.Claims.(jwtgo.MapClaims))
	return userID, scopes, nil
}

func getIDAndScopesFromClaims(claims jwtgo.MapClaims) (uuid.UUID, []string) {
	var scopes []interface{}
	var stringScopes []string
	id := claims["sub"].(string)
	i, err := uuid.FromString(id)
	if err != nil {
		log.Errorf("Could not convert user id to UUID: %s", id)
	}
	scopes = claims["scopes"].([]interface{})
	for _, scope := range scopes {
		stringScopes = append(stringScopes, scope.(string))
	}
	return i, stringScopes
}

func getScopes(ctx context.Context, scope models.UserRoles) []string {
	var scopes []string
	for key, val := range scope.AllStrings() {
		if scope, _ := scope.Value(); key <= scope.(int64) {
			scopes = append(scopes, "api:"+val)
		}
	}
	return scopes
}

//Fills the appstate for given user id
func fillAppStateForUserID(userID uuid.UUID) (*app.Appstate, error) {
	user, err := userDB.OneUserOnID(userID)
	if err != nil {
		log.Errorf("Error retrieving user: %s for appstate, error: %s", userID.String(), err.Error())
		return nil, err
	}
	return fillAppStateForUser(*user)
}

func fillAppStateForUser(user models.User) (*app.Appstate, error) {
	var appState AppState
	appState.User = user
	return appState.GetAppStateMedia(), nil
}

//GetAppStateMedia get app state media type
func (a *AppState) GetAppStateMedia() *app.Appstate {
	var appState app.Appstate
	appState.User = a.User.UserToUser()
	appState.User.Address = a.User.Addres.AddressToAddress()
	return &appState
}
