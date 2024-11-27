package neo4jstudent

import (
	"log"

	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// RetrieveStudents retrieves student nodes that belong to a particular school in a Neo4j database using the provided schoolid and a Neo4jInstance. Returns an error upon failure
func RetrieveSchoolStudents(n *neo4jutils.Neo4jInstance, schoolid int) ([]Neo4jStudent, error) {
	result, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		"MATCH (student)-[:STUDENT_AT]->(school:School {pk: $schoolid}) RETURN student AS student",
		map[string]any{
			"schoolid": schoolid, // Bind the mapped schoolid data to the "$schoolid" parameter
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		return nil, err
	}

	var studentlist []Neo4jStudent
	// Loop through results and do something with them
	for _, record := range result.Records {
		student, _ := record.Get("student") // .Get() 2nd return is whether key is present
		var s Student
		s.Props = student.(neo4j.Node).Props
		studentlist = append(studentlist, s)

		log.Println("added", s.Props["name"])
	}
	return studentlist, nil
}
