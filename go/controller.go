package main

type CustomerController interface {
	AddDeposit(e EventLogEntry)  (ResultLogEntry, int)
	//CheckForExistingTransaction()
	//GetDepositsForDate()
	//GetDepositsForWeek()
	//GetCustomerHistory()
}

type KohoCustomerController struct {
	DbManager DBManager
}

func (cc *KohoCustomerController) AddDeposit(e EventLogEntry) (ResultLogEntry, int){
	output, err := cc.DbManager.AddDeposit(e)
	return output, err
}

func NewCustomerController() CustomerController {
	return &KohoCustomerController{}
}
