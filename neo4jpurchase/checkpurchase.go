package neo4jpurchase

import (
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// CheckPurchase checks if a purchase exists in a Neo4j database using the provided purchaseid and a Neo4jInstance. Returns true if it exists and false if it does not. 
//
// Also returns error if there was a problem with the process of checking the purchase's existence
func CheckPurchase(n *neo4jutils.Neo4jInstance, purchaseid int) (bool, error) {
	result, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		"MATCH (purchase:Purchase {pk: $purchaseid}) RETURN purchase.total_amount_in_cents AS total_amount_in_cents",
		map[string]any{
			"purchaseid": purchaseid, // Bind the mapped purchaseid data to the "$purchaseid" parameter
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