package neo4jtransaction

import (
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// CheckTransaction checks if a transaction exists in a Neo4j database. Returns true if it exists and false if it does not. 
//
// Also returns error if there was a problem with the process of checking the transactions's existence
func CheckTransaction(n *neo4jutils.Neo4jInstance, transactionid int) (bool, error) {
	result, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		"MATCH (transaction:Transaction {pk: $transactionid}) RETURN transaction.pk AS pk",
		map[string]any{
			"transactionid": transactionid, // Bind the mapped transactionid data to the "$transactionid" parameter
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