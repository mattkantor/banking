package main

import "testing"

func TestCreateANewDB(t *testing.T) {

	DbManager := GetDBManager()

	e := EventLogEntry{
		CustomerId:"123",
		Id: "123",
		EventTime:"2019-01-03T12:13:34Z",
		Amount:2345.67,
	}

	success :=DbManager.loadAccount(e)
	if !success {
		t.Errorf("Could not create a new DB instance")
	}
}
