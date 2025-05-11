package neo4jpurchase

import (
	"log"

	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/purchase"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// mapPurchase maps purchase properties from the purchase.Purchase interface
// to a map[string]any for use in the Neo4j query.
func mapPurchase(p purchase.Purchase, m map[string]any) {
	m["pk"] = p.GetID()
	m["transaction_id"] = p.GetTransactionID()
	m["product_id"] = p.GetProductID()
	m["units_bought"] = p.GetUnitsBought()
	m["total_amount_in_cents"] = p.GetTotalAmountInCents()
}

// CreatePurchase creates a new purchase node and its relationships in Neo4j.
// It's recommended to check if the transaction and product exist in the database
// before calling this function.
//
// Uses:
//
//	neo4jtransaction.CheckTransaction(n *neo4jutils.Neo4jInstance, transactionid int) (bool, error)
//	neo4jproduct.CheckProduct(n *neo4jutils.Neo4jInstance, productid int) (bool, error)
func CreatePurchase(n *neo4jutils.Neo4jInstance, p purchase.Purchase) error {
	m := make(map[string]any)

	mapPurchase(p, m) // Map purchase data

	purchase := m

	// Log the data being used in the query.  This is very helpful for debugging.
	log.Printf("Creating Neo4j purchase with data: %+v\n", purchase)
	// Construct the Cypher query to create a new purchase node with the mapped properties
	query := `
		MATCH (transaction:Transaction {pk: $transaction_id})
		MATCH (product:Product {pk: $product_id})
		CREATE (p:Purchase $purchase)
		CREATE (transaction)-[:CONTAINS_PURCHASE]->(p)
		CREATE (p)-[:INVOLVES_PRODUCT]->(product)
	`
	// Execute the query.  Check the error!
	_, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		query,
		map[string]any{
			"transaction_id": p.GetTransactionID(), // Bind the mapped transaction_id data to the "$transactionid" parameter
			"product_id":     p.GetProductID(),     // Bind the mapped product_id data to the "$productid" parameter
			"purchase":      purchase,             // Bind the mapped purchase data to the "$purchase" parameter
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		log.Println(err)
	}
	return err
}
