package neo4juser

import (
	"log"

	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/user"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func mapUser(s user.User, m map[string]any) {
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

// CreateUser creates new user nodes in a Neo4j database. Returns an error upon failure
func CreateUser(n *neo4jutils.Neo4jInstance, s user.User) error {
	m := make(map[string]any)
	
	mapUser(s, m) // Map user data to the m map
	
	user := m

	// Log the mapped user data for debugging purposes
	log.Println("creating neo4j user: ", user)
	// Construct the Cypher query to create a new user node with the mapped properties
	query := "CREATE (u:User $user)"
	_, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		query,
		map[string]any{
			"user": user, // Bind the mapped user data to the "$user" parameter
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		log.Println(err)
	}
	return err
}