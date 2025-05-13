package neo4jtransaction

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/GigaDesk/eardrum-graph/mockschool"
	"github.com/GigaDesk/eardrum-graph/mockstudent"
	"github.com/GigaDesk/eardrum-graph/mocktransaction"
	"github.com/GigaDesk/eardrum-graph/neo4jschool"
	"github.com/GigaDesk/eardrum-graph/neo4jstudent"
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/joho/godotenv"
)

var (
	neo4jInstance neo4jutils.Neo4jInstance
)

func TestCreateTransactionNode(t *testing.T) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)
	//create a school node of primary key 1
	neo4jschool.CreateSchool(&neo4jInstance, mockschool.SchoolNode)
	//create shop nodes of primary keys 3,4,5 and 6 to school node of primary key 1

	//create a student nodes of primary keys 3,4,5 and 6 to school node of primary key 1
	for _, student := range mockstudent.MultipleStudentNodes {
		if err := neo4jstudent.CreateStudent(&neo4jInstance, student, 1); err != nil {
			t.Error(`Failed to add a student node`)
		}
		result, error := neo4jstudent.CheckStudent(&neo4jInstance, int(student.GetID()))
		if error != nil {
			log.Fatal("Failed to check student")
		}
		if !result {
			t.Error(fmt.Sprintf("student of id %d is not available", student.GetID()))
		}
	}
	//create transaction nodes of primary keys 3,4,5 and 6 to student node of primary key 3
	for _, transaction := range mocktransaction.MultipleTransactionNodes {
		if err := CreateTransaction(&neo4jInstance, transaction, 3); err != nil {
			t.Error(`Failed to add a transaction node`)
		}
		result, error := CheckTransaction(&neo4jInstance, int(transaction.GetID()))
		if error != nil {
			log.Fatal("Failed to check transaction")
		}
		if !result {
			t.Error(fmt.Sprintf("transaction of id %d is not available", transaction.GetID()))
		}
	}
}
