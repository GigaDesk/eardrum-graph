package neo4jproduct

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/GigaDesk/eardrum-graph/mockproduct"
	"github.com/GigaDesk/eardrum-graph/mockpurchase"
	"github.com/GigaDesk/eardrum-graph/mockschool"
	"github.com/GigaDesk/eardrum-graph/mockshop"
	"github.com/GigaDesk/eardrum-graph/mockstudent"
	"github.com/GigaDesk/eardrum-graph/mocktransaction"
	"github.com/GigaDesk/eardrum-graph/neo4jpurchase"
	"github.com/GigaDesk/eardrum-graph/neo4jschool"
	"github.com/GigaDesk/eardrum-graph/neo4jshop"
	"github.com/GigaDesk/eardrum-graph/neo4jstudent"
	"github.com/GigaDesk/eardrum-graph/neo4jtransaction"
	"github.com/joho/godotenv"
)



func TestRetrieveShopProducts(t *testing.T) {

	
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
		result, error := neo4jshop.CheckShop(&neo4jInstance, int(shop.GetID()))
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
		result, error := CheckProduct(&neo4jInstance, int(product.GetID()))
		if error != nil {
			log.Fatal("Failed to check product")
		}
		if !result {
			t.Error(fmt.Sprintf("product of id %d is not available", product.GetID()))
		}
	}
	products, error := RetrieveShopProducts(&neo4jInstance, 5)
    if error != nil {
		t.Error(`error retrieving shop products`)
	}

	if len(products) != 4{
		t.Error(fmt.Sprintf("length of retrieved products array is not: %d", 4))
	}

	for _,s:=range products{
	switch s.GetID(){
	case 3:
		log.Println(fmt.Sprintf("found product with id %d:", 3))
	case 4:
		log.Println(fmt.Sprintf("found product with id %d:", 4))
	case 5:
		log.Println(fmt.Sprintf("found product with id %d:", 5))
	case 6:
		log.Println(fmt.Sprintf("found product with id %d:", 6))
	default:
		t.Error(fmt.Sprintf("found product with id %d:", s.GetID()))
	}
	}
}

func TestRetrievePurchaseProduct(t *testing.T) {

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


	//create shop nodes of primary keys 3,4,5 and 6 to school node of primary key 1
	for _, shop := range mockshop.MultipleShopNodes {
		if err := neo4jshop.CreateShop(&neo4jInstance, shop, 1); err != nil {
			t.Error(`Failed to add a shop node`)
		}
		result, error := neo4jshop.CheckShop(&neo4jInstance, int(shop.GetID()))
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
		result, error := CheckProduct(&neo4jInstance, int(product.GetID()))
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
		result, error := neo4jtransaction.CheckTransaction(&neo4jInstance, int(transaction.GetID()))
		if error != nil {
			log.Fatal("Failed to check transaction")
		}
		if !result {
			t.Error(fmt.Sprintf("transaction of id %d is not available", transaction.GetID()))
		}
	}

	//create purchase nodes of primary keys 1,2,3 and 4
	for _, purchase := range mockpurchase.MultiplePurchaseNodes {
		if err := neo4jpurchase.CreatePurchase(&neo4jInstance, purchase); err != nil {
			t.Error(`Failed to add a purchase node`)
		}
		result, error := neo4jpurchase.CheckPurchase(&neo4jInstance, int(purchase.GetID()))
		if error != nil {
			log.Fatal("Failed to check purchase")
		}
		if !result {
			t.Error(fmt.Sprintf("purchase of id %d is not available", purchase.GetID()))
		}
	}

	product, error := RetrievePurchaseProduct(&neo4jInstance, 1)
    if error != nil {
		t.Error(`error retrieving purchase products`)
	}

	if product.GetID() != 4{
		t.Error(fmt.Sprintf("Id of retrieved product is not: %d", 4))
	}

	product, error = RetrievePurchaseProduct(&neo4jInstance, 4)
    if error != nil {
		t.Error(`error retrieving purchase product`)
	}

	if product.GetID() != 3{
		t.Error(fmt.Sprintf("Id of retrieved purchase product is not: %d", 3))
	}

	product, error = RetrievePurchaseProduct(&neo4jInstance, 2)
    if error != nil {
		t.Error(`error retrieving purchase product`)
	}

	if product.GetID() != 6{
		t.Error(fmt.Sprintf("Id of retrieved purchase product is not: %d", 6))
	}

	product, error = RetrievePurchaseProduct(&neo4jInstance, 3)
    if error != nil {
		t.Error(`error retrieving purchase product`)
	}

	if product.GetID() != 5{
		t.Error(fmt.Sprintf("Id of retrieved purchase product is not: %d", 5))
	}

}
