package neo4jstudent

import (
	"log"

	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/user"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// UpdateUser updates a user node in a Neo4j database. Returns an error upon failure
func UpdateUser(n *neo4jutils.Neo4jInstance, s user.User) error {

	// Construct the Cypher query to update a user node with the mapped properties
	query := "MATCH (s:User {pk: $pk}) SET s.updatedat = $updatedat, s.deletedat = $deletedat, s.name = $name, s.phonenumber = $phonenumber, s.account_balance_in_cents = $account_balance_in_cents, s.password = password, s.pin_code = pin_code"
	_, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		query,
		map[string]any{
			"pk":                       s.GetID(),
			"updatedat":                s.GetUpdatedAt(),
			"deletedat":                s.GetDeletedAt(),
			"name":                     s.GetName(),
			"phonenumber":              s.GetPhoneNumber(),
			"password":                 s.GetPassword(),
			"account_balance_in_cents": s.GetAccountBalanceInCents(),
			"pin_code":                 s.GetPinCode(),
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		log.Println(err)
	}
	return err
}
