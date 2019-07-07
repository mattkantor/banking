package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

type EventLogEntry struct {
	CustomerId string `json:"customer_id"`
	Id      string  `json:"id"`
	EventTime  string  `json:"time"`
	Amount     string `json:"load_amount"`
}

type ResultLogEntry struct {
	CustomerId string
	Id      string
	Accepted   bool
}

type App struct {
	Cc CustomerController

}

func NewApp() App {
	cc := NewCustomerController()
	return App{Cc: cc}
}

func main() {
	//process flags
	inFile := *flag.String("in", "../input.txt","Name of the file")
	outFile := *flag.String("out", "../results_go.txt", "File to export")
	flag.Parse()
	app := NewApp()

	app.process(inFile, outFile)

}
func (app *App) process(inputFile, outputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//var entries []EventLogEntry
	for scanner.Scan() {
		var entry EventLogEntry
		line := scanner.Text()
		err := json.Unmarshal([]byte(line),&entry)
		if err != nil{
			panic(err)
		}
		accepted, code := app.Cc.AddDeposit(entry)
		if code != 403 {
			writeLog(entry, accepted)

		}
	}
}
func writeLog(e EventLogEntry, accepted bool){
	var out = ResultLogEntry{CustomerId:e.CustomerId, Id:e.Id, Accepted: accepted}
	fmt.Println(out)

}
