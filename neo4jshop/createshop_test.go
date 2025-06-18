package neo4jshop

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/GigaDesk/eardrum-graph/mockshop"
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/joho/godotenv"
)

var (
	neo4jInstance neo4jutils.Neo4jInstance
)

func TestCreateShopNode(t *testing.T) {
	
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)
	//create shop nodes of primary keys 3,4,5 and 6
	for _,shop := range mockshop.MultipleShopNodes {
		if err := CreateShop(&neo4jInstance, shop); err != nil {
			t.Error(`Failed to add a shop node`)
		}
		result, error, _ := CheckShop(&neo4jInstance, int(shop.GetID())) 
		if error != nil {
			log.Fatal("Failed to check shop")
		}
		if !result{
			t.Error(fmt.Sprintf("shop of id %d is not available", shop.GetID()))
		}
	}
}