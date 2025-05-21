package neo4jpurchase

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
	"github.com/GigaDesk/eardrum-graph/neo4jproduct"
	"github.com/GigaDesk/eardrum-graph/neo4jschool"
	"github.com/GigaDesk/eardrum-graph/neo4jshop"
	"github.com/GigaDesk/eardrum-graph/neo4jstudent"
	"github.com/GigaDesk/eardrum-graph/neo4jtransaction"
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/joho/godotenv"
)

var (
	neo4jInstance neo4jutils.Neo4jInstance
)

func TestCreatePurchaseNode(t *testing.T) {

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
		result, error, _ := neo4jstudent.CheckStudent(&neo4jInstance, int(student.GetID()))
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
		if err := neo4jproduct.CreateProduct(&neo4jInstance, product, 5); err != nil {
			t.Error(`Failed to add a product node`)
		}
		result, error := neo4jproduct.CheckProduct(&neo4jInstance, int(product.GetID()))
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
		if err := CreatePurchase(&neo4jInstance, purchase); err != nil {
			t.Error(`Failed to add a purchase node`)
		}
		result, error := CheckPurchase(&neo4jInstance, int(purchase.GetID()))
		if error != nil {
			log.Fatal("Failed to check purchase")
		}
		if !result {
			t.Error(fmt.Sprintf("purchase of id %d is not available", purchase.GetID()))
		}
	}

}
