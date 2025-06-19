package mockpurchase


type MockPurchase struct{
	Id int
	TransactionId int
	ProductId int
	UnitsBought int
	TotalAmountInCents int
}


func (m MockPurchase) GetID() int64 {
	return int64(m.Id)
}

func (m MockPurchase) GetTransactionID() int64 {
	return int64(m.TransactionId)
}

func (m MockPurchase) GetProductID() int64 {
	return int64(m.ProductId)
}

func (m MockPurchase) GetUnitsBought() int {
	return m.UnitsBought
}

func (m MockPurchase) GetTotalAmountInCents() int64 {
	return int64(m.TotalAmountInCents)
}