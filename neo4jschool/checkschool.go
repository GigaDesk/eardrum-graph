package neo4jschool

import (
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// CheckSchool checks if a school exists in a Neo4j database using the provided schoolid and a Neo4jInstance. Returns true if it exists and false if it does not. 
//
// Also returns error if there was a problem with the process of checking the school's existence
func CheckSchool(n *neo4jutils.Neo4jInstance, schoolid int) (bool, error) {
	result, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		"MATCH (school:School {pk: $schoolid}) RETURN school.name AS name",
		map[string]any{
			"schoolid": schoolid, // Bind the mapped schoolid data to the "$schoolid" parameter
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		return false, err
	}
	// Loop through results and do something with them
	if len(result.Records) == 0 {
		return false, nil
	}
	return true, nil
}
