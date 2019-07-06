package main

import (
	"fmt"

)

type CustomerData struct{
	CustomerId string //purposefully redundant
	Transactions []string
	Deposits DepositsPerWeek
}
type DepositsPerWeek struct {
		Mondays map[string]Dailies
}

type Dailies struct{
	Day  map[int]Deposits
}
type Deposits struct {
	Amount []float64
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

	fmt.Println(customerDeposits)
	// TODO make this less messy


	/*
	check for a customer record

	yes -
	- check for a monday record
	No
	- create empty record

	yes -
	check foor a daily record

	yes -
	append to day


	if daylist, ok := mondayList.Day[dayIndex];ok{
			daylist.Amounts = append(daylist.Amounts, amountFloat)

			mondayMap[monday].Day[dayIndex] = daylist
		}else{

			mondayMap[monday].Day[dayIndex] = Deposits{Amounts: []float64{amountFloat}}
		}

	 */


	mondayMap := customerDeposits.Mondays
	fmt.Println(mondayMap)
	if len(mondayMap) == 0 {
		mondayMap = make(map[string]Dailies, 1)
		mondayMap[monday]=Dailies{}
	}
	if _, ok := mondayMap[monday]; !ok {
		
		daily[dayIndex]=Deposits{}
		mondayMap[monday] = daily
	}
	fmt.Println(mondayMap)
	if _, ok := mondayMap[monday].Day[dayIndex];!ok{
		fmt.Println("Initializing day")
		tempThing := mondayMap[monday].Day
		fmt.Println(tempThing)
		mondayMap[monday].Day[dayIndex] = Deposits{}
		amountTemp := make([]float64,0)
		mondayMap[monday].Day[dayIndex]=Deposits{Amount: amountTemp}
		// map exists - week found

	}
	amounts := mondayMap[monday].Day[dayIndex].Amount
	newTotal := append(amounts, amountFloat)
	tempDay :=  mondayMap[monday].Day[dayIndex]
	tempDay.Amount = newTotal
	mondayMap[monday].Day[dayIndex] = tempDay

	customerDeposits.Mondays = mondayMap

	dbManager.db[e.CustomerId] = CustomerData{Transactions:customerTransactions, Deposits:customerDeposits, CustomerId:e.CustomerId}

	return true
}

func (dbManager *DBManager) getCustomerData(customerId string) CustomerData{
	return dbManager.db[customerId]

}
