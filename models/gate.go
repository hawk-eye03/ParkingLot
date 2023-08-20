package models

type Gate struct {
	Base
	gateNo   int64
	gateType GateType   // ENTRY, EXIT
	operator Operator   // Current Operator at gate
	status   GateStatus // OPENED, CLOSED
}

func (g *Gate) GetID() int64 {
	return g.id
}

func (g *Gate) GetGateNo() int64 {
	return g.gateNo
}

func (g *Gate) GetGateType() GateType {
	return g.gateType
}

func (g *Gate) GetOperator() Operator {
	return g.operator
}

func (g *Gate) GetGateStatus() GateStatus {
	return g.status
}

// func (g *Gate) SetID(id int) {
// 	g.id = id
// }

func (g *Gate) SetGateID(gateID int64) {
	g.gateNo = gateID
}

func (g *Gate) SetGateType(gateType GateType) {
	g.gateType = gateType
}

func (g *Gate) SetOperator(operator Operator) {
	g.operator = operator
}

func (g *Gate) SetGateStatus(status GateStatus) {
	g.status = status
}
