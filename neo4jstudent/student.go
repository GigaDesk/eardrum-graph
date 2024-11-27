package neo4jstudent

import (
	"time"
)

//represents students queried from a neo4j database, and yet to be formated into a struct
type Student struct{
	Props map[string] any
} 
 // Returns the unique identifier of the student
func (s Student) GetID() int64 {
	return s.Props["pk"].(int64)
}

// Returns the creation timestamp of the student
func (s Student) GetCreatedAt() time.Time {
	return s.Props["createdat"].(time.Time)
}

// Returns the last update timestamp of the student
func (s Student) GetUpdatedAt() time.Time {
	return s.Props["updatedat"].(time.Time)
}

// Returns the student's registration number
func (s Student) GetRegistrationNumber() string {
	return s.Props["registration_number"].(string)
}

// Returns the name of the student
func (s Student) GetName() string  {
	return s.Props["name"].(string)
}

// Returns the phone number of the student
func (s Student) GetPhoneNumber() string  {
	return s.Props["phonenumber"].(string)
}

// Returns the password associated with the student (e.g., for student access)
func (s Student) GetPassword() string {
	return s.Props["password"].(string)
}

// Returns the student's date of admission
func (s Student) GetDateOfAdmission()  time.Time   {
	return s.Props["date_of_admission"].(time.Time)
}

// Returns the student's birthday
func (s Student) GetDateofBirth() time.Time   {
	return s.Props["date_of_birth"].(time.Time)
}

//Returns the student profile picture's image url
func (s Student) GetProfilePicture() string  {
	return s.Props["profile_picture"].(string)
}