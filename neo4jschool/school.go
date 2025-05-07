package neo4jschool

import (
	"time"
)

//represents a school queried from a neo4j database, and yet to be formated into a struct
type School struct{
	Props map[string] any
} 

// Returns the unique identifier of the school
func (s School) GetID() int64 {
	return s.Props["pk"].(int64)
}

// Returns the creation timestamp of the school
func (s School) GetCreatedAt() time.Time {
	return s.Props["createdat"].(time.Time)
}

// Returns the last update timestamp of the school
func (s School) GetUpdatedAt() time.Time {
	return s.Props["updatedat"].(time.Time)
}

// Returns the name of the school
func (s School) GetName() string  {
	return s.Props["name"].(string)
}

// Returns the phone number of the school
func (s School) GetPhoneNumber() string  {
	return s.Props["phonenumber"].(string)
}

// Returns the password of the school
func (s School) GetPassword() string  {
	return ""
}

// Returns the school's badge
func (s School) GetBadge()  string   {
	return s.Props["badge"].(string)
}

// Returns the school's website
func (s School) GetWebsite() string   {
	return s.Props["website"].(string)
}

// Returns the last deletion timestamp of the school
func (s School) GetDeletedAt() time.Time {
	return time.Time{}
}