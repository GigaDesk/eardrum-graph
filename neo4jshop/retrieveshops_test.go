package neo4jshop

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/GigaDesk/eardrum-graph/mockschool"
	"github.com/GigaDesk/eardrum-graph/mockshop"
	"github.com/GigaDesk/eardrum-graph/neo4jschool"
	"github.com/joho/godotenv"
)

func TestRetrieveSchoolShop(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)
	//create a school node of primary key 1
	neo4jschool.CreateSchool(&neo4jInstance, mockschool.SchoolNode)
	//create shop nodes of primary keys 3,4,5 and 6 to school node of primary key 1
	for _,shop := range mockshop.MultipleShopNodes {
		if err := CreateShop(&neo4jInstance, shop, 1); err != nil {
			log.Fatal("Failed to add a shop node")
		}
	}
	shops, error := RetrieveSchoolShops(&neo4jInstance,1)
    if error != nil {
		t.Error(`error retrieving school shops`)
	}

	if len(shops) != 4{
		t.Error(fmt.Sprintf("length of retrieved shops array is not: %d", 4))
	}

	for _,s:=range shops{
	switch s.GetID(){
	case 3:
		log.Println(fmt.Sprintf("found shop with id %d:", 3))
	case 4:
		log.Println(fmt.Sprintf("found shop with id %d:", 4))
	case 5:
		log.Println(fmt.Sprintf("found shop with id %d:", 5))
	case 6:
		log.Println(fmt.Sprintf("found shop with id %d:", 6))
	default:
		t.Error(fmt.Sprintf("found shop with id %d:", s.GetID()))
	}
	}
}