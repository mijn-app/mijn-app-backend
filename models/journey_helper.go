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

// ListJourney returns an array of view: default.
func (m *JourneyDB) ListJourney(ctx context.Context, organisationID uuid.UUID) []*app.Journey {
	defer goa.MeasureSince([]string{"goa", "db", "journey", "listjourney"}, time.Now())

	var native []*Journey
	var objs []*app.Journey
	err := m.Db.Scopes(JourneyFilterByOrganisation(organisationID, m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Journey", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.JourneyToJourney())
	}

	return objs
}

// JourneyToJourney loads a Journey and builds the default view of media type Journey.
func (m *Journey) JourneyToJourney() *app.Journey {
	journey := &app.Journey{}
	journey.Data = m.Data
	journey.ID = m.ID
	tmp1 := int(m.Published)
	journey.Published = &tmp1
	tmp2 := int(m.Shared)
	journey.Shared = &tmp2

	return journey
}

// OneJourney loads a Journey and builds the default view of media type Journey.
func (m *JourneyDB) OneJourney(ctx context.Context, id uuid.UUID, organisationID uuid.UUID) (*app.Journey, error) {
	defer goa.MeasureSince([]string{"goa", "db", "journey", "onejourney"}, time.Now())

	var native Journey
	err := m.Db.Scopes(JourneyFilterByOrganisation(organisationID, m.Db)).Table(m.TableName()).Preload("Organisation").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Journey", "error", err.Error())
		return nil, err
	}

	view := *native.JourneyToJourney()
	return &view, err
}

// MediaType Retrieval Functions

// ListJourneyList returns an array of view: list.
func (m *JourneyDB) ListJourneyList(ctx context.Context, organisationID uuid.UUID) []*app.JourneyList {
	defer goa.MeasureSince([]string{"goa", "db", "journey", "listjourneylist"}, time.Now())

	var native []*Journey
	var objs []*app.JourneyList
	err := m.Db.Scopes(JourneyFilterByOrganisation(organisationID, m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Journey", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.JourneyToJourneyList())
	}

	return objs
}

// JourneyToJourneyList loads a Journey and builds the list view of media type Journey.
func (m *Journey) JourneyToJourneyList() *app.JourneyList {
	journey := &app.JourneyList{}
	journey.ID = m.ID
	tmp1 := int(m.Published)
	journey.Published = &tmp1
	tmp2 := int(m.Shared)
	journey.Shared = &tmp2

	return journey
}

// OneJourneyList loads a Journey and builds the list view of media type Journey.
func (m *JourneyDB) OneJourneyList(ctx context.Context, id uuid.UUID, organisationID uuid.UUID) (*app.JourneyList, error) {
	defer goa.MeasureSince([]string{"goa", "db", "journey", "onejourneylist"}, time.Now())

	var native Journey
	err := m.Db.Scopes(JourneyFilterByOrganisation(organisationID, m.Db)).Table(m.TableName()).Preload("Organisation").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Journey", "error", err.Error())
		return nil, err
	}

	view := *native.JourneyToJourneyList()
	return &view, err
}
