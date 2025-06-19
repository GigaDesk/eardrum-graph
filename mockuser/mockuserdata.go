package mockuser

import "time"

//Mock data for carrying out user tests

var UserNode = MockUser{
	Id:          1,
	CreatedAt:   time.Now().UTC(),
	UpdatedAt:   time.Now().UTC(),
	Name:        "Leon Kenyaga",
	PhoneNumber: "+254719226150",
	Password:    "kisbhdbcvukbqiyde327&",
}

var MultipleUserNodes = []MockUser{
	{
	Id:          3,
	CreatedAt:   time.Now().UTC(),
	UpdatedAt:   time.Now().UTC(),
	Name:        "Arnold Osoro",
	PhoneNumber: "+254719226155",
	Password:    "kisbhdbcvukbqiyde327&",
	},
	{
		Id:          4,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Name:        "Julius Nyakweba",
		PhoneNumber: "+254719226156",
		Password:    "kisbhdbcvukbqiyde327&",
	},
	{
		Id:          5,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Name:        "Carlos Okumu",
		PhoneNumber: "+254719226157",
		Password:    "kisbhdbcvukbqiyde327&",
		Account_balance_in_cents: 500,
	},
	{
		Id:          6,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Name:        "Greg Otieno",
		PhoneNumber: "+254719226158",
		Password:    "kisbhdbcvukbqiyde327&",
	},
}

var UpdatedUser = MockUser{
		Id:          5,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Name:        "Carlos Okumu Nyambati",
		PhoneNumber: "+254705136690",
		Password:    "kisbhdbcvukbqiyde327&",
		Account_balance_in_cents: 5000,
}