package neo4jproduct

import (
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/product"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// CheckProduct checks if a product exists in a Neo4j database using the provided productid and a Neo4jInstance. Returns true if it exists and false if it does not.
// Also returns error if there was a problem with the process of checking the product's existence
//
// Also returns a retrieved product record
func CheckProduct(n *neo4jutils.Neo4jInstance, productid int) (bool, error, product.Product) {
	result, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		"MATCH (product:Product {pk: $productid}) RETURN product AS product",
		map[string]any{
			"productid": productid, // Bind the mapped productid data to the "$productid" parameter
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		return false, err, nil
	}
	
	if len(result.Records) == 0 {
		return false, nil, nil
	}

	var productlist []product.Product
	// Loop through results and do something with them
	for _, record := range result.Records {
		product, _ := record.Get("product") // .Get() 2nd return is whether key is present
		var p Product
		p.Props = product.(neo4j.Node).Props
		productlist = append(productlist, p)
	}
	return true, nil, productlist[0]
}