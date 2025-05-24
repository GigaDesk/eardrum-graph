package neo4jproduct

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/GigaDesk/eardrum-graph/mockschool"
	"github.com/GigaDesk/eardrum-graph/mockshop"
	"github.com/GigaDesk/eardrum-graph/mockproduct"
	"github.com/GigaDesk/eardrum-graph/neo4jschool"
	"github.com/GigaDesk/eardrum-graph/neo4jshop"
	"github.com/joho/godotenv"
)

func TestUpdateProductNode(t *testing.T) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)
	//create a school node of primary key 1
	neo4jschool.CreateSchool(&neo4jInstance, mockschool.SchoolNode)
	//create shop nodes of primary keys 3,4,5 and 6 to school node of primary key 1
	for _, shop := range mockshop.MultipleShopNodes {
		if err := neo4jshop.CreateShop(&neo4jInstance, shop, 1); err != nil {
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

	//create product nodes of primary keys 3,4,5 and 6 to shop node of primary key 5
	for _, product := range mockproduct.MultipleProductNodes {
		if err := CreateProduct(&neo4jInstance, product, 5); err != nil {
			t.Error(`Failed to add a product node`)
		}
		result, error, _ := CheckProduct(&neo4jInstance, int(product.GetID()))
		if error != nil {
			log.Fatal("Failed to check product")
		}
		if !result {
			t.Error(fmt.Sprintf("product of id %d is not available", product.GetID()))
		}
	}

    // update
	if err := UpdateProduct(&neo4jInstance, mockproduct.UpdatedProduct); err != nil {
		t.Error(err)
	}

	//check changes
	_, err, product := CheckProduct(&neo4jInstance, int(mockproduct.UpdatedProduct.GetID()))

	if err != nil {
		t.Error(err)
	}



	//throw errors for unupdated name
	if product.GetName() != mockproduct.UpdatedProduct.GetName() {
		t.Error("name is ", product.GetName(), "instead of ", mockproduct.UpdatedProduct.GetName())
	}

	//throw errors for unupdated price
	if product.GetPricePerUnitInCents() != mockproduct.UpdatedProduct.GetPricePerUnitInCents() {
		t.Error("account balance is ", product.GetPricePerUnitInCents(), "instead of ", mockproduct.UpdatedProduct.GetPricePerUnitInCents())
	}

}
