package neo4jpurchase

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/GigaDesk/eardrum-graph/mockproduct"
	"github.com/GigaDesk/eardrum-graph/mockpurchase"
	"github.com/GigaDesk/eardrum-graph/mockshop"
	"github.com/GigaDesk/eardrum-graph/mocktransaction"
	"github.com/GigaDesk/eardrum-graph/mockuser"
	"github.com/GigaDesk/eardrum-graph/neo4jproduct"
	"github.com/GigaDesk/eardrum-graph/neo4jshop"
	"github.com/GigaDesk/eardrum-graph/neo4jtransaction"
	"github.com/GigaDesk/eardrum-graph/neo4juser"
	"github.com/joho/godotenv"
)


func TestRetrievePurchaseNode(t *testing.T) {

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
		result, error,_ := neo4juser.CheckUser(&neo4jInstance, int(user.GetID()))
		if error != nil {
			log.Fatal("Failed to check user")
		}
		if !result {
			t.Error(fmt.Sprintf("user of id %d is not available", user.GetID()))
		}
	}


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

	//create product nodes of primary keys 3,4,5 and 6 to shop node of primary key 5
	for _, product := range mockproduct.MultipleProductNodes {
		if err := neo4jproduct.CreateProduct(&neo4jInstance, product, 5); err != nil {
			t.Error(`Failed to add a product node`)
		}
		result, error, _ := neo4jproduct.CheckProduct(&neo4jInstance, int(product.GetID()))
		if error != nil {
			log.Fatal("Failed to check product")
		} 
		if !result {
			t.Error(fmt.Sprintf("product of id %d is not available", product.GetID()))
		}
	}

	//create transaction nodes of primary keys 3,4,5 and 6 to student node of primary key 3
	for _, transaction := range mocktransaction.MultipleTransactionNodes {
		if err := neo4jtransaction.CreateTransaction(&neo4jInstance, transaction, 3); err != nil {
			t.Error(`Failed to add a transaction node`)
		}
		result, error, _ := neo4jtransaction.CheckTransaction(&neo4jInstance, int(transaction.GetID()))
		if error != nil {
			log.Fatal("Failed to check transaction")
		}
		if !result {
			t.Error(fmt.Sprintf("transaction of id %d is not available", transaction.GetID()))
		}
	}

	//create purchase nodes of primary keys 1,2,3 and 4
	for _, purchase := range mockpurchase.MultiplePurchaseNodes {
		if err := CreatePurchase(&neo4jInstance, purchase); err != nil {
			t.Error(`Failed to add a purchase node`)
		}
		result, error, _ := CheckPurchase(&neo4jInstance, int(purchase.GetID()))
		if error != nil {
			log.Fatal("Failed to check purchase")
		}
		if !result {
			t.Error(fmt.Sprintf("purchase of id %d is not available", purchase.GetID()))
		}
	}

	purchases, error := RetrieveTransactionPurchases(&neo4jInstance, 5)
    if error != nil {
		t.Error(`error retrieving transaction purchases`)
	}

	if len(purchases) != 2{
		t.Error(fmt.Sprintf("length of retrieved purchases array is not: %d", 2))
	}

	for _,s:=range purchases{
	switch s.GetID(){
	case 1:
		log.Println(fmt.Sprintf("found purchase with id %d:", 1))
	case 2:
		log.Println(fmt.Sprintf("found purchase with id %d:", 2))
	default:
		t.Error(fmt.Sprintf("found product with id %d:", s.GetID()))
	}
	}

	purchases, error = RetrieveTransactionPurchases(&neo4jInstance, 6)
    if error != nil {
		t.Error(`error retrieving transaction purchases`)
	}

	if len(purchases) != 1{
		t.Error(fmt.Sprintf("length of retrieved purchases array is not: %d", 1))
	}

	for _,s:=range purchases{
	switch s.GetID(){
	case 4:
		log.Println(fmt.Sprintf("found purchase with id %d:", 4))
	default:
		t.Error(fmt.Sprintf("found product with id %d:", s.GetID()))
	}
	}

	purchases, error = RetrieveTransactionPurchases(&neo4jInstance, 3)
    if error != nil {
		t.Error(`error retrieving transaction purchases`)
	}

	if len(purchases) != 1{
		t.Error(fmt.Sprintf("length of retrieved purchases array is not: %d", 1))
	}

	for _,s:=range purchases{
	switch s.GetID(){
	case 3:
		log.Println(fmt.Sprintf("found purchase with id %d:", 3))
	default:
		t.Error(fmt.Sprintf("found product with id %d:", s.GetID()))
	}
	}

}
