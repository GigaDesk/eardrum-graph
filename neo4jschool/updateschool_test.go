package neo4jschool

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/GigaDesk/eardrum-graph/mockschool"
	"github.com/joho/godotenv"
)

func TestUpdateSchool(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)
	// Create and Check each school in the slice
	for _, school := range mockschool.MultipleSchoolNodes {
		if err := CreateSchool(&neo4jInstance, school); err != nil {
			t.Error(`Failed to add multiple school nodes`)
		}
		result, error, _ := CheckSchool(&neo4jInstance, int(school.GetID()))
		if error != nil {
			log.Fatal("Failed to check school")
		}
		if !result {
			t.Error(fmt.Sprintf("school of id %d is not available", school.GetID()))
		}
	}

	// update
	if err:=UpdateSchool(&neo4jInstance, mockschool.UpdatedSchool); err!=nil{
		t.Error(err)
	}

	//check changes
	_, err, school := CheckSchool(&neo4jInstance, int(mockschool.UpdatedSchool.GetID()))

	if err!=nil{
		t.Error(err)
	}

	//throw errors for unupdated name
	if school.GetName()!= mockschool.UpdatedSchool.GetName(){
		t.Error("name is ", school.GetName(), "instead of ", mockschool.UpdatedSchool.GetName())
	}


	//throw errors for unupdated websites
	if school.GetWebsite()!= mockschool.UpdatedSchool.GetWebsite(){
		t.Error("website is ", school.GetWebsite(), "instead of ", mockschool.UpdatedSchool.GetWebsite())
	}

}
