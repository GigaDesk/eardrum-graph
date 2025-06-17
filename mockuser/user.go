package mockuser

import "time"

type MockUser struct {
	Id                       int64
	CreatedAt                time.Time
	UpdatedAt                time.Time
	DeletedAt                time.Time
	Name                     string
	PhoneNumber              string
	Password                 string
	Account_balance_in_cents int64
	PinCode                  string
}

func (m MockUser) GetID() int64 {
	return m.Id
}
func (m MockUser) GetCreatedAt() time.Time {
	return m.CreatedAt
}
func (m MockUser) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}
func (m MockUser) GetDeletedAt() time.Time {
	return m.DeletedAt
}
func (m MockUser) GetName() string {
	return m.Name
}
func (m MockUser) GetPhoneNumber() string {
	return m.PhoneNumber
}
func (m MockUser) GetPassword() string {
	return m.Password
}
func (m MockUser) GetAccountBalanceInCents() int64 {
	return m.Account_balance_in_cents
}
func (m MockUser) GetPinCode() string {
	return m.PinCode
}
