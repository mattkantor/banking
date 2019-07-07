package main

import (
	"bufio"
	"encoding/json"
	"flag"
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
	Id      string `json:"id"`
	CustomerId string `json:"customer_id"`
	Accepted   bool `json:"accepted"`
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
	os.Create(outputFile) //let it fail
	fOut, err := os.OpenFile(outputFile,1,777)
	if err != nil {
		panic(err)
	}

	defer fOut.Close()


	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		var entry EventLogEntry
		line := scanner.Text()
		err := json.Unmarshal([]byte(line),&entry)
		if err != nil{
			panic(err)
		}
		accepted, code := app.Cc.LoadCard(entry)

		if code != 403 {
			writeLog(fOut, entry, accepted)

		}
	}
}
func writeLog(fOut *os.File,e EventLogEntry, accepted bool){
	var out = ResultLogEntry{CustomerId:e.CustomerId, Id:e.Id, Accepted: accepted}

	outJson, _ := json.Marshal(out)

	if _, err := fOut.Write(outJson); err != nil {
		panic(err)
	}
	// TODO Sorry hacky
	if _, err := fOut.Write([]byte("\n")); err != nil {
		panic(err)
	}



}
