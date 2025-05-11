package neo4jtransaction

import (
	"log"

	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/transaction"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func mapTransaction(t transaction.Transaction, m map[string]any) {
	m["pk"] = t.GetID()
	m["createdat"] = t.GetCreatedAt()
	m["updatedat"] = t.GetUpdatedAt()
	m["total_amount_in_cents"] = t.GetTotalAmountInCents()
	m["balance_before_in_cents"] = t.GetBalanceBeforeInCents()
	m["balance_after_in_cents"] = t.GetBalanceAfterInCents()
}

// CreateTransaction creates new transaction nodes in a Neo4j database. Returns an error upon failure
//
// Note that it is recommended to check if the student you are adding the transaction to is available in the database. In rare cases the student might not exist and this function will not throw an error
//
// Use the function:
//
//	neo4jschool.CheckStudent(n *neo4jutils.Neo4jInstance, studentid int) (bool, error)
func CreateTransaction(n *neo4jutils.Neo4jInstance, t transaction.Transaction, studentid int) error {
	m := make(map[string]any)

	mapTransaction(t, m) // Map transaction data to the m map

	transaction := m

	// Log the mapped transaction data for debugging purposes
	log.Println("creating neo4j transaction: ", transaction)
	// Construct the Cypher query to create a new transaction node with the mapped properties
	query := "MATCH (student:Student {pk: $studentid}) CREATE (t:Transaction $transaction) CREATE (t)-[r:MADE_BY]->(student)"
	_, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		query,
		map[string]any{
			"studentid": studentid, // Bind the mapped studentid data to the "$studentid" parameter
			"transaction":  transaction,  // Bind the mapped transaction data to the "$transaction" parameter
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		log.Println(err)
	}
	return err
}