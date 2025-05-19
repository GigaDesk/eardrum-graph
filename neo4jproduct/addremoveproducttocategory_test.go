package neo4jproduct

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/GigaDesk/eardrum-graph/mockproduct"
	"github.com/GigaDesk/eardrum-graph/mockschool"
	"github.com/GigaDesk/eardrum-graph/mockshop"
	"github.com/GigaDesk/eardrum-graph/neo4jschool"
	"github.com/GigaDesk/eardrum-graph/neo4jshop"
	"github.com/joho/godotenv"
)

func TestAddProductToCategory(t *testing.T) {

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

	//create category nodes of primary keys 3,4,5 and 6 to shop node of primary key 5
	for _, category := range mockproduct.MultipleCategoryNodes {
		if err := CreateCategory(&neo4jInstance, category, 5); err != nil {
			t.Error(`Failed to add a category node`)
		}
		result, error := CheckCategory(&neo4jInstance, int(category.GetID()))
		if error != nil {
			log.Fatal("Failed to check category")
		}
		if !result {
			t.Error(fmt.Sprintf("category of id %d is not available", category.GetID()))
		}
	}

	AddProductToCategory(&neo4jInstance, 3, 3)
	AddProductToCategory(&neo4jInstance, 4, 4)
	AddProductToCategory(&neo4jInstance, 5, 4)
	AddProductToCategory(&neo4jInstance, 6, 6)

	products, error := RetrieveCategoryProducts(&neo4jInstance, 4)
    if error != nil {
		t.Error(`error retrieving category products`)
	}

	if len(products) != 2{
		t.Error(fmt.Sprintf("length of retrieved products array is not: %d", 2))
	}

	for _,s:=range products{
	switch s.GetID(){
	case 4:
		log.Println(fmt.Sprintf("found product with id %d:", 4))
	case 5:
		log.Println(fmt.Sprintf("found product with id %d:", 5))
	default:
		t.Error(fmt.Sprintf("found product with id %d:", s.GetID()))
	}
	}

	products, error = RetrieveCategoryProducts(&neo4jInstance, 3)
    if error != nil {
		t.Error(`error retrieving category products`)
	}

	if len(products) != 1{
		t.Error(fmt.Sprintf("length of retrieved products array is not: %d", 1))
	}

	for _,s:=range products{
	switch s.GetID(){
	case 3:
		log.Println(fmt.Sprintf("found product with id %d:", 3))
	default:
		t.Error(fmt.Sprintf("found product with id %d:", s.GetID()))
	}
	}

	products, error = RetrieveCategoryProducts(&neo4jInstance, 6)
    if error != nil {
		t.Error(`error retrieving category products`)
	}

	if len(products) != 1{
		t.Error(fmt.Sprintf("length of retrieved products array is not: %d", 1))
	}

	for _,s:=range products{
	switch s.GetID(){
	case 6:
		log.Println(fmt.Sprintf("found product with id %d:", 6))
	default:
		t.Error(fmt.Sprintf("found product with id %d:", s.GetID()))
	}
	}

	RemoveProductFromCategory(&neo4jInstance, 4, 4)

	products, error = RetrieveCategoryProducts(&neo4jInstance, 4)
    if error != nil {
		t.Error(`error retrieving category products`)
	}

	if len(products) != 1{
		t.Error(fmt.Sprintf("length of retrieved products array is not: %d", 1))
	}

	for _,s:=range products{
	switch s.GetID(){
	case 4:
		log.Println(fmt.Sprintf("found product with id %d:", 4))
	case 5:
		log.Println(fmt.Sprintf("found product with id %d:", 5))
	default:
		t.Error(fmt.Sprintf("found product with id %d:", s.GetID()))
	}
	}

}
