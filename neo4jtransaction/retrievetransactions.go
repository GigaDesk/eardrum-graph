package neo4jtransaction

import (
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/transaction"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// RetrieveStudentTransactions retrieves transaction nodes that belong to a particular student in a Neo4j database using the provided studentid and a Neo4jInstance. Returns an error upon failure
func RetrieveStudentTransactions(n *neo4jutils.Neo4jInstance, studentid int) ([]transaction.Transaction, error) {
	result, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		"MATCH (transaction)-[:MADE_BY]->(student:Student {pk: $studentid}) RETURN transaction AS transaction",
		map[string]any{
			"studentid": studentid, // Bind the mapped studentid data to the "$studentid" parameter
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		return nil, err
	}

	var transactionlist []transaction.Transaction
	// Loop through results and do something with them
	for _, record := range result.Records {
		student, _ := record.Get("transaction") // .Get() 2nd return is whether key is present
		var t Transaction
		t.Props = student.(neo4j.Node).Props
		transactionlist = append(transactionlist, t)
	}
	return transactionlist, nil
}