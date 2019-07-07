package main

import (
	"log"
	"regexp"
	"strconv"
	"time"
)

// I read an artcile about not writing util.go type files.

func CleanCurrency(currency string) float64{

	reg, err := regexp.Compile("[^0-9\\.]")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(currency, "")

	curr, _:= strconv.ParseFloat(  processedString, 64)
	return curr

}

func GetMondayAndoffsetForDate(date time.Time) (string, int){
	a := date.Weekday()

	dt:=date.AddDate(0,0,-int(a)+1)

	t :=  dt.Format( "2006-01-02")
	return t, int(a)
}

func ParseFileDateIntoRealDate(date string) time.Time{
	// 2000-01-01T00:00:00Z
	layout := "2006-01-02T15:04:05Z"
	timeOut, err := time.Parse( layout, date)
	if err != nil{
		panic(err)
	}
	return timeOut
}

func contains(item string, existingItems []string) bool {
	for _, n := range existingItems {
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
