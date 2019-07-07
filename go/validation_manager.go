package main




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

func (vm *ValidationManager)IsValid(customerRecords CustomerData, entry EventLogEntry) bool{
	//if no validators, return true
	for i:=0;i<len(vm.validators);i++{
		if !vm.validators[i].validate(customerRecords,entry){
			return false
		}
	}
	return true
}


