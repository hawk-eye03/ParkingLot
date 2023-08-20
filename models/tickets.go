// models/ticket.go
package models

import (
	"time"
)

type Ticket struct {
	Base
	entryTime   time.Time
	operator    Operator // operator who generated this ticket
	vehicle     Vehicle
	gate        Gate
	parkingSpot ParkingSpot
}

func (t *Ticket) GetStartTime() time.Time {
	return t.entryTime
}

func (t *Ticket) SetStartTime(startTime time.Time) {
	t.entryTime = startTime
}

func (t *Ticket) GetOperator() Operator {
	return t.operator
}

func (t *Ticket) SetOperator(operator Operator) {
	t.operator = operator
}

func (t *Ticket) GetVehicle() Vehicle {
	return t.vehicle
}

func (t *Ticket) SetVehicle(vehicle Vehicle) {
	t.vehicle = vehicle
}

func (t *Ticket) GetGate() Gate {
	return t.gate
}

func (t *Ticket) SetGate(gate Gate) {
	t.gate = gate
}

func (t *Ticket) GetParkingSpot() ParkingSpot {
	return t.parkingSpot
}

func (t *Ticket) SetParkingSpot(parkingSlot ParkingSpot) {
	t.parkingSpot = parkingSlot
}

func (t *Ticket) GetOperatorName() string {
	return t.operator.name
}

func (t *Ticket) SetOperatorName(name string) {
	t.operator.name = name
}
