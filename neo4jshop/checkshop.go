package neo4jshop

import (
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/shop"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

//CheckShop checks if a shop exists in a Neo4j database. Returns true if it exists and false if it does not.
//Also returns error if there was a problem with the process of checking the shop's existence
//
//Also returns a retrieved shop record
func CheckShop(n *neo4jutils.Neo4jInstance, shopid int) (bool, error, shop.Shop) {
	result, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		"MATCH (shop:Shop {pk: $shopid}) RETURN shop AS shop",
		map[string]any{
			"shopid": shopid, // Bind the mapped shopid data to the "$shopid" parameter
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		return false, err, nil
	}

	if len(result.Records) == 0 {
		return false, nil, nil
	}

	var shoplist []shop.Shop
	// Loop through results and do something with them
	for _, record := range result.Records {
		shop, _ := record.Get("shop") // .Get() 2nd return is whether key is present
		var s Shop
		s.Props = shop.(neo4j.Node).Props
		shoplist = append(shoplist, s)
	}
	return true, nil, shoplist[0]
}