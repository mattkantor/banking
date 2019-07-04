package main



type DepositsPerDay struct{
	DayOfWeek int8
	Deposits []float64
}

type DepositsPerWeek struct {
	Monday string
	DailyDeposits []DepositsPerDay
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

func (dbManager DBManager) AddDeposit(e EventLogEntry) (ResultLogEntry, int) {
	// TODO
	var res  = ResultLogEntry{CustomerId:"1", TxnId:"1", Accepted:true}
	return res, 200
}
