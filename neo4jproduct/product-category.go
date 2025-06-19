package neo4jproduct

import (
	"log"

	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// AddProductToCategory assigns a product to a specific Category in a Neo4j database. Returns an error upon failure
//
// Note that it is recommended to check if the product and the category you are linking are available in the database. In rare cases the product and the category might not exist and this function will not throw an error
//
// Use the functions:
//
//		neo4jproduct.CheckProduct(n *neo4jutils.Neo4jInstance, productid int) (bool, error)
//	 neo4jproduct.CheckCategory(n *neo4jutils.Neo4jInstance, categoryid int) (bool, error)
func AddProductToCategory(n *neo4jutils.Neo4jInstance, productid int, categoryid int) error {
	//log the productid and categoryid arguments for debugging purposes
	log.Println("connecting product of pk: ", productid, "to category of pk: ", categoryid)
	// Construct the Cypher query to create a new BELONGS_TO relationship between the product node and the category node
	query := "MATCH (product:Product {pk: $productid}) MATCH (category:Category {pk: $categoryid}) MERGE (product)-[:BELONGS_TO]->(category)"
	_, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		query,
		map[string]any{
			"productid":  productid,  // Bind the mapped productid data to the "$productid" parameter
			"categoryid": categoryid, // Bind the mapped categoryid data to the "$categoryid" parameter
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		log.Println(err)
	}
	return err
}

// RemoveProductFromCategory removes a product from a specific Category in a Neo4j database. Returns an error upon failure

func RemoveProductFromCategory(n *neo4jutils.Neo4jInstance, productid int, categoryid int) error {
	//log the productid and categoryid arguments for debugging purposes
	log.Println("removing product of pk: ", productid, "from category of pk: ", categoryid)
	// Construct the Cypher query to delete a BELONGS_TO relationship between the product node and the category node
	query := "MATCH (product:Product {pk: $productid})-[r:BELONGS_TO]->(category:Category {pk: $categoryid}) DELETE r"
	_, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		query,
		map[string]any{
			"productid":  productid,  // Bind the mapped productid data to the "$productid" parameter
			"categoryid": categoryid, // Bind the mapped categoryid data to the "$categoryid" parameter
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		log.Println(err)
	}
	return err
}
