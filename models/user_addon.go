package models

import (
	"context"
	"time"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	uuid "github.com/gofrs/uuid"
	"github.com/mijn-app/mijn-app-backend/app"
)

// ListUsers returns an array of users.
func (m *UserDB) ListUsers(ctx context.Context, page int, limit int) ([]*app.User, int) {
	var objs []*app.User
	var count int
	err := m.Db.Model(&User{}).
		Table(m.TableName()).
		Where("users.deleted_at IS NULL").
		Count(&count).
		Limit(limit).
		Offset(page * limit).
		Order("users.id asc").
		Find(&objs).Error
	if err != nil {
		goa.LogError(ctx, "error listing User", "error", err.Error())
		return objs, count
	}

	return objs, count
}

// OneUserOnID loads a User from DB based on id
func (m *UserDB) OneUserOnID(id uuid.UUID) (*User, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "OneUserOnID"}, time.Now())

	var native User
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	return &native, err
}

// ListUsers returns an array of users.
func (m *UserDB) ListAllUsers(ctx context.Context, page int, limit int) ([]*app.User, int) {
	var objs []*app.User
	var count int
	err := m.Db.Model(&User{}).
		Table(m.TableName()).
		Count(&count).
		Limit(limit).
		Offset(page * limit).
		Order("users.id asc").
		Find(&objs).Error
	if err != nil {
		goa.LogError(ctx, "error listing User", "error", err.Error())
		return objs, count
	}

	return objs, count
}

// GetUserByID gets user by ID
func (m *UserDB) GetUserByID(ctx context.Context, userID uuid.UUID) (*app.User, error) {
	var user User
	var res *app.User
	err := m.Db.Model(&User{}).
		Preload("Addres").
		Table(m.TableName()).
		Where("users.id = ? AND users.deleted_at IS NULL", userID).
		Find(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, err
	}

	res = user.UserToUser()
	res.Address = user.Addres.AddressToAddress()

	return res, err
}

// GetUserByID gets user by Email
func (m *UserDB) GetUserByEmail(ctx context.Context, userEmail string) (*app.UserAuth, *app.User, error) {
	var user User
	err := m.Db.Model(&User{}).
		Table(m.TableName()).
		Where("users.email = ? AND users.deleted_at IS NULL", userEmail).
		Find(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, nil, err
	}

	return user.UserToUserAuth(), user.UserToUser(), err
}

// GetUserGovByID gets a user's government id by user ID
func (m *UserDB) GetUserGovByID(ctx context.Context, userID uuid.UUID) (*app.UserGov, error) {
	var user User
	err := m.Db.Model(&User{}).
		Table(m.TableName()).
		Where("users.id = ?", userID).
		Find(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, err
	}

	return user.UserToUserGov(), err
}

func (m *UserDB) GetUserByItsMeSubjectNumber(ctx context.Context, subjectNumber string) (*User, error) {
	var user User
	err := m.Db.Model(&User{}).
		Preload("Addres").
		Table(m.TableName()).
		Where("users.itsme_subject_number = ?", subjectNumber).
		Find(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, err
	}

	return &user, err
}

func (m *UserDB) GetUserByGovID(ctx context.Context, govID string) (*User, error) {
	var user User
	err := m.Db.Model(&User{}).
		Preload("Addres").
		Table(m.TableName()).
		Where("users.gov_identifier = ?", govID).
		Find(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, err
	}

	return &user, err
}
