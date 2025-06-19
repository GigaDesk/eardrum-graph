package neo4jtransaction

import (
	"time"
)

// represents a transaction queried from a neo4j database, and yet to be formated into a struct
type Transaction struct {
	Props map[string]any
}

// Returns the unique identifier of the transaction
func (t Transaction) GetID() int64 {
	return t.Props["pk"].(int64)
}

// Returns the creation timestamp of the transaction
func (t Transaction) GetCreatedAt() time.Time {
	return t.Props["createdat"].(time.Time)
}

// Returns the last update timestamp of the transaction
func (t Transaction) GetUpdatedAt() time.Time {
	return t.Props["updatedat"].(time.Time)
}

// Returns the last deletion timestamp of the transaction
func (t Transaction) GetDeletedAt() time.Time {
	return t.Props["deletedat"].(time.Time)
}

// Returns the total amount spent in the transaction in cents
func (t Transaction) GetTotalAmountInCents() int64 {
	return t.Props["total_amount_in_cents"].(int64)
}

// Returns the cost of the transaction in cents
func (t Transaction) GetTransactionCostInCents() int64 {
	return t.Props["transaction_cost_in_cents"].(int64)
}
