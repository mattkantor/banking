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

func parseInputFile(fileName string) []EventLogEntry {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	var entries []EventLogEntry
	for scanner.Scan() {
		var entry EventLogEntry
		line := scanner.Text()
		err := json.Unmarshal([]byte(line),&entry)
		if err != nil{
			panic(err)
		}
		entries = append(entries, entry)
	}
	return entries


}
func writeLog(e EventLogEntry, accepted bool){
	var out = ResultLogEntry{CustomerId:e.CustomerId, Id:e.Id, Accepted: accepted}
	fmt.Println(out)

}

func (app *App) process(inputFile, outputFile string) {

	var inputs []EventLogEntry
	inputs = parseInputFile(inputFile)

	for i := 0; i < len(inputs); i++ {
		fmt.Println(inputs[i])
		accepted, err := app.Cc.AddDeposit(inputs[i])
		if err != 403{
			writeLog(inputs[i], accepted)
		}

	}


}
