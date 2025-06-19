package mocktransaction

import (
	"time"
)

var MultipleTransactionNodes = []MockTransaction{
	{
		Id:                      3,
		CreatedAt:               time.Now().UTC(),
		UpdatedAt:               time.Now().UTC(),
		TotalAmountInCents:      105,
	},
	{
		Id:                      4,
		CreatedAt:               time.Now().UTC(),
		UpdatedAt:               time.Now().UTC(),
		TotalAmountInCents:      205,
	},
	{
		Id:                      5,
		CreatedAt:               time.Now().UTC(),
		UpdatedAt:               time.Now().UTC(),
		TotalAmountInCents:      305,
	},
	{
		Id:                      6,
		CreatedAt:               time.Now().UTC(),
		UpdatedAt:               time.Now().UTC(),
		TotalAmountInCents:      405,
	},
}
