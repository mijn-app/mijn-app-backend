package models

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/jinzhu/gorm"
	"time"

	"github.com/goadesign/goa"
	"github.com/gofrs/uuid"
)

// HardDeleteOAuthStateOnID hard deletes an OAuth State based on token
func (m *OAuthStateDB) HardDeleteOAuthStateOnID(id uuid.UUID) error {
	defer goa.MeasureSince([]string{"goa", "db", "oAuthState", "HardDeleteOAuthStateOnID"}, time.Now())

	err := m.Db.Unscoped().Table(m.TableName()).Where("id = ?", id).Delete(&OAuthState{}).Error
	return err
}

// OneOAuthStateOnToken loads an OauthState from DB based on token
func (m *OAuthStateDB) OneOAuthStateOnToken(token string) (*OAuthState, error) {
	defer goa.MeasureSince([]string{"goa", "db", "oAuthState", "OneOAuthStateOnToken"}, time.Now())

	var native OAuthState
	err := m.Db.Scopes().Table(m.TableName()).Where("token = ?", token).Find(&native).Error
	return &native, err
}

// Generate and insert a random state token
func (m *OAuthStateDB) GenerateStateToken(ctx context.Context) (string, error) {
	token := generateRandomNumber()
	expiration := time.Now().Add(time.Hour)

	oAuthState := OAuthState{
		Token:      token,
		Expiration: expiration,
	}

	err := m.Add(ctx, &oAuthState)

	return token, err
}

// Validate a random state token
func (m *OAuthStateDB) ValidateStateToken(token string) (bool, error) {
	oAuthState, err := m.OneOAuthStateOnToken(token)
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	// As this token is processed we delete it regardless of validity outcome
	err = m.HardDeleteOAuthStateOnID(oAuthState.ID)
	if err != nil {
		return false, err
	}

	if time.Now().After(oAuthState.Expiration) {
		return false, nil
	}

	return true, nil
}

// Generate a random cryptographic number
func generateRandomNumber() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
