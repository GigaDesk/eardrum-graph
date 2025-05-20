package neo4jschool

import (
	"log"

	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/school"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// UpdateSchool updates a school node in a Neo4j database. Returns an error upon failure
func UpdateSchool(n *neo4jutils.Neo4jInstance, s school.School) error {

	// Construct the Cypher query to update a School node with the mapped properties
	query := "MATCH (s:School {pk: $pk}) SET s.updatedat = $updatedat, s.name = $name, s.phonenumber = $phonenumber, s.badge = $badge, s.website = $website"
	_, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		query,
		map[string]any{
			"pk": s.GetID(),
			"updatedat": s.GetUpdatedAt(),
			"name": s.GetName(),
			"phonenumber": s.GetPhoneNumber(),
			"badge": s.GetBadge(),
			"website": s.GetWebsite(),
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		log.Println(err)
	}
	return err
}