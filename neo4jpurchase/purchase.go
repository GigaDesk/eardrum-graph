package neo4jpurchase

// represents a purchase queried from a neo4j database, and yet to be formated into a struct
type Purchase struct {
	Props map[string]any
}

// Returns the unique identifier of the purchase
func (p Purchase) GetID() int64 {
	return p.Props["pk"].(int64)
}

// Returns the unique identifier of the product involved in the purchase
func (p Purchase) GetProductID() int64 {
	return p.Props["product_id"].(int64)
}

// Returns the unique identifier of the transaction that contains the purchase
func (p Purchase) GetTransactionID() int64 {
	return p.Props["transaction_id"].(int64)
}

// Returns the total amount spent in the purchase in cents
func (t Purchase) GetTotalAmountInCents() int64 {
	return t.Props["total_amount_in_cents"].(int64)
}

// Returns the number of units of a product in the purchase
func (t Purchase) GetUnitsBought() int {
	return t.Props["units_bought"].(int)
}
