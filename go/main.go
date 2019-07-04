package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"

	"bufio"
)

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
		json.Unmarshal([]byte(line),&entry)
		entries = append(entries, entry)
	}
	return entries


}
func writeLog(out ResultLogEntry){

}

func (app *App) process(inputFile, outputFile string) {

	var inputs []EventLogEntry
	inputs = parseInputFile(inputFile)
	for i := 0; i < len(inputs); i++ {
		out, err:=app.Cc.AddDeposit(inputs[i])
		if err != 403{
			writeLog(out)
		}

	}


}
