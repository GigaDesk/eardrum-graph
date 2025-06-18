package neo4jproduct

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/GigaDesk/eardrum-graph/mockshop"
	"github.com/GigaDesk/eardrum-graph/mockproduct"
	"github.com/GigaDesk/eardrum-graph/neo4jshop"
	"github.com/joho/godotenv"
)


func TestCreateCategoryNode(t *testing.T) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)
	//create shop nodes of primary keys 3,4,5 and 6
	for _, shop := range mockshop.MultipleShopNodes {
		if err := neo4jshop.CreateShop(&neo4jInstance, shop); err != nil {
			t.Error(`Failed to add a shop node`)
		}
		result, error, _ := neo4jshop.CheckShop(&neo4jInstance, int(shop.GetID()))
		if error != nil {
			log.Fatal("Failed to check shop")
		}
		if !result {
			t.Error(fmt.Sprintf("shop of id %d is not available", shop.GetID()))
		}
	}

	//create category nodes of primary keys 3,4,5 and 6 to shop node of primary key 5
	for _, category := range mockproduct.MultipleCategoryNodes {
		if err := CreateCategory(&neo4jInstance, category, 5); err != nil {
			t.Error(`Failed to add a category node`)
		}
		result, error, _ := CheckCategory(&neo4jInstance, int(category.GetID()))
		if error != nil {
			log.Fatal("Failed to check category")
		}
		if !result {
			t.Error(fmt.Sprintf("category of id %d is not available", category.GetID()))
		}
	}
}