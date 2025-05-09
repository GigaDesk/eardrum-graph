package neo4jpurchase

import (
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/purchase"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// RetrieveTransactionPurchases retrieves purchase nodes that belong to a particular transaction in a Neo4j database using the provided transactionid and a Neo4jInstance. Returns an error upon failure
func RetrieveTransactionPurchases(n *neo4jutils.Neo4jInstance, transactionid int) ([]purchase.Purchase, error) {
	result, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		"MATCH (transaction:Transaction{pk: $transactionid})-[:CONTAINS_PURCHASE]->(purchase:Purchase) RETURN purchase AS purchase",
		map[string]any{
			"transactionid": transactionid, // Bind the mapped transactionid data to the "$transactionid" parameter
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		return nil, err
	}

	var purchaselist []purchase.Purchase
	// Loop through results and do something with them
	for _, record := range result.Records {
		student, _ := record.Get("purchase") // .Get() 2nd return is whether key is present
		var p Purchase
		p.Props = student.(neo4j.Node).Props
		purchaselist = append(purchaselist, p)
	}
	return purchaselist, nil
}