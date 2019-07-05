package main

type CustomerController interface {
	AddDeposit(e EventLogEntry)  (ResultLogEntry, int)
	//CheckForExistingTransaction()
	//GetDepositsForDate()
	//GetDepositsForWeek()
	//GetCustomerHistory()
}

type KohoCustomerController struct {
	DbManager *DBManager
	Vm *ValidationManager
}


func (cc *KohoCustomerController) AddDeposit(e EventLogEntry) (accepted bool, code int){

	vm := NewValidationManager()

	cd := cc.DbManager.getCustomerData(e.CustomerId)

	customerTransactions := cd.Transactions
	transactionId := e.TxnId

	if contains(transactionId,customerTransactions){
		return false, 403
	}

	if vm.IsValid(cd, e) {
		err = cc.DbManager.addDeposit(e)
		if err != nil {
			panic(err)
			return false, 500
		}
		return true, 200
	}
	return false, 406
}

func NewCustomerController() CustomerController {
	Dbm:= GetDBManager()
	vm := NewValidationManager()
	vm.addValidaator(&MaxItemsPerDayValidator{})
	vm.addValidaator(&MaxAmountPerDayValidator{})
	vm.addValidaator(&MaxAmountPerWeekValidtor{})
	return &KohoCustomerController{
		DbManager:Dbm,
		Vm: vm,


	}
}
