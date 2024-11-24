package neo4jschool

import (
	"log"
	"time"

	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var m = make(map[string]any)

// Neo4jSchool represents a school entity with its properties.
type Neo4jSchool interface {
	GetID() int              // Returns the unique identifier of the school
	GetCreatedAt() time.Time // Returns the creation timestamp of the school
	GetUpdatedAt() time.Time // Returns the last update timestamp of the school
	GetName() string         // Returns the name of the school
	GetPhoneNumber() string  // Returns the phone number of the school
	GetPassword() string     // Returns the password associated with the school (e.g., for admin access)
	GetBadge() string        // Returns a badge or identifier associated with the school
	GetWebsite() string      // Returns the website URL of the school
}

func mapSchool(s Neo4jSchool) {
	m["pk"] = s.GetID()
	m["createdat"] = s.GetCreatedAt()
	m["updatedat"] = s.GetUpdatedAt()
	m["name"] = s.GetName()
	m["phonenumber"] = s.GetPhoneNumber()
	m["password"] = s.GetPassword()
	m["badge"] = s.GetBadge()
	m["website"] = s.GetWebsite()
}

// CreateSchool creates a new school node in a Neo4j database using the provided Neo4jSchool interface and a Neo4jInstance. Returns an error upon failure
func CreateSchool(n *neo4jutils.Neo4jInstance, s Neo4jSchool) error {
	mapSchool(s) // Map school data to the global m map
	
	school := m

	// Log the mapped school data for debugging purposes
	log.Println("creating neo4j school: ", school)

	// Construct the Cypher query to create a new School node with the mapped properties
	query := "CREATE (s:School $school)"
	_, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		query,
		map[string]any{
			"school": school, // Bind the mapped school data to the "$school" parameter
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		log.Println(err)
	}
	return err
}
