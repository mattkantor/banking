package main

import "time"

type EventLogEntry struct {
	UserId    string
	EventId   string
	Amount    float64
	EventTime time.Time
}

type UserLogEntry struct {
	EventTime time.Time
	Amount    float64
}

type UserLog struct {
	UserId   string
	LogEntry []UserLogEntry
}

func main() {

}

// ReadPackedFile is a function to unpack a tar.gz
func ReadJsonFile(filepath string) {

}

func processFile(srcFile string) {

}
