package neo4jtransaction

import (
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/transaction"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// CheckTransaction checks if a transaction exists in a Neo4j database. Returns true if it exists and false if it does not.
// Also returns error if there was a problem with the process of checking the transactions's existence
//
// Also returns a retrieved transaction record
func CheckTransaction(n *neo4jutils.Neo4jInstance, transactionid int) (bool, error, transaction.Transaction) {
	result, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		"MATCH (transaction:Transaction {pk: $transactionid}) RETURN transaction AS transaction",
		map[string]any{
			"transactionid": transactionid, // Bind the mapped transactionid data to the "$transactionid" parameter
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
		if err != nil {
			return false, err, nil
		}
		if len(result.Records) == 0 {
			return false, nil, nil
		}
	
		var transactionlist []transaction.Transaction
		// Loop through results and do something with them
		for _, record := range result.Records {
			transaction, _ := record.Get("transaction") // .Get() 2nd return is whether key is present
			var t Transaction
			t.Props = transaction.(neo4j.Node).Props
			transactionlist = append(transactionlist, t)
		}
		return true, nil, transactionlist[0]
}