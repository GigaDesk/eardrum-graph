package neo4jproduct

import (
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/product"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// RetrieveShopProducts retrieves product nodes that belong to a particular shop in a Neo4j database using the provided shopid and a Neo4jInstance. Returns an error upon failure
func RetrieveShopProducts(n *neo4jutils.Neo4jInstance, shopid int) ([]product.Product, error) {
	result, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		"MATCH (product)-[:SOLD_AT]->(shop:Shop {pk: $shopid}) RETURN product AS product",
		map[string]any{
			"shopid": shopid, // Bind the mapped shopid data to the "$shopid" parameter
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		return nil, err
	}

	var productlist []product.Product
	// Loop through results and do something with them
	for _, record := range result.Records {
		product, _ := record.Get("product") // .Get() 2nd return is whether key is present
		var p Product
		p.Props = product.(neo4j.Node).Props
		productlist = append(productlist, p)
	}
	return productlist, nil
}


// RetrievePurchasesProduct retrieves a product node that belongs to a particular purchase in a Neo4j database using the provided purchaseid and a Neo4jInstance. Returns an error upon failure
func RetrievePurchaseProduct(n *neo4jutils.Neo4jInstance, purchaseid int) (product.Product, error) {
	result, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		"MATCH (purchase:Purchase {pk: $purchaseid})-[:INVOLVES_PRODUCT]->(product:Product) RETURN product AS product",
		map[string]any{
			"purchaseid": purchaseid, // Bind the mapped purchaseid data to the "$purchaseid" parameter
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		return nil, err
	}

	var productlist []product.Product
	// Loop through results and do something with them
	for _, record := range result.Records {
		product, _ := record.Get("product") // .Get() 2nd return is whether key is present
		var p Product
		p.Props = product.(neo4j.Node).Props
		productlist = append(productlist, p)
	}
	return productlist[0], nil
}