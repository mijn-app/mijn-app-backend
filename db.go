package main

import (
	"github.com/jinzhu/gorm"
	"github.com/mijn-app/mijn-app-backend/models"
)

var db *gorm.DB

// Set database vars
var userDB *models.UserDB
var journeyDB *models.JourneyDB
var jwtDB *models.JWTDB
var orderDB *models.OrderDB
var orderLogDB *models.OrderLogDB
var orderAttachmentDB *models.OrderAttachmentDB
var organisationDB *models.OrganisationDB
var userOrganisationDataDB *models.UserOrganisationDataDB
var userOrganisationMemberDB *models.UserOrganisationMemberDB
var organisationAPIDB *models.OrganisationAPIDB
var organisationAPIParamDB *models.OrganisationAPIParamDB
var addressDB *models.AddressDB
var oAuthStateDB *models.OAuthStateDB

func initDatabase(logMode bool, testMode bool) {
	db = models.InitDatabase(logMode, testMode)

	userDB = models.NewUserDB(db)
	journeyDB = models.NewJourneyDB(db)
	jwtDB = models.NewJWTDB(db)
	orderDB = models.NewOrderDB(db)
	orderLogDB = models.NewOrderLogDB(db)
	orderAttachmentDB = models.NewOrderAttachmentDB(db)
	organisationDB = models.NewOrganisationDB(db)
	userOrganisationDataDB = models.NewUserOrganisationDataDB(db)
	userOrganisationMemberDB = models.NewUserOrganisationMemberDB(db)
	organisationAPIDB = models.NewOrganisationAPIDB(db)
	organisationAPIParamDB = models.NewOrganisationAPIParamDB(db)
	addressDB = models.NewAddressDB(db)
	oAuthStateDB = models.NewOAuthStateDB(db)
}
