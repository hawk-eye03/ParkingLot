package services

import (
	"time"

	"github.com/hawk-eye03/ParkingLot/models"
	"github.com/hawk-eye03/ParkingLot/repo"
	"github.com/hawk-eye03/ParkingLot/strategies/spotassignmentstrategy"
)

type TicketService struct {
	gateRepo               repo.GateRepo
	vehicleRepo            repo.VehicleRepo
	spotAssignmentStrategy spotassignmentstrategy.SpotAssignmentStrategy
	parkingLotRepo         repo.ParkingLotRepo
	ticketRepo             repo.TicketRepo
}

func NewTicketService(gateRepo repo.GateRepo, vehicleRepo repo.VehicleRepo, spotAssignmentStrategy spotassignmentstrategy.SpotAssignmentStrategy, parkingLotRepo repo.ParkingLotRepo, ticketRepo repo.TicketRepo) *TicketService {
	return &TicketService{
		gateRepo:               gateRepo,
		vehicleRepo:            vehicleRepo,
		spotAssignmentStrategy: spotAssignmentStrategy,
		parkingLotRepo:         parkingLotRepo,
		ticketRepo:             ticketRepo,
	}
}

func (ts *TicketService) GenerateTicket(gateID int64, vehicleType models.VehicleType, vehicleNumber string) *models.Ticket {
	// get gate from id, else throw error if gateID not present
	Gate := ts.gateRepo.FindGateByID(gateID)
	if Gate.GetID() != gateID {
		return nil
	}

	// check if vehicle present, if not create a new entry for vehicle
	Vehicle := ts.vehicleRepo.GetVehicleByNumber(vehicleNumber)
	if Vehicle.GetVehicleByNumber() == "" {
		newVehicle := models.Vehicle{}
		newVehicle.SetVehicleByNumber(vehicleNumber)
		newVehicle.SetVehicleType(vehicleType)
		Vehicle = newVehicle
		ts.vehicleRepo.CreateVehcile(Vehicle)
	}

	parkingLot := ts.parkingLotRepo.GetParkingLotByGate(Gate)
	if parkingLot == nil {
		// parking lot not found for given gate
		return nil
	}
	// get Parking Spot from spot assignment strategy
	Spot := ts.spotAssignmentStrategy.FindSpot(vehicleType, *parkingLot)
	if Spot == nil {
		// no spot available
		// returning empty ticket for now, should return a proper error message
		return nil

	}

	newTicket := models.Ticket{}
	newTicket.SetOperator(Gate.GetOperator())
	newTicket.SetParkingSpot(*Spot)
	newTicket.SetGate(Gate)
	newTicket.SetVehicle(Vehicle)
	newTicket.SetStartTime(time.Now())

	// create new ticket in DB, handle error gracefully, this is just for purpose of understanding
	ts.ticketRepo.CreateTicket(newTicket)
	return &newTicket
}
