// API "MijnApp": Model Addon
//
// Command:
// $ goagen
// --design=github.com/mijn-app/mijn-app-backend/design
// --out=$(GOPATH)/src/github.com/mijn-app/mijn-app-backend
// --version=v1.3.1

package models

import (
	"context"
	"time"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	uuid "github.com/gofrs/uuid"
	"github.com/mijn-app/mijn-app-backend/app"
)

// OneAddressByUser loads a Address by UserID.
func (m *AddressDB) OneAddressByUser(ctx context.Context, userID uuid.UUID) (*app.Address, error) {
	defer goa.MeasureSince([]string{"goa", "db", "address", "oneaddress"}, time.Now())

	var native Address
	err := m.Db.Scopes().Table(m.TableName()).Where("user_id = ?", userID).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Address on UserID", "error", err.Error())
		return nil, err
	}

	view := *native.AddressToAddress()
	return &view, err
}
