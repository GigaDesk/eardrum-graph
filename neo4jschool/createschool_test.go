package neo4jschool

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/GigaDesk/eardrum-graph/mockschool"
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/joho/godotenv"
)

var (
	neo4jInstance  neo4jutils.Neo4jInstance
)
func TestAddSchoolNode(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)

	start := time.Now()


	if err := CreateSchool(&neo4jInstance, mockschool.SchoolNode); err != nil{
		t.Error(`Failed to add a school node`)
	}
    elapsed := time.Since(start)
    fmt.Printf("Function took %s\n", elapsed)
}

func TestAddSamePhoneNumberSchoolNode(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)

	start := time.Now()


	if err := CreateSchool(&neo4jInstance, mockschool.SamePhoneNumberSchoolNode); err==nil{
		t.Error(`Added a school node with the same phonenumber`)
	}
    elapsed := time.Since(start)
    fmt.Printf("Function took %s\n", elapsed)
}
func TestAddSameIdSchoolNode(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)

	start := time.Now()


	if err := CreateSchool(&neo4jInstance, mockschool.SameIdSchoolNode); err==nil{
		t.Error(`Added a school node with the same Id`)
	}
    elapsed := time.Since(start)
    fmt.Printf("Function took %s\n", elapsed)
}
func TestAddMultipleSchoolNodes(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)

	start := time.Now()

    for _ , school := range mockschool.MultipleSchoolNodes{
	if err := CreateSchool(&neo4jInstance, school); err!=nil{
		t.Error(`Failed to add multiple school nodes`)
	}
}
    elapsed := time.Since(start)
    fmt.Printf("Function took %s\n", elapsed)
}
