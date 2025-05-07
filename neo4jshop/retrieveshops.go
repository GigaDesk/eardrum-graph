package neo4jshop

import (
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/shop"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// RetrieveSchoolShops retrieves shop nodes that belong to a particular school in a Neo4j database using the provided schoolid and a Neo4jInstance. Returns an error upon failure
func RetrieveSchoolShops(n *neo4jutils.Neo4jInstance, schoolid int) ([]shop.Shop, error) {
	result, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		"MATCH (shop)-[:SHOP_AT]->(school:School {pk: $schoolid}) RETURN shop AS shop",
		map[string]any{
			"schoolid": schoolid, // Bind the mapped schoolid data to the "$schoolid" parameter
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		return nil, err
	}

	var shoplist []shop.Shop
	// Loop through results and do something with them
	for _, record := range result.Records {
		student, _ := record.Get("shop") // .Get() 2nd return is whether key is present
		var s Shop
		s.Props = student.(neo4j.Node).Props
		shoplist = append(shoplist, s)
	}
	return shoplist, nil
}