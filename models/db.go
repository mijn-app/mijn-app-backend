package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mijn-app/mijn-app-backend/utils/database"
)

// Init db
func InitDatabase(logMode bool, testMode bool) *gorm.DB {
	m := []interface{}{
		&User{},
		&Journey{},
		&Order{},
		&OrderLog{},
		&OrderAttachment{},
		&Organisation{},
		&UserOrganisationData{},
		&UserOrganisationMember{},
		&OrganisationAPI{},
		&OrganisationAPIParam{},
		&JWT{},
		&Address{},
		&OAuthState{},
	}

	db := database.Init("mijn-app", m, logMode, testMode)
	db.DB().SetMaxOpenConns(50)

	return db
}
