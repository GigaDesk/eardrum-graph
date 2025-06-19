package neo4jproduct

import "time"

// represents a product category queried from a neo4j database, and yet to be formated into a struct
type Category struct {
	Props map[string]any
}

// Returns the unique identifier of the product category
func (c Category) GetID() int64 {
	return c.Props["pk"].(int64)
}

// Returns the creation timestamp of the product category
func (c Category) GetCreatedAt() time.Time {
	return c.Props["createdat"].(time.Time)
}

// Returns the last update timestamp of the product category
func (c Category) GetUpdatedAt() time.Time {
	return c.Props["updatedat"].(time.Time)
}

// Returns the last deletion timestamp of the product category
func (c Category) GetDeletedAt() time.Time {
	return c.Props["deletedat"].(time.Time)
}

// Returns the name of the product category
func (c Category) GetName() string {
	return c.Props["name"].(string)
}

// Returns the description of the product category
func (c Category) GetDescription() string {
	return c.Props["description"].(string)
}