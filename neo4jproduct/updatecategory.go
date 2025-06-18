package neo4jproduct

import (
	"log"

	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/product"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// UpdateCategory updates a category node in a Neo4j database. Returns an error upon failure
func UpdateCategory(n *neo4jutils.Neo4jInstance, c product.Category) error {

	// Construct the Cypher query to update a Category node with the mapped properties
	query := "MATCH (c:Category {pk: $pk}) SET c.updatedat = $updatedat, c.deletedat = $deletedat, c.name = $name, c.description = $description"
	_, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		query,
		map[string]any{
			"pk":          c.GetID(),
			"updatedat":   c.GetUpdatedAt(),
			"deletedat": c.GetDeletedAt(),
			"name":        c.GetName(),
			"description": c.GetDescription(),
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		log.Println(err)
	}
	return err
}
