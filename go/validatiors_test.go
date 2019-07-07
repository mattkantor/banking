package main

import "testing"

const customerId = "123"


func TestShouldValidateDailyCOunt(t *testing.T){
	var maxItemsPerDayValidator MaxItemsPerDayValidator
	e:=EventLogEntry{CustomerId:customerId, Id:"123",Amount:"$2933.34", EventTime:"2019-07-01T12:34:56Z"}
	deposits:=map[string]map[int][]float64{}

	deposits["2019-07-01"]=map[int][]float64{}
	deposits["2019-07-01"][1] = []float64{1234.56,2345.45}
	customerData := CustomerData{Transactions:[]string{"123","456"},
		CustomerId:customerId,
		Deposits: deposits}
	v := maxItemsPerDayValidator.validate(customerData, e)
	if v != true{
		t.Error("validator failed")
	}
}
func TestShouldFailDailyCount(t *testing.T){
	var maxItemsPerDayValidator MaxItemsPerDayValidator
	e:=EventLogEntry{CustomerId:customerId, Id:"123",Amount:"$2933.34", EventTime:"2019-07-01T12:34:56Z"}
	deposits:=map[string]map[int][]float64{}
	deposits["2019-07-01"]=map[int][]float64{}
	deposits["2019-07-01"][1] = []float64{1234.56,2345.45,3444}
	customerData := CustomerData{Transactions:[]string{"123","456","765"},
		CustomerId:customerId,
		Deposits: deposits}
	v := maxItemsPerDayValidator.validate(customerData, e)
	if v == true{
		t.Error("validator failed - should not allow another txn")
	}
}
func TestShouldValidateDailyAmount(t *testing.T){
	var maxAmountPerDayValidator MaxAmountPerDayValidator
	e:=EventLogEntry{CustomerId:customerId, Id:"123",Amount:"$1933.34", EventTime:"2019-07-01T12:34:56Z"}
	deposits:=map[string]map[int][]float64{}

	deposits["2019-07-01"]=map[int][]float64{}
	deposits["2019-07-01"][1] = []float64{1234.56,1345.45}
	customerData := CustomerData{Transactions:[]string{"123","456"},
		CustomerId:customerId,
		Deposits: deposits}
	v := maxAmountPerDayValidator.validate(customerData, e)
	if v != true{
		t.Error("validator failed - should have added more momney per day")
	}
}
func TestShouldFailDailyAmount(t *testing.T){
	var maxAmountPerDayValidator MaxAmountPerDayValidator
	e:=EventLogEntry{CustomerId:customerId, Id:"123",Amount:"$2933.34", EventTime:"2019-07-01T12:34:56Z"}
	deposits:=map[string]map[int][]float64{}
	deposits["2019-07-01"]=map[int][]float64{}
	deposits["2019-07-01"][1] = []float64{2234.56,2345.45}
	customerData := CustomerData{Transactions:[]string{"123","456"},
		CustomerId:customerId,
		Deposits: deposits}
	v := maxAmountPerDayValidator.validate(customerData, e)
	if v == true{
		t.Error("validator failed - should not allow another txn too much $ per day")
	}
}
//weekly

func TestShouldValidateWeeklyAmount(t *testing.T){
	var maxAmountPerWeekValidator MaxAmountPerWeekValidator
	e:=EventLogEntry{CustomerId:customerId, Id:"123",Amount:"$2933.34", EventTime:"2019-07-01T12:34:56Z"}
	deposits:=map[string]map[int][]float64{}

	deposits["2019-07-01"]=map[int][]float64{}
	deposits["2019-07-01"][1] = []float64{2234.56,2345.45}
	deposits["2019-07-01"][2] = []float64{2234.56,2345.45}
	deposits["2019-07-01"][3] = []float64{2234.56,2345.45}


	customerData := CustomerData{Transactions:[]string{"123","456","111","112","113","114"},
		CustomerId:customerId,
		Deposits: deposits}
	v := maxAmountPerWeekValidator.validate(customerData, e)
	if v != true{
		t.Error("validator failed - should have added more momney per day")
	}
}
func TestShouldFailWeeklyAmount(t *testing.T){
	var maxAmountPerWeekValidator MaxAmountPerWeekValidator
	e:=EventLogEntry{CustomerId:customerId, Id:"123",Amount:"$2933.34", EventTime:"2019-07-01T12:34:56Z"}
	deposits:=map[string]map[int][]float64{}
	deposits["2019-07-01"]=map[int][]float64{}
	deposits["2019-07-01"][1] = []float64{2234.56,2345.45}
	deposits["2019-07-01"][2] = []float64{2234.56,2345.45}
	deposits["2019-07-01"][3] = []float64{5000}
	deposits["2019-07-01"][4] = []float64{5000}
	deposits["2019-07-01"][5] = []float64{5000}


	customerData := CustomerData{Transactions:[]string{"123","456","111","112","113","114","222"},
		CustomerId:customerId,
		Deposits: deposits}
	v := maxAmountPerWeekValidator.validate(customerData, e)
	if v == true{
		t.Error("validator failed - should not allow another txn too much $ per day")
	}
}