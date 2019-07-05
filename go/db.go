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
	Deposits []DepositsPerWeek
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

func (dbManager *DBManager) addDeposit(e EventLogEntry) bool {
	// TODO

	customerTransactions := dbManager.db[e.CustomerId].Transactions
	customerTransactions = append(customerTransactions, e.TxnId)

	customerDeposits := dbManager.db[e.CustomerId].Deposits
	//monday := customerDeposits[
	make(customerDeposits)





	return true
}

func (dbManager *DBManager) getCustomerData(customerId string) CustomerData{
	return dbManager.db[customerId]

}
