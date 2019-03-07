package main

import (
	"fmt"
	"github.com/goadesign/goa"
	"io/ioutil"
	"github.com/mijn-app/mijn-app-backend/app"
	"net/http"
)

// ContractController implements the contract resource.
type ContractController struct {
	*goa.Controller
}

// NewContractController creates a contract controller.
func NewContractController(service *goa.Service) *ContractController {
	return &ContractController{Controller: service.NewController("ContractController")}
}

// List runs the list action.
func (c *ContractController) List(ctx *app.ListContractContext) error {
	// ContractController_List: start_implement

	uid, _, _ := getAccountIDAndScopesFromJWT(ctx)
	gov, err := userDB.GetUserGovByID(ctx, uid)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	// Put your logic here
	url := fmt.Sprintf(conf.NLX.URL+conf.NLX.SolviteersContractsApiPath+"/personen/%s/contracten", *gov.GovIdentifier)

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

	// ContractController_List: end_implement
}

// Show runs the show action.
func (c *ContractController) Show(ctx *app.ShowContractContext) error {
	// ContractController_Show: start_implement

	uid, _, _ := getAccountIDAndScopesFromJWT(ctx)
	gov, err := userDB.GetUserGovByID(ctx, uid)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	// Put your logic here
	url := fmt.Sprintf(conf.NLX.URL+conf.NLX.SolviteersContractsApiPath+"/personen/%s/contracten/%s", *gov.GovIdentifier, ctx.ContractID)

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

	// ContractController_Show: end_implement
}
