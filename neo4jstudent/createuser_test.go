package neo4jstudent

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/GigaDesk/eardrum-graph/mockuser"
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/joho/godotenv"
)

var (
	neo4jInstance neo4jutils.Neo4jInstance
)

func TestCreateUserNode(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)

	//create user nodes of primary keys 3,4,5 and 6
	for _,user := range mockuser.MultipleUserNodes {
		if err := CreateUser(&neo4jInstance, user); err != nil {
			t.Error(`Failed to add a user node`)
		}
		result, error, _ := CheckUser(&neo4jInstance, int(user.GetID())) 
		if error != nil {
			log.Fatal("Failed to check user")
		}
		if !result{
			t.Error(fmt.Sprintf("user of id %d is not available", user.GetID()))
		}
	}
}
