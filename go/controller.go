package main


type CustomerController struct {
	DbManager *DBManager
	Vm *ValidationManager
}


func (cc *CustomerController) AddDeposit(e EventLogEntry) (accepted bool, code int){



	cd := cc.DbManager.getCustomerData(e.CustomerId)

	customerTransactions := cd.Transactions
	transactionId := e.Id

	if contains(transactionId,customerTransactions){
		return false, 403
	}

	if cc.Vm.IsValid(cd, e) {
		ok := cc.DbManager.loadAccount(e)
		if !ok {

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
	vm.addValidaator(&MaxAmountPerWeekValidator{})
	return CustomerController{
		DbManager:Dbm,
		Vm: vm,

	}
}
