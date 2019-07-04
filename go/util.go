package main

import "time"

// I read an artcile about not writing util.go type files.

func CleanCurrency(currency string) float64{
	return 0
}

func GetMondayAndoffsetForDate(date string) (string, string){
		return "0000-00-00", "0"
}

func ParseFileDateIntoRealDate(date string) (time.Time){
	return time.Now()
}