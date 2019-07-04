package main

import "time"

type CustomerController interface {
	AddDeposit(customer string, txn string, date time.Time, amount float64)
	CheckForExistingTransaction()
	GetDepositsForDate()
	GetDepositsForWeek()
	GetCustomerHistory()
}

var instance DBManager

// DBManager holds the database config
type DBManager struct {
	// TODO make this a proper data storage structure
	db            map[string]string
	isInitialized bool
}

// GetDBManager is the constructor for singleton
func GetDBManager() DBManager {
	if instance.isInitialized {
		return instance
	}
	return newDbManager()

}

func newDbManager() DBManager {

	return DBManager{isInitialized: true, db: make(map[string]string)}
}

func (dbManager DBManager) AddDeposit(customer string, txn string, date time.Time, amount float64) (bool, error) {
	// TODO
	return true, nil
}
