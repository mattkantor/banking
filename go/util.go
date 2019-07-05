package main

import (
	"time"
)

// I read an artcile about not writing util.go type files.

func CleanCurrency(currency string) float64{
	return 0
}

func GetMondayAndoffsetForDate(date time.Time) (string, int){
	a := date.Weekday()
	date.AddDate(0,0,-int(a))
	dt:=date.AddDate(0,0,-int(a))

	t :=  dt.Format( "2006-01-02")
	return t, int(a)
}

func ParseFileDateIntoRealDate(date string) time.Time{
	// 2000-01-01T00:00:00Z

	timeOut, err := time.Parse( "2006-01-02T15:04:05Z", date)
	if err != nil{
		panic(err)
	}
	return timeOut
}

func contains(item string, existing_items []string) bool {
	for _, n := range existing_items {
		if item == n {
			return true
		}
	}
	return false
}

func sum(input []float64) float64 {
	sum := 0.00

	for i := range input {
		sum = sum +  input[i]
	}


	return sum
}
