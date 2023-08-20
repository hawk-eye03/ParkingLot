package models

type ParkingSpot struct {
	Base
	parkingID      int
	spotStatus     SpotStatus    // EMPTY, FILLED
	vehicleTypes   []VehicleType // Supported vehicle types
	currentVehicle Vehicle       // why are we storing it?
}

// func (ps *ParkingSpot) GetID() int {
// 	return ps.id
// }

func (ps *ParkingSpot) GetParkingID() int {
	return ps.parkingID
}

func (ps *ParkingSpot) GetStatus() SpotStatus {
	return ps.spotStatus
}

func (ps *ParkingSpot) GetVehicleTypes() []VehicleType {
	return ps.vehicleTypes
}

func (ps *ParkingSpot) GetCurrentVehicle() Vehicle {
	return ps.currentVehicle
}

// func (ps *ParkingSpot) SetID(id int) {
// 	ps.id = id
// }

func (ps *ParkingSpot) SetParkingID(parkingID int) {
	ps.parkingID = parkingID
}

func (ps *ParkingSpot) SetStatus(status SpotStatus) {
	ps.spotStatus = status
}

func (ps *ParkingSpot) SetVehicleTypes(vehicleTypes []VehicleType) {
	ps.vehicleTypes = vehicleTypes
}

func (ps *ParkingSpot) SetCurrentVehicle(currentVehicle Vehicle) {
	ps.currentVehicle = currentVehicle
}
