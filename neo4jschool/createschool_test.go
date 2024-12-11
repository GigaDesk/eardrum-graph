package neo4jschool

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/GigaDesk/eardrum-graph/mockschool"
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/joho/godotenv"
)

var (
	neo4jInstance  neo4jutils.Neo4jInstance
)

func TestCreateSchool(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)
// Create and Check each school in the slice
    for _ , school := range mockschool.MultipleSchoolNodes{
	if err := CreateSchool(&neo4jInstance, school); err!=nil{
		t.Error(`Failed to add multiple school nodes`)
	}
	result, error := CheckSchool(&neo4jInstance, int(school.GetID())) 
	if error != nil {
		log.Fatal("Failed to check school")
	}
	if !result{
		t.Error(fmt.Sprintf("school of id %d is not available", school.GetID()))
	}
}
  
}
