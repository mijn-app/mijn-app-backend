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
	"github.com/mijn-app/mijn-app-backend/app"
	"time"
)

// Model to store an order from an authorized client.
type Order struct {
	ID        uuid.UUID `sql:"type:uuid;default:uuid_generate_v4()" gorm:"primary_key"` // primary key
	CreatedAt time.Time
	Data      string
	DeletedAt *time.Time
	JourneyID uuid.UUID   `sql:"type:uuid" gorm:"index:idx_order_journey_id"` // Belongs To Journey
	Status    OrderStatus `sql:"type:smallint"`                               // enum OrderStatus
	UpdatedAt time.Time
	UserID    uuid.UUID `sql:"type:uuid" gorm:"index:idx_order_user_id"` // Belongs To User
	Journey   Journey
	User      User
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Order) TableName() string {
	return "orders"

}

// OrderDB is the implementation of the storage interface for
// Order.
type OrderDB struct {
	Db *gorm.DB
}

// NewOrderDB creates a new storage type.
func NewOrderDB(db *gorm.DB) *OrderDB {
	return &OrderDB{Db: db}
}

// DB returns the underlying database.
func (m *OrderDB) DB() interface{} {
	return m.Db
}

// OrderStorage represents the storage interface.
type OrderStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Order, error)
	Get(ctx context.Context, id uuid.UUID) (*Order, error)
	Add(ctx context.Context, order *Order) error
	Update(ctx context.Context, order *Order) error
	Delete(ctx context.Context, id uuid.UUID) error

	ListOrder(ctx context.Context, journeyID int, userID int) []*app.Order
	OneOrder(ctx context.Context, id uuid.UUID, journeyID int, userID int) (*app.Order, error)

	ListOrderItem(ctx context.Context, journeyID int, userID int) []*app.OrderItem
	OneOrderItem(ctx context.Context, id uuid.UUID, journeyID int, userID int) (*app.OrderItem, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *OrderDB) TableName() string {
	return "orders"

}

// Belongs To Relationships

// OrderFilterByJourney is a gorm filter for a Belongs To relationship.
func OrderFilterByJourney(journeyID uuid.UUID, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if journeyID != uuid.Nil {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("journey_id = ?", journeyID)
		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// Belongs To Relationships

// OrderFilterByUser is a gorm filter for a Belongs To relationship.
func OrderFilterByUser(userID uuid.UUID, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if userID != uuid.Nil {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", userID)
		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// CRUD Functions

// Get returns a single Order as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *OrderDB) Get(ctx context.Context, id uuid.UUID) (*Order, error) {
	defer goa.MeasureSince([]string{"goa", "db", "order", "get"}, time.Now())

	var native Order
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Order
func (m *OrderDB) List(ctx context.Context) ([]*Order, error) {
	defer goa.MeasureSince([]string{"goa", "db", "order", "list"}, time.Now())

	var objs []*Order
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *OrderDB) Add(ctx context.Context, model *Order) error {
	defer goa.MeasureSince([]string{"goa", "db", "order", "add"}, time.Now())

	model.ID = uuid.Must(uuid.NewV4())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Order", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *OrderDB) Update(ctx context.Context, model *Order) error {
	defer goa.MeasureSince([]string{"goa", "db", "order", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Order", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *OrderDB) Delete(ctx context.Context, id uuid.UUID) error {
	defer goa.MeasureSince([]string{"goa", "db", "order", "delete"}, time.Now())

	err := m.Db.Where("id = ?", id).Delete(&Order{}).Error
	if err != nil {
		goa.LogError(ctx, "error deleting Order", "error", err.Error())
		return err
	}

	return nil
}
