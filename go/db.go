package main

import "fmt"

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

// assume customer already exists because this happens after the write.
// We should not assume this in general
func (dbManager *DBManager) recordTransaction(customerId, txn string){
	customerData := dbManager.getCustomerData(customerId)

	customerTransactions := customerData.Transactions
	customerTransactions = append(customerTransactions, txn)
	customerData.Transactions = customerTransactions
	dbManager.db[customerId] = customerData
}

func (dbManager *DBManager)recordLoad(e EventLogEntry) bool{

	customerData := dbManager.getCustomerData(e.CustomerId)

	monday, dayIndex :=  GetMondayAndoffsetForDate(ParseFileDateIntoRealDate(e.EventTime))
	amountFloat:=CleanCurrency(e.Amount)

	//assume it's true for the sake of this exercise

	customerDeposits := customerData.Deposits

	if  customerDeposits[monday]==nil {


		customerDeposits[monday]=map[int][]float64{}
	}
	if _, ok := customerDeposits[monday][dayIndex];!ok{
		customerDeposits[monday][dayIndex]= []float64{}
	}

	amounts := customerDeposits[monday][dayIndex]
	newTotal := append(amounts, amountFloat)

	customerDeposits[monday][dayIndex] = newTotal

	customerData.Deposits = customerDeposits

	dbManager.db[e.CustomerId] = customerData
	return true // false would be for a failure which doesn't panic. No case for this here
}




func (dbManager *DBManager) getCustomerData(customerId string) CustomerData{
	var ok bool

	if _, ok = dbManager.db[customerId];!ok{
		// new customer - yay!
		fmt.Println("Im new!")
		customerData := CustomerData{CustomerId:customerId, Transactions:[]string{}, Deposits: map[string]map[int][]float64{}}
		dbManager.db[customerId] = customerData

	}
	return dbManager.db[customerId]

}
