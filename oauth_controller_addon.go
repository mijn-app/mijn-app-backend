package main

import (
	"context"
	"golang.org/x/oauth2"
	"io/ioutil"
	"github.com/mijn-app/mijn-app-backend/app"
	"github.com/mijn-app/mijn-app-backend/models"
	"github.com/mijn-app/mijn-app-backend/utils/log"
	"github.com/mijn-app/mijn-app-backend/utils/oauth"
	"os"
	"time"
)

var allowedProviders = []oauth.Provider {oauth.ProviderItsMe}

// Initiate the ItsMe OAuth client if needed
func initItsMeAuth() error {
	privateSigningKey, err := os.Open("keys/itsme_private_signing.json")
	privateSigningKeyValue, err := ioutil.ReadAll(privateSigningKey)
	privateEncryptionKey, err := os.Open("keys/itsme_private_encryption.json")
	privateEncryptionKeyValue, err := ioutil.ReadAll(privateEncryptionKey)
	publicSigningKey, err := os.Open("keys/itsme_public_signing.json")
	publicSigningKeyValue, err := ioutil.ReadAll(publicSigningKey)
	publicEncryptionKey, err := os.Open("keys/itsme_public_encryption.json")
	publicEncryptionKeyValue, err := ioutil.ReadAll(publicEncryptionKey)
	if err != nil {
		log.Fatalf("ItsMe keys missing from environment, error: %s", err.Error())
	}

	return oauth.InitItsMeOAuth(
		conf.ItsMeOAuth.ClientID,
		conf.ItsMeOAuth.RedirectUrl,
		conf.ItsMeOAuth.LoginServiceCode,
		conf.ItsMeOAuth.DiscoveryUrl,
		&[]string{},
		string(privateSigningKeyValue),
		string(privateEncryptionKeyValue),
		string(publicSigningKeyValue),
		string(publicEncryptionKeyValue),
	)
}

func initOauthForProvider(provider oauth.Provider) error {
	switch provider {
	case oauth.ProviderItsMe:
		return initItsMeAuth()
	}
	return nil
}

func getConfigForProvider(provider oauth.Provider) *oauth2.Config {
	switch provider {
	case oauth.ProviderItsMe:
		return oauth.ItsMeOAuthConfiguration
	}
	return nil
}

func isOAuthProviderAllowed(providerToCheck oauth.Provider) bool {
	for _, allowedProvider := range allowedProviders {
		if providerToCheck == allowedProvider {
			return true
		}
	}

	return false
}

func createHandleResponseObjects(ctx context.Context, user models.User) (string, error) {
	var newScopes []string
	newScopes = append(newScopes, "api:USER")

	if user.Role == models.UserRolesADMIN {
		newScopes = append(newScopes, "api:ADMIN")
	}

	in120m := time.Now().Add(time.Hour * 2).Unix()
	signedToken := createJWTToken(nil, in120m, user.ID, newScopes,"")

	return signedToken, nil
}

func buildFoundResponse(ctx *app.CallbackOauthContext, redirectLocation string) error {
	ctx.ResponseData.Header().Set("Location", redirectLocation)
	ctx.ResponseData.Header().Set("Access-Control-Expose-Headers", "Location")

	return ctx.Found()
}

func buildErrorResponse(ctx *app.CallbackOauthContext) error {
	return ctx.BadRequest("Unable to process request")
}

func addItsMeSubjectNumberToUser(ctx context.Context, govID string, itsMeSubjectNumber string) error {
	user, err := userDB.GetUserByGovID(ctx, govID)
	if err != nil {
		return err
	}

	user.ItsmeSubjectNumber = itsMeSubjectNumber
	err = userDB.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
