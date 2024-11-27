package neo4jstudent

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/GigaDesk/eardrum-graph/mockstudent"
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/joho/godotenv"
)

var (
	neo4jInstance  neo4jutils.Neo4jInstance
)
func TestAddStudentNode(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)

	start := time.Now()


	if err := CreateStudent(&neo4jInstance, mockstudent.StudentNode, 1); err != nil{
		t.Error(`Failed to add a student node`)
	}
    elapsed := time.Since(start)
    fmt.Printf("Function took %s\n", elapsed)
}

func TestAddSameRegistrationNumberStudentNode(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)

	start := time.Now()


	if err := CreateStudent(&neo4jInstance, mockstudent.SameRegistrationNumberStudentNode, 1); err==nil{
		t.Error(`Added a student node with the same registration number`)
	}
    elapsed := time.Since(start)
    fmt.Printf("Function took %s\n", elapsed)
}
func TestAddSameIdStudentNode(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)

	start := time.Now()


	if err := CreateStudent(&neo4jInstance, mockstudent.SameIdStudentNode, 1); err==nil{
		t.Error(`Added a student node with the same Id`)
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

    for _ , student := range mockstudent.MultipleStudentNodes{
	if err := CreateStudent(&neo4jInstance, student, 1); err!=nil{
		t.Error(`Failed to add multiple student nodes`)
	}
}
    elapsed := time.Since(start)
    fmt.Printf("Function took %s\n", elapsed)
}