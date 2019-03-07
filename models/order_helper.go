// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "MijnApp": Model Helpers
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

// MediaType Retrieval Functions

// ListOrder returns an array of view: default.
func (m *OrderDB) ListOrder(ctx context.Context, journeyID uuid.UUID, userID uuid.UUID) []*app.Order {
	defer goa.MeasureSince([]string{"goa", "db", "order", "listorder"}, time.Now())

	var native []*Order
	var objs []*app.Order
	err := m.Db.Scopes(OrderFilterByJourney(journeyID, m.Db), OrderFilterByUser(userID, m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Order", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.OrderToOrder())
	}

	return objs
}

// OrderToOrder loads a Order and builds the default view of media type Order.
func (m *Order) OrderToOrder() *app.Order {
	order := &app.Order{}
	order.CreatedAt = &m.CreatedAt
	order.Data = &m.Data
	order.ID = &m.ID
	tmp1 := int(m.Status)
	order.Status = &tmp1
	tmp2 := &m.User
	order.User = tmp2.UserToUser() // Belongs To User

	return order
}

// OneOrder loads a Order and builds the default view of media type Order.
func (m *OrderDB) OneOrder(ctx context.Context, id uuid.UUID, journeyID uuid.UUID, userID uuid.UUID) (*app.Order, error) {
	defer goa.MeasureSince([]string{"goa", "db", "order", "oneorder"}, time.Now())

	var native Order
	err := m.Db.Scopes(OrderFilterByJourney(journeyID, m.Db), OrderFilterByUser(userID, m.Db)).Table(m.TableName()).Preload("Journey").Preload("User").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Order", "error", err.Error())
		return nil, err
	}

	view := *native.OrderToOrder()
	return &view, err
}

// MediaType Retrieval Functions

// ListOrderItem returns an array of view: item.
func (m *OrderDB) ListOrderItem(ctx context.Context, journeyID uuid.UUID, userID uuid.UUID) []*app.OrderItem {
	defer goa.MeasureSince([]string{"goa", "db", "order", "listorderitem"}, time.Now())

	var native []*Order
	var objs []*app.OrderItem
	err := m.Db.Scopes(OrderFilterByJourney(journeyID, m.Db), OrderFilterByUser(userID, m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Order", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.OrderToOrderItem())
	}

	return objs
}

// OrderToOrderItem loads a Order and builds the item view of media type Order.
func (m *Order) OrderToOrderItem() *app.OrderItem {
	order := &app.OrderItem{}
	order.ID = &m.ID
	tmp1 := int(m.Status)
	order.Status = &tmp1

	return order
}

// OneOrderItem loads a Order and builds the item view of media type Order.
func (m *OrderDB) OneOrderItem(ctx context.Context, id uuid.UUID, journeyID uuid.UUID, userID uuid.UUID) (*app.OrderItem, error) {
	defer goa.MeasureSince([]string{"goa", "db", "order", "oneorderitem"}, time.Now())

	var native Order
	err := m.Db.Scopes(OrderFilterByJourney(journeyID, m.Db), OrderFilterByUser(userID, m.Db)).Table(m.TableName()).Preload("Journey").Preload("User").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Order", "error", err.Error())
		return nil, err
	}

	view := *native.OrderToOrderItem()
	return &view, err
}