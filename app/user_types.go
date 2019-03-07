// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "MijnApp": Application User Types
//
// Command:
// $ goagen
// --design=github.com/mijn-app/mijn-app-backend/design
// --out=$(GOPATH)/src/github.com/mijn-app/mijn-app-backend
// --version=v1.4.1

package app

import (
	"github.com/goadesign/goa"
	"time"
	"unicode/utf8"
)

// createAddressPayload user type.
type createAddressPayload struct {
	Country     *string `form:"country,omitempty" json:"country,omitempty" yaml:"country,omitempty" xml:"country,omitempty"`
	HouseNumber *string `form:"house_number,omitempty" json:"house_number,omitempty" yaml:"house_number,omitempty" xml:"house_number,omitempty"`
	Location    *string `form:"location,omitempty" json:"location,omitempty" yaml:"location,omitempty" xml:"location,omitempty"`
	Street      *string `form:"street,omitempty" json:"street,omitempty" yaml:"street,omitempty" xml:"street,omitempty"`
	Zipcode     *string `form:"zipcode,omitempty" json:"zipcode,omitempty" yaml:"zipcode,omitempty" xml:"zipcode,omitempty"`
}

// Validate validates the createAddressPayload type instance.
func (ut *createAddressPayload) Validate() (err error) {
	if ut.Location == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "location"))
	}
	if ut.Country == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "country"))
	}
	if ut.Street == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "street"))
	}
	if ut.Zipcode == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "zipcode"))
	}
	if ut.HouseNumber == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "house_number"))
	}
	return
}

// Publicize creates CreateAddressPayload from createAddressPayload
func (ut *createAddressPayload) Publicize() *CreateAddressPayload {
	var pub CreateAddressPayload
	if ut.Country != nil {
		pub.Country = *ut.Country
	}
	if ut.HouseNumber != nil {
		pub.HouseNumber = *ut.HouseNumber
	}
	if ut.Location != nil {
		pub.Location = *ut.Location
	}
	if ut.Street != nil {
		pub.Street = *ut.Street
	}
	if ut.Zipcode != nil {
		pub.Zipcode = *ut.Zipcode
	}
	return &pub
}

// CreateAddressPayload user type.
type CreateAddressPayload struct {
	Country     string `form:"country" json:"country" yaml:"country" xml:"country"`
	HouseNumber string `form:"house_number" json:"house_number" yaml:"house_number" xml:"house_number"`
	Location    string `form:"location" json:"location" yaml:"location" xml:"location"`
	Street      string `form:"street" json:"street" yaml:"street" xml:"street"`
	Zipcode     string `form:"zipcode" json:"zipcode" yaml:"zipcode" xml:"zipcode"`
}

// Validate validates the CreateAddressPayload type instance.
func (ut *CreateAddressPayload) Validate() (err error) {
	if ut.Location == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "location"))
	}
	if ut.Country == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "country"))
	}
	if ut.Street == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "street"))
	}
	if ut.Zipcode == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "zipcode"))
	}
	if ut.HouseNumber == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "house_number"))
	}
	return
}

// createUserPayload user type.
type createUserPayload struct {
	DateOfBirth    *time.Time `form:"date_of_birth,omitempty" json:"date_of_birth,omitempty" yaml:"date_of_birth,omitempty" xml:"date_of_birth,omitempty"`
	Email          *string    `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	FirstName      *string    `form:"first_name,omitempty" json:"first_name,omitempty" yaml:"first_name,omitempty" xml:"first_name,omitempty"`
	Gender         *int       `form:"gender,omitempty" json:"gender,omitempty" yaml:"gender,omitempty" xml:"gender,omitempty"`
	GovIdentifier  *string    `form:"gov_identifier,omitempty" json:"gov_identifier,omitempty" yaml:"gov_identifier,omitempty" xml:"gov_identifier,omitempty"`
	LastName       *string    `form:"last_name,omitempty" json:"last_name,omitempty" yaml:"last_name,omitempty" xml:"last_name,omitempty"`
	LastNamePrefix *string    `form:"last_name_prefix,omitempty" json:"last_name_prefix,omitempty" yaml:"last_name_prefix,omitempty" xml:"last_name_prefix,omitempty"`
	Password       *string    `form:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
	PhoneNumber    *string    `form:"phone_number,omitempty" json:"phone_number,omitempty" yaml:"phone_number,omitempty" xml:"phone_number,omitempty"`
}

// Validate validates the createUserPayload type instance.
func (ut *createUserPayload) Validate() (err error) {
	if ut.FirstName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "first_name"))
	}
	if ut.LastName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "last_name"))
	}
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "email"))
	}
	if ut.GovIdentifier == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "gov_identifier"))
	}
	if ut.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "password"))
	}
	if ut.Email != nil {
		if ok := goa.ValidatePattern(`^[^@\s:]+@[^@\s]+\.[^@\s]+$`, *ut.Email); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.email`, *ut.Email, `^[^@\s:]+@[^@\s]+\.[^@\s]+$`))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 8, true))
		}
	}
	return
}

// Publicize creates CreateUserPayload from createUserPayload
func (ut *createUserPayload) Publicize() *CreateUserPayload {
	var pub CreateUserPayload
	if ut.DateOfBirth != nil {
		pub.DateOfBirth = ut.DateOfBirth
	}
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.FirstName != nil {
		pub.FirstName = *ut.FirstName
	}
	if ut.Gender != nil {
		pub.Gender = ut.Gender
	}
	if ut.GovIdentifier != nil {
		pub.GovIdentifier = *ut.GovIdentifier
	}
	if ut.LastName != nil {
		pub.LastName = *ut.LastName
	}
	if ut.LastNamePrefix != nil {
		pub.LastNamePrefix = ut.LastNamePrefix
	}
	if ut.Password != nil {
		pub.Password = *ut.Password
	}
	if ut.PhoneNumber != nil {
		pub.PhoneNumber = ut.PhoneNumber
	}
	return &pub
}

// CreateUserPayload user type.
type CreateUserPayload struct {
	DateOfBirth    *time.Time `form:"date_of_birth,omitempty" json:"date_of_birth,omitempty" yaml:"date_of_birth,omitempty" xml:"date_of_birth,omitempty"`
	Email          string     `form:"email" json:"email" yaml:"email" xml:"email"`
	FirstName      string     `form:"first_name" json:"first_name" yaml:"first_name" xml:"first_name"`
	Gender         *int       `form:"gender,omitempty" json:"gender,omitempty" yaml:"gender,omitempty" xml:"gender,omitempty"`
	GovIdentifier  string     `form:"gov_identifier" json:"gov_identifier" yaml:"gov_identifier" xml:"gov_identifier"`
	LastName       string     `form:"last_name" json:"last_name" yaml:"last_name" xml:"last_name"`
	LastNamePrefix *string    `form:"last_name_prefix,omitempty" json:"last_name_prefix,omitempty" yaml:"last_name_prefix,omitempty" xml:"last_name_prefix,omitempty"`
	Password       string     `form:"password" json:"password" yaml:"password" xml:"password"`
	PhoneNumber    *string    `form:"phone_number,omitempty" json:"phone_number,omitempty" yaml:"phone_number,omitempty" xml:"phone_number,omitempty"`
}

// Validate validates the CreateUserPayload type instance.
func (ut *CreateUserPayload) Validate() (err error) {
	if ut.FirstName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "first_name"))
	}
	if ut.LastName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "last_name"))
	}
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "email"))
	}
	if ut.GovIdentifier == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "gov_identifier"))
	}
	if ut.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "password"))
	}
	if ok := goa.ValidatePattern(`^[^@\s:]+@[^@\s]+\.[^@\s]+$`, ut.Email); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.email`, ut.Email, `^[^@\s:]+@[^@\s]+\.[^@\s]+$`))
	}
	if utf8.RuneCountInString(ut.Password) < 8 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.password`, ut.Password, utf8.RuneCountInString(ut.Password), 8, true))
	}
	return
}

// pinPayload user type.
type pinPayload struct {
	Pin *string `form:"pin,omitempty" json:"pin,omitempty" yaml:"pin,omitempty" xml:"pin,omitempty"`
}

// Validate validates the pinPayload type instance.
func (ut *pinPayload) Validate() (err error) {
	if ut.Pin == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "pin"))
	}
	if ut.Pin != nil {
		if ok := goa.ValidatePattern(`^\d{4,6}$`, *ut.Pin); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.pin`, *ut.Pin, `^\d{4,6}$`))
		}
	}
	if ut.Pin != nil {
		if utf8.RuneCountInString(*ut.Pin) < 4 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.pin`, *ut.Pin, utf8.RuneCountInString(*ut.Pin), 4, true))
		}
	}
	if ut.Pin != nil {
		if utf8.RuneCountInString(*ut.Pin) > 6 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.pin`, *ut.Pin, utf8.RuneCountInString(*ut.Pin), 6, false))
		}
	}
	return
}

// Publicize creates PinPayload from pinPayload
func (ut *pinPayload) Publicize() *PinPayload {
	var pub PinPayload
	if ut.Pin != nil {
		pub.Pin = *ut.Pin
	}
	return &pub
}

// PinPayload user type.
type PinPayload struct {
	Pin string `form:"pin" json:"pin" yaml:"pin" xml:"pin"`
}

// Validate validates the PinPayload type instance.
func (ut *PinPayload) Validate() (err error) {
	if ut.Pin == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "pin"))
	}
	if ok := goa.ValidatePattern(`^\d{4,6}$`, ut.Pin); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.pin`, ut.Pin, `^\d{4,6}$`))
	}
	if utf8.RuneCountInString(ut.Pin) < 4 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.pin`, ut.Pin, utf8.RuneCountInString(ut.Pin), 4, true))
	}
	if utf8.RuneCountInString(ut.Pin) > 6 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.pin`, ut.Pin, utf8.RuneCountInString(ut.Pin), 6, false))
	}
	return
}

// updateAddressPayload user type.
type updateAddressPayload struct {
	Country     *string `form:"country,omitempty" json:"country,omitempty" yaml:"country,omitempty" xml:"country,omitempty"`
	HouseNumber *string `form:"house_number,omitempty" json:"house_number,omitempty" yaml:"house_number,omitempty" xml:"house_number,omitempty"`
	Location    *string `form:"location,omitempty" json:"location,omitempty" yaml:"location,omitempty" xml:"location,omitempty"`
	Street      *string `form:"street,omitempty" json:"street,omitempty" yaml:"street,omitempty" xml:"street,omitempty"`
	Zipcode     *string `form:"zipcode,omitempty" json:"zipcode,omitempty" yaml:"zipcode,omitempty" xml:"zipcode,omitempty"`
}

// Publicize creates UpdateAddressPayload from updateAddressPayload
func (ut *updateAddressPayload) Publicize() *UpdateAddressPayload {
	var pub UpdateAddressPayload
	if ut.Country != nil {
		pub.Country = ut.Country
	}
	if ut.HouseNumber != nil {
		pub.HouseNumber = ut.HouseNumber
	}
	if ut.Location != nil {
		pub.Location = ut.Location
	}
	if ut.Street != nil {
		pub.Street = ut.Street
	}
	if ut.Zipcode != nil {
		pub.Zipcode = ut.Zipcode
	}
	return &pub
}

// UpdateAddressPayload user type.
type UpdateAddressPayload struct {
	Country     *string `form:"country,omitempty" json:"country,omitempty" yaml:"country,omitempty" xml:"country,omitempty"`
	HouseNumber *string `form:"house_number,omitempty" json:"house_number,omitempty" yaml:"house_number,omitempty" xml:"house_number,omitempty"`
	Location    *string `form:"location,omitempty" json:"location,omitempty" yaml:"location,omitempty" xml:"location,omitempty"`
	Street      *string `form:"street,omitempty" json:"street,omitempty" yaml:"street,omitempty" xml:"street,omitempty"`
	Zipcode     *string `form:"zipcode,omitempty" json:"zipcode,omitempty" yaml:"zipcode,omitempty" xml:"zipcode,omitempty"`
}

// updateUserPayload user type.
type updateUserPayload struct {
	DateOfBirth    *time.Time `form:"date_of_birth,omitempty" json:"date_of_birth,omitempty" yaml:"date_of_birth,omitempty" xml:"date_of_birth,omitempty"`
	Email          *string    `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	FirstName      *string    `form:"first_name,omitempty" json:"first_name,omitempty" yaml:"first_name,omitempty" xml:"first_name,omitempty"`
	Gender         *int       `form:"gender,omitempty" json:"gender,omitempty" yaml:"gender,omitempty" xml:"gender,omitempty"`
	GovIdentifier  *string    `form:"gov_identifier,omitempty" json:"gov_identifier,omitempty" yaml:"gov_identifier,omitempty" xml:"gov_identifier,omitempty"`
	LastName       *string    `form:"last_name,omitempty" json:"last_name,omitempty" yaml:"last_name,omitempty" xml:"last_name,omitempty"`
	LastNamePrefix *string    `form:"last_name_prefix,omitempty" json:"last_name_prefix,omitempty" yaml:"last_name_prefix,omitempty" xml:"last_name_prefix,omitempty"`
	Password       *string    `form:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
	PhoneNumber    *string    `form:"phone_number,omitempty" json:"phone_number,omitempty" yaml:"phone_number,omitempty" xml:"phone_number,omitempty"`
}

// Validate validates the updateUserPayload type instance.
func (ut *updateUserPayload) Validate() (err error) {
	if ut.Email != nil {
		if ok := goa.ValidatePattern(`^[^@\s:]+@[^@\s]+\.[^@\s]+$`, *ut.Email); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.email`, *ut.Email, `^[^@\s:]+@[^@\s]+\.[^@\s]+$`))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 8, true))
		}
	}
	return
}

// Publicize creates UpdateUserPayload from updateUserPayload
func (ut *updateUserPayload) Publicize() *UpdateUserPayload {
	var pub UpdateUserPayload
	if ut.DateOfBirth != nil {
		pub.DateOfBirth = ut.DateOfBirth
	}
	if ut.Email != nil {
		pub.Email = ut.Email
	}
	if ut.FirstName != nil {
		pub.FirstName = ut.FirstName
	}
	if ut.Gender != nil {
		pub.Gender = ut.Gender
	}
	if ut.GovIdentifier != nil {
		pub.GovIdentifier = ut.GovIdentifier
	}
	if ut.LastName != nil {
		pub.LastName = ut.LastName
	}
	if ut.LastNamePrefix != nil {
		pub.LastNamePrefix = ut.LastNamePrefix
	}
	if ut.Password != nil {
		pub.Password = ut.Password
	}
	if ut.PhoneNumber != nil {
		pub.PhoneNumber = ut.PhoneNumber
	}
	return &pub
}

// UpdateUserPayload user type.
type UpdateUserPayload struct {
	DateOfBirth    *time.Time `form:"date_of_birth,omitempty" json:"date_of_birth,omitempty" yaml:"date_of_birth,omitempty" xml:"date_of_birth,omitempty"`
	Email          *string    `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	FirstName      *string    `form:"first_name,omitempty" json:"first_name,omitempty" yaml:"first_name,omitempty" xml:"first_name,omitempty"`
	Gender         *int       `form:"gender,omitempty" json:"gender,omitempty" yaml:"gender,omitempty" xml:"gender,omitempty"`
	GovIdentifier  *string    `form:"gov_identifier,omitempty" json:"gov_identifier,omitempty" yaml:"gov_identifier,omitempty" xml:"gov_identifier,omitempty"`
	LastName       *string    `form:"last_name,omitempty" json:"last_name,omitempty" yaml:"last_name,omitempty" xml:"last_name,omitempty"`
	LastNamePrefix *string    `form:"last_name_prefix,omitempty" json:"last_name_prefix,omitempty" yaml:"last_name_prefix,omitempty" xml:"last_name_prefix,omitempty"`
	Password       *string    `form:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
	PhoneNumber    *string    `form:"phone_number,omitempty" json:"phone_number,omitempty" yaml:"phone_number,omitempty" xml:"phone_number,omitempty"`
}

// Validate validates the UpdateUserPayload type instance.
func (ut *UpdateUserPayload) Validate() (err error) {
	if ut.Email != nil {
		if ok := goa.ValidatePattern(`^[^@\s:]+@[^@\s]+\.[^@\s]+$`, *ut.Email); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`type.email`, *ut.Email, `^[^@\s:]+@[^@\s]+\.[^@\s]+$`))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`type.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 8, true))
		}
	}
	return
}
