package main

import (
	"fmt"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"github.com/mijn-app/mijn-app-backend/app"
	"github.com/mijn-app/mijn-app-backend/models"
	"strconv"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	// UserController_Create: start_implement

	foundUser, _, err := userDB.GetUserByEmail(ctx, ctx.Payload.Email)
	if foundUser != nil && err == nil {
		return ctx.Conflict(ErrConflict("Duplicate key: email"))
	}

	user := models.UserFromCreateUserPayload(ctx.Payload)
	user.Role = models.UserRolesUSER
	pass, err := bcrypt.GenerateFromPassword([]byte(ctx.Payload.Password), bcrypt.DefaultCost)
	if err != nil {
		goa.LogError(ctx, "Error generating new password, error: %s", err.Error())
		return ctx.InternalServerError()
	}
	user.Password = string(pass)

	tx := userDB.Db.Begin()

	err = tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		goa.LogError(ctx, "Error occurred when creating user in DB, error: %s", err.Error())
		return ctx.InternalServerError()
	}

	tx.Commit()
	return ctx.OK(user.UserToUser())

	// UserController_Create: end_implement
}

// List runs the list action.
func (c *UserController) List(ctx *app.ListUserContext) error {
	// UserController_List: start_implement

	_, scopes, _ := getAccountIDAndScopesFromJWT(ctx)
	// Only admins may query accounts other than their own
	if !isAdmin(scopes) {
		return ctx.Unauthorized()
	}

	var users []*app.User
	var count int

	users, count = userDB.ListUsers(ctx, 0, 99999)

	ctx.ResponseData.Header().Set("X-List-Count", strconv.Itoa(count))
	ctx.ResponseData.Header().Set("Access-Control-Expose-Headers", "X-List-Count")
	return ctx.OK(users)

	// UserController_List: end_implement
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	// UserController_Show: start_implement

	userID, scopes, _ := getAccountIDAndScopesFromJWT(ctx)
	fmt.Println(isAdmin(scopes))
	// Only admins may query accounts other than their own
	if userID != ctx.UserID && !isAdmin(scopes) {
		return ctx.Unauthorized()
	}

	user, err := userDB.GetUserByID(ctx, ctx.UserID)
	if err != nil && err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}
	if err != nil {
		goa.LogError(ctx, "Error occurred when retrieving user from DB, error: %s", err.Error())
		return ctx.InternalServerError()
	}

	return ctx.OK(user)

	// UserController_Show: end_implement
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	// UserController_Update: start_implement

	userID, scopes, _ := getAccountIDAndScopesFromJWT(ctx)
	// Only users that are owner of the address details and dashboard user can access this endpoint
	if userID != ctx.UserID && !isAdmin(scopes) {
		return ctx.Unauthorized()
	}

	_, err := userDB.OneUser(ctx, ctx.UserID)
	if err != nil && err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}
	if err != nil {
		goa.LogError(ctx, "Error occurred when retrieving user from DB, error: %s", err.Error())
		return ctx.InternalServerError()
	}

	user := models.UserFromUpdateUserPayload(ctx.Payload)
	user.ID = ctx.UserID

	err = userDB.Update(ctx, user)
	if err != nil {
		goa.LogError(ctx, "Unable to update user in DB, error: %s", err.Error())
		return ctx.InternalServerError()
	}

	updatedUser, err := userDB.OneUserCompact(ctx, ctx.UserID)
	if err != nil {
		goa.LogError(ctx, "Error occurred when retrieving user from DB, error: %s", err.Error())
		return ctx.InternalServerError()
	}

	return ctx.OKCompact(updatedUser)

	// UserController_Update: end_implement
}
