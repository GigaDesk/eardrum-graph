package neo4jproduct

import (
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// CheckCategory checks if a product category exists in a Neo4j database using the provided categoryid and a Neo4jInstance. Returns true if it exists and false if it does not. 
//
// Also returns error if there was a problem with the process of checking the category's existence
func CheckCategory(n *neo4jutils.Neo4jInstance, categoryid int) (bool, error) {
	result, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		"MATCH (category:Category {pk: $categoryid}) RETURN category.name AS name",
		map[string]any{
			"categoryid": categoryid, // Bind the mapped categoryid data to the "$categoryid" parameter
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		return false, err
	}
	// Loop through results and do something with them
	if len(result.Records) == 0 {
		return false, nil
	}
	return true, nil
}