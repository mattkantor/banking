package main


type DepositsPerWeek struct {
		MondayDate map[string]Daily
}

type Daily struct{
	Day    map[int][]float64
}

type CustomerData struct{
	CustomerId string //purposefully redundant
	Transactions []string
	Deposits DepositsPerWeek
}

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
	// TODO
	monday, dayIndex :=  GetMondayAndoffsetForDate(ParseFileDateIntoRealDate(e.EventTime))

	customerData := dbManager.db[e.CustomerId]

	customerTransactions := customerData.Transactions

	customerTransactions = append(customerTransactions, e.Id)

	customerDeposits := dbManager.db[e.CustomerId].Deposits

	// TODO make
	var depositsPerDay  = customerDeposits.MondayDate[monday].Day[dayIndex]

	depositsPerDay = append(depositsPerDay, e.Amount)

	dbManager.db[e.CustomerId] = CustomerData{Transactions:customerTransactions, Deposits:customerDeposits, CustomerId:e.CustomerId}

	return true
}

func (dbManager *DBManager) getCustomerData(customerId string) CustomerData{
	return dbManager.db[customerId]

}
