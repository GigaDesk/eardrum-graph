package neo4jproduct

import (
	"time"
)

// represents a product queried from a neo4j database, and yet to be formated into a struct
type Product struct {
	Props map[string]any
}

// Returns the unique identifier of the product
func (s Product) GetID() int64 {
	return s.Props["pk"].(int64)
}

// Returns the creation timestamp of the product
func (s Product) GetCreatedAt() time.Time {
	return s.Props["createdat"].(time.Time)
}

// Returns the last update timestamp of the product
func (s Product) GetUpdatedAt() time.Time {
	return s.Props["updatedat"].(time.Time)
}

// Returns the last deletion timestamp of the product
func (m Product) GetDeletedAt() time.Time {
	return time.Time{}
}

// Returns the name of the product
func (s Product) GetName() string {
	return s.Props["name"].(string)
}

// Returns the product's account balance in cents
func (s Product) GetPricePerUnitInCents() int64 {
	return s.Props["price_per_unit_in_cents"].(int64)
}