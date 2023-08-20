package models

type Vehicle struct {
	Base
	vehicleNumber string
	vehicleType   VehicleType
	// currentParkingSlot ParkingSlot  (we can have this as well)?
}

func (g *Vehicle) GetVehicleByNumber() string {
	return g.vehicleNumber
}

func (g *Vehicle) GetVehicleType() VehicleType {
	return g.vehicleType
}

func (g *Vehicle) SetVehicleByNumber(vehicleID string) {
	g.vehicleNumber = vehicleID
}

func (g *Vehicle) SetVehicleType(vehicleType VehicleType) {
	g.vehicleType = vehicleType
}
