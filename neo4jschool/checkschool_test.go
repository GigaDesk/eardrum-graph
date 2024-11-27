package neo4jschool

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestCheckSchoolNode(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)

	start := time.Now()

    result, err := CheckSchool(&neo4jInstance, 1)
	if err != nil{
		t.Error(`Failed to check school node's existence`)
	}
	if !result{
		t.Error(`Failed to find an existing school node`)
	}
	result1, err1 := CheckSchool(&neo4jInstance, 2)
	if err1 != nil{
		t.Error(`Failed to check school node's existence`)
	}
	if result1{
		t.Error(`Found a non-existing school node`)
	}

    elapsed := time.Since(start)
    log.Printf("Function took %s\n", elapsed)
}
