package neo4jstudent

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/GigaDesk/eardrum-graph/mockschool"
	"github.com/GigaDesk/eardrum-graph/mockstudent"
	"github.com/GigaDesk/eardrum-graph/neo4jschool"
	"github.com/joho/godotenv"
)

func TestRetrieveSchoolStudents(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)
	//create a school node of primary key 1
	neo4jschool.CreateSchool(&neo4jInstance, mockschool.SchoolNode)
	//create a student nodes of primary keys 3,4,5 and 6 to school node of primary key 1
	for _,student := range mockstudent.MultipleStudentNodes {
		if err := CreateStudent(&neo4jInstance, student, 1); err != nil {
			log.Fatal("Failed to add a student node")
		}
	}
	students, error := RetrieveSchoolStudents(&neo4jInstance,1)
    if error != nil {
		t.Error(`error retrieving school students`)
	}

	if len(students) != 4{
		t.Error(fmt.Sprintf("length of retrieved students array is not: %d", 4))
	}

	for _,s:=range students{
	switch s.GetID(){
	case 3:
		log.Println(fmt.Sprintf("found student with id %d:", 3))
	case 4:
		log.Println(fmt.Sprintf("found student with id %d:", 4))
	case 5:
		log.Println(fmt.Sprintf("found student with id %d:", 5))
	case 6:
		log.Println(fmt.Sprintf("found student with id %d:", 6))
	default:
		t.Error(fmt.Sprintf("found student with id %d:", s.GetID()))
	}
	}
}