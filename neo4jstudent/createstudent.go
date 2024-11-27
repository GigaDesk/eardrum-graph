package neo4jstudent

import (
	"log"
	"time"

	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var m = make(map[string]any)

// Neo4jStudent represents a student entity with its properties.
type Neo4jStudent interface {
	GetID() int64              // Returns the unique identifier of the student
	GetCreatedAt() time.Time // Returns the creation timestamp of the student
	GetUpdatedAt() time.Time // Returns the last update timestamp of the student
	GetRegistrationNumber() string // Returns the student's registration number
	GetName() string         // Returns the name of the student
	GetPhoneNumber() string  // Returns the phone number of the student
	GetPassword() string     // Returns the password associated with the student (e.g., for student access)
	GetDateOfAdmission()  time.Time       // Returns the student's date of admission
	GetDateofBirth() time.Time     // Returns the student's birthday
	GetProfilePicture() string  //Returns the student profile picture's image url
}

func mapSchool(s Neo4jStudent) {
	m["pk"] = s.GetID()
	m["createdat"] = s.GetCreatedAt()
	m["updatedat"] = s.GetUpdatedAt()
	m["registration_number"]= s.GetRegistrationNumber()
	m["name"] = s.GetName()
	m["phonenumber"] = s.GetPhoneNumber()
	m["password"] = s.GetPassword()
	m["date_of_admission"] = s.GetDateOfAdmission()
	m["date_of_birth"] = s.GetDateofBirth()
	m["profile_picture"] = s.GetProfilePicture()
}

// CreateStudents creates new student nodes in a Neo4j database using the provided Neo4jSchool interface and a Neo4jInstance. Returns an error upon failure
//
//Note that it is recommended to check if the school you are adding the students to is available in the database. In rare cases the school might not exist and this function will not throw an error
//
//Use the function: 
//  school.CheckSchool(n *neo4jutils.Neo4jInstance, schoolid int) (bool, error)
func CreateStudent(n *neo4jutils.Neo4jInstance, s Neo4jStudent, schoolid int) error {
	
	mapSchool(s) // Map student data to the global m map
	
	student := m

	// Log the mapped student data for debugging purposes
	log.Println("creating neo4j student: ", student)
	// Construct the Cypher query to create a new student node with the mapped properties
	query := "MATCH (school:School {pk: $schoolid}) CREATE (s:Student $student) CREATE (s)-[r:STUDENT_AT]->(school)"
	_, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		query,
		map[string]any{
			"schoolid": schoolid, // Bind the mapped schoolid data to the "$schoolid" parameter
			"student": student, // Bind the mapped student data to the "$student" parameter
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		log.Println(err)
	}
	return err
}