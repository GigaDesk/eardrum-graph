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
	return time.Time{}
}

// Returns the student account balance after the transaction
func (t Transaction) GetBalanceAfter() int64 {
	return t.Props["balance_after"].(int64)
}

// Returns the student account balance before the transaction
func (t Transaction) GetBalanceBefore() int64 {
	return t.Props["balance_before"].(int64)
}

// Returns the total amount spent in the transaction
func (t Transaction) GetTotalAmount() int64 {
	return t.Props["total_amount"].(int64)
}
