package models

import (
	"context"
	"time"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	uuid "github.com/gofrs/uuid"
	"github.com/mijn-app/mijn-app-backend/app"
)

// ListOrderWithUser returns an array of view: default.
func (m *OrderDB) ListOrderWithUser(ctx context.Context) []*app.Order {
	defer goa.MeasureSince([]string{"goa", "db", "order", "listorderwithuser"}, time.Now())

	var native []*Order
	var objs []*app.Order
	err := m.Db.Table(m.TableName()).Preload("User").Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Order", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.OrderToOrder())
	}

	return objs
}

// OneOrderByID loads a Order on ID.
func (m *OrderDB) OneOrderByID(ctx context.Context, id uuid.UUID) (*app.Order, error) {
	defer goa.MeasureSince([]string{"goa", "db", "order", "oneorderbyid"}, time.Now())

	var native Order
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Order", "error", err.Error())
		return nil, err
	}

	view := *native.OrderToOrder()
	return &view, err
}
