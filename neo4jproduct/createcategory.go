package neo4jproduct

import (
	"log"

	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/product"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func mapCategory(p product.Category, m map[string]any) {
	m["pk"] = p.GetID()
	m["createdat"] = p.GetCreatedAt()
	m["updatedat"] = p.GetUpdatedAt()
	m["name"] = p.GetName()
	m["description"] = p.GetDescription()
}

// CreateCategory creates a new category node in a Neo4j database. Returns an error upon failure
//
// Note that it is recommended to check if the shop you are adding the category to is available in the database. In rare cases the shop might not exist and this function will not throw an error
//
// Use the function:
//
//	neo4jshop.CheckShop(n *neo4jutils.Neo4jInstance, shopid int) (bool, error)
func CreateCategory(n *neo4jutils.Neo4jInstance, c product.Category, shopid int) error {
	m := make(map[string]any)

	mapCategory(c, m) // Map category data to the m map

	category := m

	// Log the mapped category data for debugging purposes
	log.Println("creating neo4j category: ", category)
	// Construct the Cypher query to create a new category node with the mapped properties
	query := "MATCH (shop:Shop {pk: $shopid}) CREATE (c:Category $category) CREATE (shop)-[r:HAS_CATEGORY]->(c)"
	_, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		query,
		map[string]any{
			"shopid": shopid, // Bind the mapped shopid data to the "$shopid" parameter
			"category": category,  // Bind the mapped category data to the "$category" parameter
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		log.Println(err)
	}
	return err
}