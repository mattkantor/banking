package main



const maxItemsPerDay = 3
const maxAmountPerDay = 5000
const maxAmountPerWeek = 20000

func TransactionExists(customerRecords CustomerData,txnId string ) bool {
	transactions := customerRecords.Transactions
	for i := 0; i < len(transactions); i++ {

		if transactions[i] == txnId {
			return true
		}
	}
	return false
}

func IsValid(customerRecords CustomerData, date string, amount float64) bool{
	if amount > maxAmountPerDay {
		return false
	}




	return false
}