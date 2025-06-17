package neo4juser

import (
	"time"
)

// represents a user queried from a neo4j database, and yet to be formated into a struct
type User struct {
	Props map[string]any
}

// Returns the unique identifier of the user
func (s User) GetID() int64 {
	return s.Props["pk"].(int64)
}

// Returns the creation timestamp of the user
func (s User) GetCreatedAt() time.Time {
	return s.Props["createdat"].(time.Time)
}

// Returns the last update timestamp of the user
func (s User) GetUpdatedAt() time.Time {
	return s.Props["updatedat"].(time.Time)
}

// Retuurns the last deletion timestamp of the user
func (s  User) GetDeletedAt() time.Time {
	return time.Time{}
}

// Returns the name of the user
func (s User) GetName() string {
	return s.Props["name"].(string)
}

// Returns the phone number of the user
func (s User) GetPhoneNumber() string {
	return s.Props["phonenumber"].(string)
}

// Returns the user's account balance in cents
func (s User) GetAccountBalanceInCents() int64 {
	return s.Props["account_balance_in_cents"].(int64)
}

// Returns the security pin code of the user
func (s User) GetPinCode() string {
	return s.Props["pin_code"].(string)
}

// Returns the security password of the user
func (s User) GetPassword() string {
	return s.Props["password"].(string)
}
