// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "MijnApp": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/mijn-app/mijn-app-backend/design
// --out=$(GOPATH)/src/github.com/mijn-app/mijn-app-backend
// --version=v1.4.1

package app

import (
	"github.com/goadesign/goa"
	uuid "github.com/gofrs/uuid"
	"time"
)

// Mediatype for addresses (default view)
//
// Identifier: application/vnd.address+json; view=default
type Address struct {
	// Address country
	Country *string `form:"country,omitempty" json:"country,omitempty" yaml:"country,omitempty" xml:"country,omitempty"`
	// Address house number
	HouseNumber *string `form:"house_number,omitempty" json:"house_number,omitempty" yaml:"house_number,omitempty" xml:"house_number,omitempty"`
	// Address unique identifier
	ID *uuid.UUID `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// Address location
	Location *string `form:"location,omitempty" json:"location,omitempty" yaml:"location,omitempty" xml:"location,omitempty"`
	// Address streetnaam
	Street *string    `form:"street,omitempty" json:"street,omitempty" yaml:"street,omitempty" xml:"street,omitempty"`
	UserID *uuid.UUID `form:"user_id,omitempty" json:"user_id,omitempty" yaml:"user_id,omitempty" xml:"user_id,omitempty"`
	// Address zipcode
	Zipcode *string `form:"zipcode,omitempty" json:"zipcode,omitempty" yaml:"zipcode,omitempty" xml:"zipcode,omitempty"`
}

// Current app state (default view)
//
// Identifier: application/vnd.appstate+json; view=default
type Appstate struct {
	User *User `form:"user" json:"user" yaml:"user" xml:"user"`
}

// Validate validates the Appstate media type instance.
func (mt *Appstate) Validate() (err error) {
	if mt.User == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "user"))
	}
	return
}

// The health check for the API (default view)
//
// Identifier: application/vnd.health+json; view=default
type Health struct {
	// True if API is healthy
	Health bool `form:"health" json:"health" yaml:"health" xml:"health"`
}

// Journey mediatype (default view)
//
// Identifier: application/vnd.journey+json; view=default
type Journey struct {
	// The journey data as a json blob
	Data string `form:"data" json:"data" yaml:"data" xml:"data"`
	// ID of the journey in the DB
	ID uuid.UUID `form:"id" json:"id" yaml:"id" xml:"id"`
	// Whether this journey has been published
	Published *int `form:"published,omitempty" json:"published,omitempty" yaml:"published,omitempty" xml:"published,omitempty"`
	// Whether this journey has been shared
	Shared *int `form:"shared,omitempty" json:"shared,omitempty" yaml:"shared,omitempty" xml:"shared,omitempty"`
}

// Validate validates the Journey media type instance.
func (mt *Journey) Validate() (err error) {

	if mt.Data == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	return
}

// Journey mediatype (list view)
//
// Identifier: application/vnd.journey+json; view=list
type JourneyList struct {
	// ID of the journey in the DB
	ID uuid.UUID `form:"id" json:"id" yaml:"id" xml:"id"`
	// Whether this journey has been published
	Published *int `form:"published,omitempty" json:"published,omitempty" yaml:"published,omitempty" xml:"published,omitempty"`
	// Whether this journey has been shared
	Shared *int `form:"shared,omitempty" json:"shared,omitempty" yaml:"shared,omitempty" xml:"shared,omitempty"`
}

// Order mediatype (default view)
//
// Identifier: application/vnd.order+json; view=default
type Order struct {
	// Date of creation
	CreatedAt *time.Time `form:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty"`
	// The order data as a json blob
	Data *string `form:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty" xml:"data,omitempty"`
	// ID of the order in the DB
	ID *uuid.UUID `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// The status of the order
	Status *int  `form:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty" xml:"status,omitempty"`
	User   *User `form:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty" xml:"user,omitempty"`
}

// Order mediatype (item view)
//
// Identifier: application/vnd.order+json; view=item
type OrderItem struct {
	// ID of the order in the DB
	ID *uuid.UUID `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// The status of the order
	Status *int `form:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty" xml:"status,omitempty"`
}

// OrderCollection is the media type for an array of Order (default view)
//
// Identifier: application/vnd.order+json; type=collection; view=default
type OrderCollection []*Order

// OrderCollection is the media type for an array of Order (item view)
//
// Identifier: application/vnd.order+json; type=collection; view=item
type OrderItemCollection []*OrderItem

// OrderAttachment mediatype (default view)
//
// Identifier: application/vnd.orderattachment+json; view=default
type Orderattachment struct {
	// ID of the attachment for the order in the DB
	ID *uuid.UUID `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// The path to the attachment
	Path *string `form:"path,omitempty" json:"path,omitempty" yaml:"path,omitempty" xml:"path,omitempty"`
}

// OrderLog mediatype (default view)
//
// Identifier: application/vnd.orderlog+json; view=default
type Orderlog struct {
	// ID of the log in the DB
	ID *uuid.UUID `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// The message-body for this log
	Message *string `form:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty" xml:"message,omitempty"`
}

// Organisation mediatype (default view)
//
// Identifier: application/vnd.organisation+json; view=default
type Organisation struct {
	// ID of the ogranisation in the DB
	ID *uuid.UUID `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// Display name of the organisation
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// OrganisationAPI mediatype (default view)
//
// Identifier: application/vnd.organisationapi+json; view=default
type Organisationapi struct {
	// ID of the api endpoint in the DB
	ID *uuid.UUID `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// The name of an organisation's API endpoint
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	// The api type
	Type *int `form:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty" xml:"type,omitempty"`
	// The api endpoint url
	URL *string `form:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty" xml:"url,omitempty"`
}

// OrganisationAPIParam mediatype (default view)
//
// Identifier: application/vnd.organisationapiparam+json; view=default
type Organisationapiparam struct {
	// ID of the api endpoint parameter
	ID *uuid.UUID `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// The api endpoint parameter key
	Key *string `form:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty" xml:"key,omitempty"`
	// The display name for the api endpoint parameter
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// User mediatype (auth view)
//
// Identifier: application/vnd.user+json; view=auth
type UserAuth struct {
	// Email of the user
	Email *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	// ID of the user in the DB
	ID *uuid.UUID `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// The role of the user
	Role *int `form:"role,omitempty" json:"role,omitempty" yaml:"role,omitempty" xml:"role,omitempty"`
}

// User mediatype (compact view)
//
// Identifier: application/vnd.user+json; view=compact
type UserCompact struct {
	// Email of the user
	Email *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	// First name of the user
	FirstName *string `form:"first_name,omitempty" json:"first_name,omitempty" yaml:"first_name,omitempty" xml:"first_name,omitempty"`
	// ID of the user in the DB
	ID *uuid.UUID `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// Family name of the user
	LastName *string `form:"last_name,omitempty" json:"last_name,omitempty" yaml:"last_name,omitempty" xml:"last_name,omitempty"`
}

// User mediatype (default view)
//
// Identifier: application/vnd.user+json; view=default
type User struct {
	Address     *Address   `form:"address,omitempty" json:"address,omitempty" yaml:"address,omitempty" xml:"address,omitempty"`
	DateOfBirth *time.Time `form:"date_of_birth,omitempty" json:"date_of_birth,omitempty" yaml:"date_of_birth,omitempty" xml:"date_of_birth,omitempty"`
	// Email of the user
	Email *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	// First name of the user
	FirstName *string `form:"first_name,omitempty" json:"first_name,omitempty" yaml:"first_name,omitempty" xml:"first_name,omitempty"`
	Gender    *int    `form:"gender,omitempty" json:"gender,omitempty" yaml:"gender,omitempty" xml:"gender,omitempty"`
	// Government identification token
	GovIdentifier *string `form:"gov_identifier,omitempty" json:"gov_identifier,omitempty" yaml:"gov_identifier,omitempty" xml:"gov_identifier,omitempty"`
	// ID of the user in the DB
	ID *uuid.UUID `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// Family name of the user
	LastName *string `form:"last_name,omitempty" json:"last_name,omitempty" yaml:"last_name,omitempty" xml:"last_name,omitempty"`
	// Prefix for the users family name
	LastNamePrefix *string `form:"last_name_prefix,omitempty" json:"last_name_prefix,omitempty" yaml:"last_name_prefix,omitempty" xml:"last_name_prefix,omitempty"`
	PhoneNumber    *string `form:"phone_number,omitempty" json:"phone_number,omitempty" yaml:"phone_number,omitempty" xml:"phone_number,omitempty"`
	// The role of the user
	Role *int `form:"role,omitempty" json:"role,omitempty" yaml:"role,omitempty" xml:"role,omitempty"`
}

// User mediatype (gov view)
//
// Identifier: application/vnd.user+json; view=gov
type UserGov struct {
	// Government identification token
	GovIdentifier *string `form:"gov_identifier,omitempty" json:"gov_identifier,omitempty" yaml:"gov_identifier,omitempty" xml:"gov_identifier,omitempty"`
}

// UserCollection is the media type for an array of User (auth view)
//
// Identifier: application/vnd.user+json; type=collection; view=auth
type UserAuthCollection []*UserAuth

// UserCollection is the media type for an array of User (compact view)
//
// Identifier: application/vnd.user+json; type=collection; view=compact
type UserCompactCollection []*UserCompact

// UserCollection is the media type for an array of User (default view)
//
// Identifier: application/vnd.user+json; type=collection; view=default
type UserCollection []*User

// UserCollection is the media type for an array of User (gov view)
//
// Identifier: application/vnd.user+json; type=collection; view=gov
type UserGovCollection []*UserGov

// UserOrganisationData mediatype (default view)
//
// Identifier: application/vnd.userorganisationdatamedia+json; view=default
type Userorganisationdatamedia struct {
	// The user's identifier in the organisation's remote endpoints
	ForeignID *string `form:"foreign_id,omitempty" json:"foreign_id,omitempty" yaml:"foreign_id,omitempty" xml:"foreign_id,omitempty"`
	// The user's token for the organisation's remote endpoints
	ForeignToken *string `form:"foreign_token,omitempty" json:"foreign_token,omitempty" yaml:"foreign_token,omitempty" xml:"foreign_token,omitempty"`
	// ID of the user's organisation-data in the DB
	ID *uuid.UUID `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
}

// UserOrganisationMember mediatype (default view)
//
// Identifier: application/vnd.userorganisationmember+json; view=default
type Userorganisationmember struct {
	// ID of the user's organisation-membership in the DB
	ID *uuid.UUID `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// The user's role inside the organisation
	Role *int `form:"role,omitempty" json:"role,omitempty" yaml:"role,omitempty" xml:"role,omitempty"`
}
