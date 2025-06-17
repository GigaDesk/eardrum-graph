package neo4jstudent

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/GigaDesk/eardrum-graph/mockuser"
	"github.com/joho/godotenv"
)

func TestUpdateUserNode(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)
	//create user nodes of primary keys 3,4,5 and 6
	for _, user := range mockuser.MultipleUserNodes {
		if err := CreateUser(&neo4jInstance, user); err != nil {
			t.Error(`Failed to add a user node`)
		}
		result, error, _ := CheckUser(&neo4jInstance, int(user.GetID()))
		if error != nil {
			log.Fatal("Failed to check user")
		}
		if !result {
			t.Error(fmt.Sprintf("user of id %d is not available", user.GetID()))
		}
	}

	// update
	if err:= UpdateUser(&neo4jInstance, mockuser.UpdatedUser); err !=nil{
		t.Error(err)
	}

	//check changes
	_, err, user := CheckUser(&neo4jInstance, int(mockuser.UpdatedUser.GetID()))

	if err != nil {
		t.Error(err)
	}


	//throw errors for unupdated name
	if user.GetName() != mockuser.UpdatedUser.GetName() {
		t.Error("name is ", user.GetName(), "instead of ", mockuser.UpdatedUser.GetName())
	}

	//throw errors for unupdated phone number
	if user.GetPhoneNumber() != mockuser.UpdatedUser.GetPhoneNumber() {
		t.Error("phone number is ", user.GetPhoneNumber(), "instead of ", mockuser.UpdatedUser.GetPhoneNumber())
	}

	//throw errors for unupdated account balance
	if user.GetAccountBalanceInCents() != mockuser.UpdatedUser.GetAccountBalanceInCents() {
		t.Error("account balance is ", user.GetAccountBalanceInCents(), "instead of ", mockuser.UpdatedUser.GetAccountBalanceInCents())
	}
}
