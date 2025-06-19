package neo4jshop

import (
	"log"

	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/shop"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func mapShop(s shop.Shop, m map[string]any) {
	m["pk"] = s.GetID()
	m["createdat"] = s.GetCreatedAt()
	m["updatedat"] = s.GetUpdatedAt()
	m["deletedat"] = s.GetDeletedAt()
	m["name"] = s.GetName()
	m["phonenumber"] = s.GetPhoneNumber()
	m["account_balance_in_cents"] = s.GetAccountBalanceInCents()
	m["password"] = s.GetPassword()
	m["pin_code"] = s.GetPinCode()
}

// CreateShop creates new shop nodes in a Neo4j database. Returns an error upon failure
func CreateShop(n *neo4jutils.Neo4jInstance, s shop.Shop ) error {
	m := make(map[string]any)

	mapShop(s, m) // Map shop data to the m map

	shop := m

	// Log the mapped shop data for debugging purposes
	log.Println("creating neo4j shop: ", shop)
	// Construct the Cypher query to create a new shop node with the mapped properties
	query := "CREATE (s:Shop $shop)"
	_, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		query,
		map[string]any{
			"shop":  shop,  // Bind the mapped shop data to the "$shop" parameter
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		log.Println(err)
	}
	return err
}
