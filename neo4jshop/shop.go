package neo4jshop

import (
	"time"
)

// represents a shop queried from a neo4j database, and yet to be formated into a struct
type Shop struct {
	Props map[string]any
}

// Returns the unique identifier of the shop
func (s Shop) GetID() int64 {
	return s.Props["pk"].(int64)
}

// Returns the creation timestamp of the shop
func (s Shop) GetCreatedAt() time.Time {
	return s.Props["createdat"].(time.Time)
}

// Returns the last update timestamp of the shop
func (s Shop) GetUpdatedAt() time.Time {
	return s.Props["updatedat"].(time.Time)
}

// Returns the last deletion timestamp of the shop
func (s Shop) GetDeletedAt() time.Time {
	return s.Props["deletedat"].(time.Time)
}

// Returns the name of the shop
func (s Shop) GetName() string {
	return s.Props["name"].(string)
}

// Returns the phone number of the shop
func (s Shop) GetPhoneNumber() string {
	return s.Props["phonenumber"].(string)
}

// Returns the shop's account balance in cents
func (s Shop) GetAccountBalanceInCents() int64 {
	return s.Props["account_balance_in_cents"].(int64)
}

// Returns the security pin code of the shop
func (s Shop) GetPinCode() string {
	return s.Props["pin_code"].(string)
}

// Returns the security password of the shop
func (s Shop) GetPassword() string {
	return s.Props["password"].(string)
}