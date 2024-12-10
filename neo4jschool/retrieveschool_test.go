package neo4jschool

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/GigaDesk/eardrum-graph/mockschool"
	"github.com/GigaDesk/eardrum-graph/mockstudent"
	"github.com/GigaDesk/eardrum-graph/neo4jstudent"
	"github.com/joho/godotenv"
)

func TestRetrieveSchool(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)

	// add school nodes of primary keys: 3,4,5,6
	for _, school := range mockschool.MultipleSchoolNodes {
		CreateSchool(&neo4jInstance, school)
	}
	// attach student of primary key 4 to school of primary key 6
	neo4jstudent.CreateStudent(&neo4jInstance, mockstudent.MultipleStudentNodes[1], 6)
	// attach student of primary key 3 to school of primary key 4
	neo4jstudent.CreateStudent(&neo4jInstance, mockstudent.MultipleStudentNodes[0], 4)
    //retrieve the school of student of primary key 3
	school1, err1 := RetrieveStudentSchool(&neo4jInstance, 3)
	if err1 != nil{
		log.Fatal(err1)
	}
	//throw an error if the primary key of the school is not 4
	if school1.GetID() != 4{
		t.Error(fmt.Sprintf("student of primary key %d belongs to school of primary key %d", mockstudent.MultipleStudentNodes[0].Id, school1.GetID()))
	}
	//retrieve the school of student of primary key 4
	school2, err2 := RetrieveStudentSchool(&neo4jInstance, 4)
	if err2 != nil{
		log.Fatal(err2)
	}
	//throw an error if the primary key of the school is not 6
	if school2.GetID() != 6{
		t.Error(fmt.Sprintf("student of primary key %d belongs to school of primary key %d", mockstudent.MultipleStudentNodes[1].Id, school2.GetID()))
	}
}
