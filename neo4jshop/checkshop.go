package neo4jshop

import (
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// CheckShop checks if a shop exists in a Neo4j database. Returns true if it exists and false if it does not. 
//
// Also returns error if there was a problem with the process of checking the shop's existence
func CheckShop(n *neo4jutils.Neo4jInstance, shopid int) (bool, error) {
	result, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		"MATCH (shop:Shop {pk: $shopid}) RETURN shop.name AS name",
		map[string]any{
			"shopid": shopid, // Bind the mapped shopid data to the "$shopid" parameter
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