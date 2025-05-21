package neo4jshop

import (
	"log"

	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/shop"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// UpdateShop updates a shop node in a Neo4j database. Returns an error upon failure
func UpdateShop(n *neo4jutils.Neo4jInstance, s shop.Shop) error {

	// Construct the Cypher query to update a Shop node with the mapped properties
	query := "MATCH (s:Shop {pk: $pk}) SET s.updatedat = $updatedat, s.name = $name, s.phonenumber = $phonenumber, s.account_balance_in_cents = $account_balance_in_cents"
	_, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		query,
		map[string]any{
			"pk":                       s.GetID(),
			"updatedat":                s.GetUpdatedAt(),
			"name":                     s.GetName(),
			"phonenumber":              s.GetPhoneNumber(),
			"account_balance_in_cents": s.GetAccountBalanceInCents(),
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		log.Println(err)
	}
	return err
}
