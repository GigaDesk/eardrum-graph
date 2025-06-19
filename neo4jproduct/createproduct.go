package neo4jproduct

import (
	"log"

	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/product"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func mapProduct(p product.Product, m map[string]any) {
	m["pk"] = p.GetID()
	m["createdat"] = p.GetCreatedAt()
	m["updatedat"] = p.GetUpdatedAt()
	m["deletedat"] = p.GetDeletedAt()
	m["name"] = p.GetName()
	m["price_per_unit_in_cents"] = p.GetPricePerUnitInCents()
}

// CreateProduct creates new product nodes in a Neo4j database. Returns an error upon failure
//
// Note that it is recommended to check if the shop you are adding the product to is available in the database. In rare cases the shop might not exist and this function will not throw an error
//
// Use the function:
//
//	neo4jshop.CheckShop(n *neo4jutils.Neo4jInstance, shopid int) (bool, error)
func CreateProduct(n *neo4jutils.Neo4jInstance, p product.Product, shopid int) error {
	m := make(map[string]any)

	mapProduct(p, m) // Map product data to the m map

	product := m

	// Log the mapped product data for debugging purposes
	log.Println("creating neo4j product: ", product)
	// Construct the Cypher query to create a new product node with the mapped properties
	query := "MATCH (shop:Shop {pk: $shopid}) CREATE (p:Product $product) CREATE (p)-[r:SOLD_AT]->(shop)"
	_, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		query,
		map[string]any{
			"shopid": shopid, // Bind the mapped shopid data to the "$shopid" parameter
			"product": product,  // Bind the mapped product data to the "$product" parameter
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		log.Println(err)
	}
	return err
}