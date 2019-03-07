// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "MijnApp": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/mijn-app/mijn-app-backend/design
// --out=$(GOPATH)/src/github.com/mijn-app/mijn-app-backend
// --version=v1.4.1

package app

import (
	"fmt"
	"strings"
)

// AddressHref returns the resource href.
func AddressHref(userID interface{}) string {
	paramuserID := strings.TrimLeftFunc(fmt.Sprintf("%v", userID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v1/users/%v/address", paramuserID)
}

// AvglogHref returns the resource href.
func AvglogHref(avglogID interface{}) string {
	paramavglogID := strings.TrimLeftFunc(fmt.Sprintf("%v", avglogID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v1/avglogs/%v", paramavglogID)
}

// ContractHref returns the resource href.
func ContractHref(contractID interface{}) string {
	paramcontractID := strings.TrimLeftFunc(fmt.Sprintf("%v", contractID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v1/contracts/%v", paramcontractID)
}

// OrderHref returns the resource href.
func OrderHref(orderID interface{}) string {
	paramorderID := strings.TrimLeftFunc(fmt.Sprintf("%v", orderID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v1/order/%v", paramorderID)
}

// UserHref returns the resource href.
func UserHref(userID interface{}) string {
	paramuserID := strings.TrimLeftFunc(fmt.Sprintf("%v", userID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v1/users/%v", paramuserID)
}