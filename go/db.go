package main

// a simplesitic approach
type CustomerData struct{
	CustomerId string //purposefully redundant
	Transactions []string
	Deposits map[string]map[int][]float64
}
//type DepositsPerWeek struct {
//		Mondays map[string]Dailies
//}

//type Dailies struct{
//	Day  map[int]Deposits
//}
//type Deposits struct {
//	Amount []float64
//}


var instance DBManager

// DBManager holds the database config
type DBManager struct {
	// TODO make this a proper data storage structure
	db            map[string]CustomerData
	isInitialized bool
}

// GetDBManager is the constructor for singleton
func GetDBManager() *DBManager {
	if instance.isInitialized {
		return &instance
	}
	return newDbManager()

}

func newDbManager() *DBManager {

	return &DBManager{isInitialized: true, db: make(map[string]CustomerData)}
}



func (dbManager *DBManager) loadAccount(e EventLogEntry) bool {
	// TODO init customer
	monday, dayIndex :=  GetMondayAndoffsetForDate(ParseFileDateIntoRealDate(e.EventTime))
	amountFloat:=CleanCurrency(e.Amount)
	var customerData CustomerData
	var ok bool
	if customerData, ok = dbManager.db[e.CustomerId];!ok{
		// new customer
		customerData = CustomerData{CustomerId:e.CustomerId}

	}

	customerTransactions := customerData.Transactions
	customerTransactions = append(customerTransactions, e.Id)

	customerDeposits := dbManager.db[e.CustomerId].Deposits




	if len(customerDeposits) == 0 {
		customerDeposits = map[string]map[int][]float64{}
		//customerDeposits[monday]=make(map[int],0)
	}
	if _, ok := customerDeposits[monday]; !ok {

		customerDeposits[monday]=map[int][]float64{}
		//customerDeposits[monday][dayIndex]= make([]float64,0)
	}
	if _, ok := customerDeposits[monday][dayIndex];!ok{
		customerDeposits[monday][dayIndex]= []float64{}
	}

	amounts := customerDeposits[monday][dayIndex]

	newTotal := append(amounts, amountFloat)

	customerDeposits[monday][dayIndex] = newTotal
	dbManager.db[e.CustomerId] = CustomerData{Transactions:customerTransactions, Deposits:customerDeposits, CustomerId:e.CustomerId}


	return true
}

func (dbManager *DBManager) getCustomerData(customerId string) CustomerData{
	return dbManager.db[customerId]

}
