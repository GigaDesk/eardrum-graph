package neo4jschool

import (
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/school"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// CheckSchool checks if a school exists in a Neo4j database using the provided schoolid and a Neo4jInstance. Returns true if it exists and false if it does not.
//
// Also returns error if there was a problem with the process of checking the school's existence
//
// Also return the school record itself
func CheckSchool(n *neo4jutils.Neo4jInstance, schoolid int) (bool, error, school.School) {
	result, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		"MATCH (school:School {pk: $schoolid}) RETURN school AS school",
		map[string]any{
			"schoolid": schoolid, // Bind the mapped schoolid data to the "$schoolid" parameter
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		return false, err, nil
	}
	
	
	if len(result.Records) == 0 {
		return false, nil, nil
	}

	var schoollist []school.School
	// Loop through results and do something with them
	for _, record := range result.Records {
		school, _ := record.Get("school") // .Get() 2nd return is whether key is present
		var s School
		s.Props = school.(neo4j.Node).Props
		schoollist = append(schoollist, s)
	}
	return true, nil, schoollist[0]
}
