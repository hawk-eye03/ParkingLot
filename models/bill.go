package models

import (
	"time"
)

type Bill struct {
	Base
	amount   int
	operator Operator // operator generated the bill
	// vehicle  Vehicle
	gate     Gate // the gate where the bill is generated
	exitTime time.Time
	ticket   Ticket
}

// func NewBill(id, billID, amount int, operator Operator, gate Gate, exitTime time.Time) *Bill {
// 	return &Bill{
// 		id:       id,
// 		billID:   billID,
// 		amount:   amount,
// 		operator: operator,
// 		gate:     gate,
// 		exitTime: exitTime,
// 	}
// }

// func (b *Bill) GetID() int {
// 	return b.id
// }

// func (b *Bill) SetID(id int) {
// 	b.id = id
// }

// func (b *Bill) GetBillID() int {
// 	return b.billID
// }

// func (b *Bill) SetBillID(billID int) {
// 	b.billID = billID
// }

func (b *Bill) GetAmount() int {
	return b.amount
}

func (b *Bill) SetAmount(amount int) {
	b.amount = amount
}

func (b *Bill) GetOperator() Operator {
	return b.operator
}

func (b *Bill) SetOperator(operator Operator) {
	b.operator = operator
}

func (b *Bill) GetGate() Gate {
	return b.gate
}

func (b *Bill) SetGate(gate Gate) {
	b.gate = gate
}

func (b *Bill) GetExitTime() time.Time {
	return b.exitTime
}

func (b *Bill) SetExitTime(exitTime time.Time) {
	b.exitTime = exitTime
}

func (b *Bill) GetTicket() Ticket {
	return b.ticket
}

func (b *Bill) SetTicket(ticket Ticket) {
	b.ticket = ticket
}
