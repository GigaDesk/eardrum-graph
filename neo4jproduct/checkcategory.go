package neo4jproduct

import (
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/product"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// CheckCategory checks if a product category exists in a Neo4j database using the provided categoryid and a Neo4jInstance. Returns true if it exists and false if it does not.
// Also returns error if there was a problem with the process of checking the category's existence
//
// Also returns a retrieved category record
func CheckCategory(n *neo4jutils.Neo4jInstance, categoryid int) (bool, error, product.Category) {
	result, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		"MATCH (category:Category {pk: $categoryid}) RETURN category AS category",
		map[string]any{
			"categoryid": categoryid, // Bind the mapped categoryid data to the "$categoryid" parameter
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		return false, err, nil
	}
	
	if len(result.Records) == 0 {
		return false, nil, nil
	}

	var categorylist []product.Category
	// Loop through results and do something with them
	for _, record := range result.Records {
		category, _ := record.Get("category") // .Get() 2nd return is whether key is present
		var c Category
		c.Props = category.(neo4j.Node).Props
		categorylist = append(categorylist, c)
	}
	return true, nil, categorylist[0]
}