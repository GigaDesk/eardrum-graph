package neo4jschool

import (
	"log"

	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/school"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func mapSchool(s school.School, m map[string]any) {
	m["pk"] = s.GetID()
	m["createdat"] = s.GetCreatedAt()
	m["updatedat"] = s.GetUpdatedAt()
	m["name"] = s.GetName()
	m["phonenumber"] = s.GetPhoneNumber()
	m["badge"] = s.GetBadge()
	m["website"] = s.GetWebsite()
}

// CreateSchool creates a new school node in a Neo4j database. Returns an error upon failure
func CreateSchool(n *neo4jutils.Neo4jInstance, s school.School) error {
	m := make(map[string]any)
	
	mapSchool(s, m) // Map school data to the global m map
	
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
