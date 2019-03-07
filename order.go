package main

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/mijn-app/mijn-app-backend/app"
	"github.com/mijn-app/mijn-app-backend/models"
)

// OrderController implements the order resource.
type OrderController struct {
	*goa.Controller
}

// NewOrderController creates a order controller.
func NewOrderController(service *goa.Service) *OrderController {
	return &OrderController{Controller: service.NewController("OrderController")}
}

// Create runs the create action.
func (c *OrderController) Create(ctx *app.CreateOrderContext) error {
	// OrderController_Create: start_implement

	// Put your logic here
	userID, _, _ := getAccountIDAndScopesFromJWTNoSecurity(ctx.RequestData)
	// Only admins may query accounts other than their own

	order := models.Order{
		Data:   ctx.Payload.Data,
		UserID: userID,
	}

	if !IsJSON(order.Data) {
		goa.LogError(ctx, "String is not valide JSON so we can't save the order.")
		return ctx.InternalServerError()
	}

	err := orderDB.Add(ctx, &order)
	if err != nil {
		goa.LogError(ctx, "Error occurred when creating order in DB, error: %s", err.Error())
		return ctx.InternalServerError()
	}

	return ctx.NoContent()

	// OrderController_Create: end_implement
}

// List runs the list action.
func (c *OrderController) List(ctx *app.ListOrderContext) error {
	// OrderController_List: start_implement

	// Put your logic here
	orders := orderDB.ListOrderWithUser(ctx)

	return ctx.OK(orders)

	// OrderController_List: end_implement
}

// Show runs the show action.
func (c *OrderController) Show(ctx *app.ShowOrderContext) error {
	// OrderController_Show: start_implement

	// Put your logic here
	order, err := orderDB.OneOrderByID(ctx, ctx.OrderID)
	if err != nil && err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}
	if err != nil {
		goa.LogError(ctx, "Error occurred when retrieving order from DB, error: %s", err.Error())
		return ctx.InternalServerError()
	}

	return ctx.OK(order)

	// OrderController_Show: end_implement
}
