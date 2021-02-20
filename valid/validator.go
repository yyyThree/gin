package valid

import (
	"github.com/asaskevich/govalidator"
)

func (myValid *myValid) IsNotNull() *myValid {
	myValid.validations["IsNotNull"] = func() bool {
		return govalidator.IsNotNull(govalidator.ToString(myValid.value))
	}
	return myValid
}

func (myValid *myValid) IsIn(params ...string) *myValid {
	myValid.validations["IsIn"] = func() bool {
		return govalidator.IsIn(govalidator.ToString(myValid.value), params...)
	}
	return myValid
}