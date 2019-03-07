package main

import (
	"fmt"
	"github.com/goadesign/goa"
	"io/ioutil"
	"github.com/mijn-app/mijn-app-backend/app"
	"net/http"
)

// AvglogController implements the avglog resource.
type AvglogController struct {
	*goa.Controller
}

// NewAvglogController creates a avglog controller.
func NewAvglogController(service *goa.Service) *AvglogController {
	return &AvglogController{Controller: service.NewController("AvglogController")}
}

// List runs the list action.
func (c *AvglogController) List(ctx *app.ListAvglogContext) error {
	// AvglogController_List: start_implement

	uid, _, _ := getAccountIDAndScopesFromJWT(ctx)
	gov, err := userDB.GetUserGovByID(ctx, uid)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	// Put your logic here
	url := fmt.Sprintf(conf.NLX.URL+conf.NLX.SolviteersAVGLoggingApiPath+"/personen/%s/logregels", *gov.GovIdentifier)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+conf.NLX.SolviteersKey)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if resp.StatusCode == http.StatusNotFound {
		return ctx.NotFound(goa.ErrNotFound(err))
	}
	/**
	 * TODO: Add checks for 401, 403 when these are implemented at Solviteers
	 *
	 * https://solviteers.github.io/mijnapp-api-docs-swaggerui/#/contracten/get_personen__bsn__contracten
	 */
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ctx.BadRequest(goa.ErrBadRequest(err))
		}
		return ctx.OK(body)
	}
	return ctx.BadRequest(goa.ErrBadRequest(err))

	// AvglogController_List: end_implement
}

// Show runs the show action.
func (c *AvglogController) Show(ctx *app.ShowAvglogContext) error {
	// AvglogController_Show: start_implement

	uid, _, _ := getAccountIDAndScopesFromJWT(ctx)
	gov, err := userDB.GetUserGovByID(ctx, uid)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	// Put your logic here
	url := fmt.Sprintf(conf.NLX.URL+conf.NLX.SolviteersAVGLoggingApiPath+"/personen/%s/logregels/%s", *gov.GovIdentifier, ctx.AvglogID)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+conf.NLX.SolviteersKey)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if resp.StatusCode == http.StatusNotFound {
		return ctx.NotFound(goa.ErrNotFound(err))
	}
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ctx.BadRequest(goa.ErrBadRequest(err))
		}
		return ctx.OK(body)
	}
	return ctx.BadRequest(goa.ErrBadRequest(err))

	// AvglogController_Show: end_implement
}
