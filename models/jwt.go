// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "MijnApp": Models
//
// Command:
// $ goagen
// --design=github.com/mijn-app/mijn-app-backend/design
// --out=$(GOPATH)/src/github.com/mijn-app/mijn-app-backend
// --version=v1.4.1

package models

import (
	"context"
	"github.com/goadesign/goa"
	uuid "github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

// JWT Token
type JWT struct {
	ID           int `gorm:"primary_key"` // primary key
	CreatedAt    time.Time
	DeletedAt    *time.Time
	Pin          string
	Platform     string
	UniqueID     string
	UpdatedAt    time.Time
	UserID       uuid.UUID `sql:"type:uuid" gorm:"index:idx_jwt_user_id"` // Belongs To User
	Verification string
	Expiration   time.Time // timestamp
	User         User
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m JWT) TableName() string {
	return "j_w_t_s"

}

// JWTDB is the implementation of the storage interface for
// JWT.
type JWTDB struct {
	Db *gorm.DB
}

// NewJWTDB creates a new storage type.
func NewJWTDB(db *gorm.DB) *JWTDB {
	return &JWTDB{Db: db}
}

// DB returns the underlying database.
func (m *JWTDB) DB() interface{} {
	return m.Db
}

// JWTStorage represents the storage interface.
type JWTStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*JWT, error)
	Get(ctx context.Context, id int) (*JWT, error)
	Add(ctx context.Context, jwt *JWT) error
	Update(ctx context.Context, jwt *JWT) error
	Delete(ctx context.Context, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *JWTDB) TableName() string {
	return "j_w_t_s"

}

// Belongs To Relationships

// JWTFilterByUser is a gorm filter for a Belongs To relationship.
func JWTFilterByUser(userID uuid.UUID, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if userID != uuid.Nil {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", userID)
		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// CRUD Functions

// Get returns a single JWT as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *JWTDB) Get(ctx context.Context, id int) (*JWT, error) {
	defer goa.MeasureSince([]string{"goa", "db", "jwt", "get"}, time.Now())

	var native JWT
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of JWT
func (m *JWTDB) List(ctx context.Context) ([]*JWT, error) {
	defer goa.MeasureSince([]string{"goa", "db", "jwt", "list"}, time.Now())

	var objs []*JWT
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *JWTDB) Add(ctx context.Context, model *JWT) error {
	defer goa.MeasureSince([]string{"goa", "db", "jwt", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding JWT", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *JWTDB) Update(ctx context.Context, model *JWT) error {
	defer goa.MeasureSince([]string{"goa", "db", "jwt", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating JWT", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *JWTDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "jwt", "delete"}, time.Now())

	err := m.Db.Where("id = ?", id).Delete(&JWT{}).Error
	if err != nil {
		goa.LogError(ctx, "error deleting JWT", "error", err.Error())
		return err
	}

	return nil
}
