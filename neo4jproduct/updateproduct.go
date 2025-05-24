package neo4jproduct

import (
	"log"

	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/product"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// UpdateProduct updates a product node in a Neo4j database. Returns an error upon failure
func UpdateProduct(n *neo4jutils.Neo4jInstance, p product.Product) error {

	// Construct the Cypher query to update a Product node with the mapped properties
	query := "MATCH (p:Product {pk: $pk}) SET p.updatedat = $updatedat, p.name = $name, p.price_per_unit_in_cents = $price_per_unit_in_cents"
	_, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		query,
		map[string]any{
			"pk":                      p.GetID(),
			"updatedat":               p.GetUpdatedAt(),
			"name":                    p.GetName(),
			"price_per_unit_in_cents": p.GetPricePerUnitInCents(),
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		log.Println(err)
	}
	return err
}
