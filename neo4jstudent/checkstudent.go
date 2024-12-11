package neo4jstudent

import (
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// CheckStudent checks if a student exists in a Neo4j database using the provided studentid and a Neo4jInstance. Returns true if it exists and false if it does not. 
//
// Also returns error if there was a problem with the process of checking the student's existence
func CheckStudent(n *neo4jutils.Neo4jInstance, studentid int) (bool, error) {
	result, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		"MATCH (student:Student {pk: $studentid}) RETURN student.name AS name",
		map[string]any{
			"studentid": studentid, // Bind the mapped studentid data to the "$studentid" parameter
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