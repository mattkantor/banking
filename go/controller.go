package main

type CustomerController struct {
	DbManager *DBManager
	Vm *ValidationManager
}


func (cc *CustomerController) LoadCard(e EventLogEntry) (accepted bool, code int){

	cd := cc.DbManager.getCustomerData(e.CustomerId)
	customerTransactions := cd.Transactions
	transactionId := e.Id

	// dont do anything if the id is already used

	/*
	return codes
	200 - ok
	406 - not ok
	500 - broken not ok (shouldn't really get here)
	403 - duplicate ID

 */
	if contains(transactionId, customerTransactions){
		return false, 403
	}
	//in theory it would be nice to do it in a more rollback type fashion
	cc.DbManager.recordTransaction(e.CustomerId,e.Id)

	if cc.Vm.IsValid(cd, e) {
		ok := cc.DbManager.recordLoad(e)
		if !ok {
			return false, 500
		}else{
			return true, 200
		}
	}else{
		return false, 406
	}



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
