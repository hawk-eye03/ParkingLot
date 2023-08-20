package services

import "github.com/hawk-eye03/ParkingLot/models"

type ITicketService interface {
	GenerateTicket(gateID int64, vehicleType models.VehicleType, vehicleNumber string) *models.Ticket
}
