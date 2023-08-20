package dtos

import "github.com/hawk-eye03/ParkingLot/models"

type GenerateTicketRequestDto struct {
	vehicleNumber string
	vehicleType   models.VehicleType
	gateID        int64
}

func (r *GenerateTicketRequestDto) SetGateID(id int64) {
	r.gateID = id
}

func (r GenerateTicketRequestDto) GetGateID() int64 {
	return r.gateID
}

func (r GenerateTicketRequestDto) GetVehicleNumber() string {
	return r.vehicleNumber
}

func (r GenerateTicketRequestDto) GetVehicleType() models.VehicleType {
	return r.vehicleType
}

func (r *GenerateTicketRequestDto) SetVehicleNumber(number string) {
	r.vehicleNumber = number
}

func (r *GenerateTicketRequestDto) SetVehicleType(vehicleType models.VehicleType) {
	r.vehicleType = vehicleType
}
