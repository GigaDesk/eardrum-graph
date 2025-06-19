package neo4jproduct

import (
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/product"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// RetrieveShopCategories retrieves category nodes that belong to a particular shop in a Neo4j database using the provided shopid and a Neo4jInstance. Returns an error upon failure
func RetrieveShopCategories(n *neo4jutils.Neo4jInstance, shopid int) ([]product.Category, error) {
	result, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		"MATCH (shop:Shop {pk: $shopid})-[:HAS_CATEGORY]->(category:Category) RETURN category AS category",
		map[string]any{
			"shopid": shopid, // Bind the mapped shopid data to the "$shopid" parameter
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		return nil, err
	}

	var categorylist []product.Category
	// Loop through results and do something with them
	for _, record := range result.Records {
		category, _ := record.Get("category") // .Get() 2nd return is whether key is present
		var c Category
		c.Props = category.(neo4j.Node).Props
		categorylist = append(categorylist, c)
	}
	return categorylist, nil
}