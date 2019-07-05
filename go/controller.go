package main


type CustomerController struct {
	DbManager *DBManager
	Vm *ValidationManager
}


func (cc *CustomerController) AddDeposit(e EventLogEntry) (accepted bool, code int){

	vm := NewValidationManager()

	cd := cc.DbManager.getCustomerData(e.CustomerId)

	customerTransactions := cd.Transactions
	transactionId := e.Id

	if contains(transactionId,customerTransactions){
		return false, 403
	}

	if vm.IsValid(cd, e) {
		err := cc.DbManager.loadAccount(e)
		if err {
			panic("Error adding deposit")
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
	return CustomerController{
		DbManager:Dbm,
		Vm: vm,

	}
}
