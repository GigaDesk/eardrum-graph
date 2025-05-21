package neo4jstudent

import (
	"log"

	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-interfaces/student"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// UpdateStudent updates a student node in a Neo4j database. Returns an error upon failure
func UpdateStudent(n *neo4jutils.Neo4jInstance, s student.Student) error {

	// Construct the Cypher query to update a Student node with the mapped properties
	query := "MATCH (s:Student {pk: $pk}) SET s.updatedat = $updatedat, s.registration_number = $registration_number, s.name = $name, s.phonenumber = $phonenumber, s.date_of_admission = $date_of_admission, s.date_of_birth = $date_of_birth,  s.profile_picture = $profile_picture, s.account_balance_in_cents = $account_balance_in_cents"
	_, err := neo4j.ExecuteQuery(n.Ctx, n.Driver,
		query,
		map[string]any{
			"pk":                       s.GetID(),
			"updatedat":                s.GetUpdatedAt(),
			"registration_number":      s.GetRegistrationNumber(),
			"name":                     s.GetName(),
			"phonenumber":              s.GetPhoneNumber(),
			"date_of_birth":            s.GetDateofBirth(),
			"date_of_admission":        s.GetDateOfAdmission(),
			"profile_picture":          s.GetProfilePicture(),
			"account_balance_in_cents": s.GetAccountBalanceInCents(),
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(n.Db))
	if err != nil {
		log.Println(err)
	}
	return err
}
