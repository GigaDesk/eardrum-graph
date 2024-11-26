package mockstudent

import "time"


type MockStudent struct {
	Id          int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	RegistrationNumber string
	Name string
	PhoneNumber string
	Password string
	DateOfAdmission time.Time
	DateofBirth time.Time
	ProfilePicture string
}




func (m MockStudent) GetID() int {
	return m.Id
}
func (m MockStudent) GetCreatedAt() time.Time {
	return m.CreatedAt
}
func (m MockStudent) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}
func (m MockStudent) GetRegistrationNumber() string {
	return m.RegistrationNumber
}
func (m MockStudent) GetName() string {
	return m.Name
}
func (m MockStudent) GetPhoneNumber() string {
	return m.PhoneNumber
}
func (m MockStudent) GetPassword() string {
	return m.Password
}
func (m MockStudent) GetDateOfAdmission() time.Time {
	return m.DateOfAdmission
}
func (m MockStudent) GetDateofBirth() time.Time {
	return m.DateofBirth
}
func (m MockStudent) GetProfilePicture() string {
	return m.ProfilePicture
}