package neo4jshop

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/GigaDesk/eardrum-graph/mockshop"
	"github.com/joho/godotenv"
)

func TestUpdateShopNode(t *testing.T) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)
	//create shop nodes of primary keys 3,4,5 and 6
	for _, shop := range mockshop.MultipleShopNodes {
		if err := CreateShop(&neo4jInstance, shop); err != nil {
			t.Error(`Failed to add a shop node`)
		}
		result, error, _ := CheckShop(&neo4jInstance, int(shop.GetID()))
		if error != nil {
			log.Fatal("Failed to check shop")
		}
		if !result {
			t.Error(fmt.Sprintf("shop of id %d is not available", shop.GetID()))
		}
	}

	// update
	if err := UpdateShop(&neo4jInstance, mockshop.UpdatedShop); err != nil {
		t.Error(err)
	}

	//check changes
	_, err, shop := CheckShop(&neo4jInstance, int(mockshop.UpdatedShop.GetID()))

	if err != nil {
		t.Error(err)
	}



	//throw errors for unupdated name
	if shop.GetName() != mockshop.UpdatedShop.GetName() {
		t.Error("name is ", shop.GetName(), "instead of ", mockshop.UpdatedShop.GetName())
	}

	//throw errors for unupdated phone number
	if shop.GetPhoneNumber() != mockshop.UpdatedShop.GetPhoneNumber() {
		t.Error("phone number is ", shop.GetPhoneNumber(), "instead of ", mockshop.UpdatedShop.GetPhoneNumber())
	}


	//throw errors for unupdated account balance
	if shop.GetAccountBalanceInCents() != mockshop.UpdatedShop.GetAccountBalanceInCents() {
		t.Error("account balance is ", shop.GetAccountBalanceInCents(), "instead of ", mockshop.UpdatedShop.GetAccountBalanceInCents())
	}

}
