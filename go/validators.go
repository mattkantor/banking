package main


const maxItemsPerDay = 3
const maxAmountPerDay = 5000
const maxAmountPerWeek = 20000

type MaxItemsPerDayValidator struct{}
type MaxAmountPerDayValidator struct {}
type MaxAmountPerWeekValidator struct {}


func (m *MaxItemsPerDayValidator )validate(data CustomerData, e EventLogEntry) bool{

	monday, dayIndex :=  GetMondayAndoffsetForDate(ParseFileDateIntoRealDate(e.EventTime))
	if mondayItems, ok := data.Deposits[monday]; ok {
		if dayItems, ok := mondayItems[dayIndex];ok{
			return  len(dayItems) < maxItemsPerDay

		}
	}
	return true
}

func (m *MaxAmountPerDayValidator )validate(data CustomerData, e EventLogEntry) bool{
	monday, dayIndex :=  GetMondayAndoffsetForDate(ParseFileDateIntoRealDate(e.EventTime))
	if mondayItems, ok := data.Deposits[monday]; ok {
		if dayItems, ok := mondayItems[dayIndex];ok{
			return sum(dayItems)  + CleanCurrency(e.Amount) <= maxAmountPerDay
		}
	}
	return true
}

func (m *MaxAmountPerWeekValidator )validate(data CustomerData, e EventLogEntry) bool{
	monday, _ :=  GetMondayAndoffsetForDate(ParseFileDateIntoRealDate(e.EventTime))
	if items, ok := data.Deposits[monday]; ok {
		var total float64 = 0.0
		for i:=0;i<len(items);i++{
			total += sum(items[i])
		}
		return total  + CleanCurrency(e.Amount) <= maxAmountPerWeek
	}else{
		return true
	}




}
