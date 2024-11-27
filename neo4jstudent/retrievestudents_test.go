package neo4jstudent

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestRetrieveSchoolStudents(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)

	start := time.Now()
    
	students, error := RetrieveSchoolStudents(&neo4jInstance, 1)

	if error != nil{
		t.Error(`Failed to retrieve student nodes`)
	}
    
	for _, student := range students{
    
	switch student.GetID() {
	case 4:
		log.Println("successfully retrieved node: ", student.GetName())
	case 3:
		log.Println("successfully retrieved node: ", student.GetName())	
	case 1:
		log.Println("successfully retrieved node: ", student.GetName())	
	case 6:
		log.Println("successfully retrieved node: ", student.GetName())	
	case 5:
		log.Println("successfully retrieved node: ", student.GetName())	
	default:
		t.Error(`retrieved non-existent student node:`, student.GetName())

	}

    elapsed := time.Since(start)
    fmt.Printf("Function took %s\n", elapsed)
}
}