package controllers

import (
	"github.com/hawk-eye03/ParkingLot/dtos"
	"github.com/hawk-eye03/ParkingLot/services"
)

type TicketController struct {
	ticketServ services.TicketService
}

func NewTicketController(ts services.TicketService) *TicketController {
	return &TicketController{
		ticketServ: ts,
	}
}

func (t *TicketController) GenerateTicket(req dtos.GenerateTicketRequestDto) *dtos.GenerateTicketResponseDto {
	vehicleType := req.GetVehicleType()
	vehicleNumber := req.GetVehicleNumber()
	gateID := req.GetGateID()

	ticket := t.ticketServ.GenerateTicket(gateID, vehicleType, vehicleNumber)
	if ticket == nil {
		// if error received from service
		return nil
	}

	ticketResp := dtos.GenerateTicketResponseDto{}
	ticketResp.SetOperatorName(ticket.GetOperatorName())
	parkingSpot := ticket.GetParkingSpot()
	ticketResp.SetSpotNumber(parkingSpot.GetParkingID())
	ticketResp.SetResponseStatus(dtos.SUCCESS)
	ticketResp.SetTicketID(ticket.GetID())

	return &ticketResp
}
