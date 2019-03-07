package main

import (
	"github.com/goadesign/goa"
	"github.com/mijn-app/mijn-app-backend/app"
	"github.com/mijn-app/mijn-app-backend/utils/sendinblue"
)

// MailController implements the mail resource.
type MailController struct {
	*goa.Controller
}

// NewMailController creates a mail controller.
func NewMailController(service *goa.Service) *MailController {
	return &MailController{Controller: service.NewController("MailController")}
}

// Send runs the send action.
func (c *MailController) Send(ctx *app.SendMailContext) error {
	// MailController_Send: start_implement

	// Put your logic here
	err := sendinblue.SendEmailTemplate(24, "info@mijn-app.io", map[string]string{
		"ORG":     ctx.Payload.Org,
		"MAIL":    ctx.Payload.Email,
		"PHONE":   ctx.Payload.Phonenumber,
		"MESSAGE": ctx.Payload.Message,
	})

	if err != nil {
		goa.LogError(ctx, "Mail is not sent for email: %s, error: %s", ctx.Payload.Email, err.Error())
		return ctx.BadRequest()
	}

	return ctx.NoContent()

	// MailController_Send: end_implement
}
