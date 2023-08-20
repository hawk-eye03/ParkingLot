package models

import "time"

type ParkingLot struct {
	Base
	capacity int
	floors   []ParkingFloor
	gates    []Gate
}

type Base struct {
	id            int64
	createdAt     time.Time
	lastUpdatedAt time.Time
}

func (b *Base) GetID() int64 {
	return b.id
}

func (b *Base) SetID(id int64) {
	b.id = id
}

func (b *Base) GetCreatedAt() time.Time {
	return b.createdAt
}

func (b *Base) SetCreatedAt(createdAt time.Time) {
	b.createdAt = createdAt
}

func (b *Base) GetLastUpdatedAt() time.Time {
	return b.lastUpdatedAt
}

func (b *Base) SetLastUpdatedAt(lastUpdatedAt time.Time) {
	b.lastUpdatedAt = lastUpdatedAt
}

func (p *ParkingLot) GetCapacity() int {
	return p.capacity
}

func (p *ParkingLot) GetFloors() []ParkingFloor {
	return p.floors
}

func (p *ParkingLot) GetGates() []Gate {
	return p.gates
}

// func (p *ParkingLot) SetID(id int) {
// 	p.id = id
// }

func (p *ParkingLot) SetCapacity(capacity int) {
	p.capacity = capacity
}

func (p *ParkingLot) SetFloors(floors []ParkingFloor) {
	p.floors = floors
}

func (p *ParkingLot) SetGates(gates []Gate) {
	p.gates = gates
}
