package main

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/mijn-app/mijn-app-backend/app"
	"github.com/mijn-app/mijn-app-backend/models"
)

// AddressController implements the address resource.
type AddressController struct {
	*goa.Controller
}

// NewAddressController creates a address controller.
func NewAddressController(service *goa.Service) *AddressController {
	return &AddressController{Controller: service.NewController("AddressController")}
}

// Create runs the create action.
func (c *AddressController) Create(ctx *app.CreateAddressContext) error {
	// AddressController_Create: start_implement

	UserID, scopes, _ := getAccountIDAndScopesFromJWT(ctx)

	// Only users that are owner of the address details and dashboard user can access this endpoint
	if ctx.UserID != UserID || isDashboard(scopes) {
		return ctx.Unauthorized()
	}

	foundAddress, err := addressDB.OneAddressByUser(ctx, ctx.UserID)
	if foundAddress != nil && err == nil {
		return ctx.Conflict(ErrConflict("Address for userID already exists"))
	}

	// Put your logic here
	address := models.AddressFromCreateAddressPayload(ctx.Payload)
	address.UserID = UserID

	tx := addressDB.Db.Begin()
	err = tx.Create(&address).Error
	if err != nil {
		tx.Rollback()
		goa.LogError(ctx, "Error occurred when creating address in DB, error: %s", err.Error())
		return ctx.InternalServerError()
	}

	tx.Commit()

	return ctx.NoContent()

	// AddressController_Create: end_implement
}

// Show runs the show action.
func (c *AddressController) Show(ctx *app.ShowAddressContext) error {
	// AddressController_Show: start_implement

	// Put your logic here
	UserID, scopes, _ := getAccountIDAndScopesFromJWT(ctx)

	address, err := addressDB.OneAddressByUser(ctx, ctx.UserID)
	if err != nil && err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}
	if err != nil {
		goa.LogError(ctx, "Error occurred when retrieving address from DB, error: %s", err.Error())
		return ctx.InternalServerError()
	}

	// Only users that are owner of the address details and dashboard user can access this endpoint
	if ctx.UserID != UserID && !isDashboard(scopes) {
		return ctx.Unauthorized()
	}

	return ctx.OK(address)

	// AddressController_Show: end_implement
}

// Update runs the update action.
func (c *AddressController) Update(ctx *app.UpdateAddressContext) error {
	// AddressController_Update: start_implement

	// Put your logic here
	UserID, scopes, _ := getAccountIDAndScopesFromJWT(ctx)

	// Only users that are owner of the address details and dashboard user can access this endpoint
	if ctx.UserID != UserID && !isDashboard(scopes) {
		return ctx.Unauthorized()
	}

	address, err := addressDB.OneAddressByUser(ctx, ctx.UserID)
	if err != nil && err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}
	if err != nil {
		goa.LogError(ctx, "Error occurred when retrieving address from DB, error: %s", err.Error())
		return ctx.InternalServerError()
	}

	err = addressDB.UpdateFromUpdateAddressPayload(ctx, ctx.Payload, *address.ID)
	if err != nil {
		goa.LogError(ctx, "Error updating address in DB, error: %s", err.Error())
		return ctx.InternalServerError()
	}

	return ctx.NoContent()

	// AddressController_Update: end_implement
}
