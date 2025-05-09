package neo4jproduct

import (
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// CheckProduct checks if a product exists in a Neo4j database using the provided productid and a Neo4jInstance. Returns true if it exists and false if it does not. 
//
// Also returns error if there was a problem with the process of checking the product's existence
func CheckProduct(n *neo4jutils.Neo4jInstance, productid int) (bool, error) {
	result, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		"MATCH (product:Product {pk: $productid}) RETURN product.name AS name",
		map[string]any{
			"productid": productid, // Bind the mapped productid data to the "$productid" parameter
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