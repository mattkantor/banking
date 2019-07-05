package main



const maxItemsPerDay = 3
const maxAmountPerDay = 5000
const maxAmountPerWeek = 20000

// This is a little on the "java-y" side but its a nice way to set up your validators.  DI is another approach
// why use an if statement when youc an dynamically inject a bunch of struts with methods into manager?

type Validator interface {
	 validate(data CustomerData, input EventLogEntry) bool
}

type ValidationManager struct {
	validators []Validator

}

func NewValidationManager() *ValidationManager{
	vm := ValidationManager{}

	return &vm
}



func (vm *ValidationManager)addValidaator(v Validator){
	vm.validators = append(vm.validators, v)
}

func (vm *ValidationManager) IsValid(customerRecords CustomerData, entry EventLogEntry) bool{
	//if no validators, return true
	for i:=0;i<len(vm.validators);i++{
		if !vm.validators[i].validate(customerRecords,entry){
			return false
		}
	}
	return true
}

type MaxItemsPerDayValidator struct{}
type MaxAmountPerDayValidator struct {}
type MaxAmountPerWeekValidtor struct {}


func (m *MaxItemsPerDayValidator )validate(data CustomerData, e EventLogEntry) bool{
	monday, dayIndex :=  GetMondayAndoffsetForDate(ParseFileDateIntoRealDate(e.EventTime))
	items := data.Deposits.MondayDate[monday].Day[dayIndex]
	return len(items) <= maxItemsPerDay
}

func (m *MaxAmountPerDayValidator )validate(data CustomerData, e EventLogEntry) bool{
	monday, dayIndex :=  GetMondayAndoffsetForDate(ParseFileDateIntoRealDate(e.EventTime))
	items := data.Deposits.MondayDate[monday].Day[dayIndex]

	return sum(items)  + e.Amount < maxAmountPerDay
}

func (m *MaxAmountPerWeekValidtor )validate(data CustomerData, e EventLogEntry) bool{
	monday, _ :=  GetMondayAndoffsetForDate(ParseFileDateIntoRealDate(e.EventTime))
	items := data.Deposits.MondayDate[monday].Day
	var total float64 = 0.0
	for i:=0;i<len(items);i++{
		total += sum(items[i])
	}
	return total  + e.Amount < maxAmountPerDay

}



//----------
//
//type TransactionExistsValidator struct {}
//
//func (m *TransactionExistsValidator) validte(customerRecords CustomerData,txnId string ) bool {
//	transactions := customerRecords.Transactions
//	for i := 0; i < len(transactions); i++ {
//
//		if transactions[i] == txnId {
//			return true
//		}
//	}
//	return false
//}
