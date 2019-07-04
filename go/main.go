package main

import "time"

type EventLogEntry struct {
	CustomerId string
	TxnId      string
	EventTime  time.Time
	Amount     float64
}

type ResultLogEntry struct {
	CustomerId string
	TxnId      string
	Accepted   bool
}

type App struct {
	Manager DBManager
}

func NewApp() App {
	manager := GetDBManager()
	return App{Manager: manager}
}

func main() {

	app := NewApp()
	app.process()

}

func (app *App) process() {
	// TODO get the file
}

// ReadPackedFile is a function to unpack a tar.gz
func ReadJsonFile(filepath string) {

}
