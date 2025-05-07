package neo4jschool

import (
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/school"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// RetrieveStudentSchool retrieves a school node that a particular student belongs to in a Neo4j database using the provided studentid and a Neo4jInstance. Returns an error upon failure
func RetrieveStudentSchool(n *neo4jutils.Neo4jInstance, studentid int) (school.School, error) {
	result, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		"MATCH (student:Student {pk: $studentid})-[:STUDENT_AT]->(school:School) RETURN school AS school",
		map[string]any{
			"studentid": studentid, // Bind the mapped studentid data to the "$studentid" parameter
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		return nil, err
	}

	var schoollist []school.School
	// Loop through results and do something with them
	for _, record := range result.Records {
		school, _ := record.Get("school") // .Get() 2nd return is whether key is present
		var s School
		s.Props = school.(neo4j.Node).Props
		schoollist = append(schoollist, s)
	}
	return schoollist[0], nil
}
