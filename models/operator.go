package models

type Operator struct {
	Base
	name  string
	empID int
}

// func (o *Operator) GetID() int {
// 	return o.id
// }

func (o *Operator) GetName() string {
	return o.name
}

func (o *Operator) GetEmpID() int {
	return o.empID
}

// func (o *Operator) SetID(id int) {
// 	o.id = id
// }

func (o *Operator) SetName(name string) {
	o.name = name
}

func (o *Operator) SetEmpID(empID int) {
	o.empID = empID
}
