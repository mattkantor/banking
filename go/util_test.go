package main

import (
	"testing"
)


func TestParseFileDateIntoRealDate(t *testing.T) {

	// 2000-01-01T00:00:00Z
	 timeOut := ParseFileDateIntoRealDate(" 2000-01-01T00:00:01Z")
	 if timeOut.Year() != 2000 {
		 t.Errorf("Date conversion year is wrong, %d should be %s", timeOut.Year() , "2000")
	 }
	if timeOut.Month() != 1 {
		t.Error("Date conversion month is wrong")
	}
	if timeOut.Second() != 2 {
		t.Error("Date conversion second is wrong")
	}
}



func TestGetMondayAndoffsetForDate(t *testing.T){
	timeOut := ParseFileDateIntoRealDate(" 2019-07-05T10:00:01Z")
	monday, dayindex := GetMondayAndoffsetForDate(timeOut)
	if monday != "2019-07-01"  && dayindex != 5 {
		t.Errorf("Monday not computed correctly from ,%s,  %s, %d",timeOut, monday, dayindex)
	}

}

func TestSumFloats(t *testing.T){
	var a []float64  = make([]float64,3)
	a[0] = 111.11
	a[1] = 222.22
	a[2] = 333.33

	if sum(a) != 666.66 {
		t.Error("Sumthing is broken!")
	}


}
