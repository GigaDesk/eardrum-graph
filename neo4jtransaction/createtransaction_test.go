package neo4jtransaction

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/GigaDesk/eardrum-graph/mockuser"
	"github.com/GigaDesk/eardrum-graph/mocktransaction"
	"github.com/GigaDesk/eardrum-graph/neo4juser"
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

	//create user nodes of primary keys 3,4,5 and 6
	for _, user := range mockuser.MultipleUserNodes {
		if err := neo4juser.CreateUser(&neo4jInstance, user); err != nil {
			t.Error(`Failed to add a user node`)
		}
		result, error, _ := neo4juser.CheckUser(&neo4jInstance, int(user.GetID()))
		if error != nil {
			log.Fatal("Failed to check user")
		}
		if !result {
			t.Error(fmt.Sprintf("user of id %d is not available", user.GetID()))
		}
	}
	//create transaction nodes of primary keys 3,4,5 and 6 to user node of primary key 3
	for _, transaction := range mocktransaction.MultipleTransactionNodes {
		if err := CreateTransaction(&neo4jInstance, transaction, 3); err != nil {
			t.Error(`Failed to add a transaction node`)
		}
		result, error, _ := CheckTransaction(&neo4jInstance, int(transaction.GetID()))
		if error != nil {
			log.Fatal("Failed to check transaction")
		}
		if !result {
			t.Error(fmt.Sprintf("transaction of id %d is not available", transaction.GetID()))
		}
	}
}
