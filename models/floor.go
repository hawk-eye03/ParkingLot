package models

type ParkingFloor struct {
	Base
	id           int
	floorID      int
	parkingSpots []ParkingSpot
}

func (pf *ParkingFloor) GetID() int {
	return pf.id
}

func (pf *ParkingFloor) GeFloortID() int {
	return pf.floorID
}

func (pf *ParkingFloor) GetParkingSlots() []ParkingSpot {
	return pf.parkingSpots
}

func (pf *ParkingFloor) SetID(id int) {
	pf.id = id
}

func (pf *ParkingFloor) SetFloortID(floorID int) {
	pf.floorID = floorID
}

func (pf *ParkingFloor) SetParkingSlots(parkingSlots []ParkingSpot) {
	pf.parkingSpots = parkingSlots
}
