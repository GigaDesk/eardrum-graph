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
	m["deletedat"] = t.GetDeletedAt()
	m["total_amount_in_cents"] = t.GetTotalAmountInCents()
	m["transaction_cost_in_cents"] = t.GetTransactionCostInCents()
}

// CreateTransaction creates new transaction nodes in a Neo4j database. Returns an error upon failure
//
// Note that it is recommended to check if the user you are adding the transaction to is available in the database. In rare cases the user might not exist and this function will not throw an error
//
// Use the function:
//
//	neo4juser.CheckUser(n *neo4jutils.Neo4jInstance, userid int) (bool, error)
func CreateTransaction(n *neo4jutils.Neo4jInstance, t transaction.Transaction, userid int) error {
	m := make(map[string]any)

	mapTransaction(t, m) // Map transaction data to the m map

	transaction := m

	// Log the mapped transaction data for debugging purposes
	log.Println("creating neo4j transaction: ", transaction)
	// Construct the Cypher query to create a new transaction node with the mapped properties
	query := "MATCH (user:User {pk: $userid}) CREATE (t:Transaction $transaction) CREATE (t)-[r:MADE_BY]->(user)"
	_, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		query,
		map[string]any{
			"userid": userid, // Bind the mapped userid data to the "$userid" parameter
			"transaction":  transaction,  // Bind the mapped transaction data to the "$transaction" parameter
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		log.Println(err)
	}
	return err
}