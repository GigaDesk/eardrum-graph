package neo4jstudent

import (
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/user"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// CheckUser checks if a user exists in a Neo4j database using the provided userid and a Neo4jInstance. Returns true if it exists and false if it does not.
//
// Also returns error if there was a problem with the process of checking the user's existence
//
// Also returns the retrieved user record
func CheckUser(n *neo4jutils.Neo4jInstance, userid int) (bool, error, user.User) {
	result, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		"MATCH (user:User {pk: $userid}) RETURN user AS user",
		map[string]any{
			"userid": userid, // Bind the mapped userid data to the "$userid" parameter
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		return false, err, nil
	}

	if len(result.Records) == 0 {
		return false, nil, nil
	}

	var userlist []user.User
	// Loop through results and do something with them
	for _, record := range result.Records {
		user, _ := record.Get("user") // .Get() 2nd return is whether key is present
		var s User
		s.Props = user.(neo4j.Node).Props
		userlist = append(userlist, s)
	}
	return true, nil, userlist[0]
}