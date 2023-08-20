package main

import (
	"fmt"

	"github.com/hawk-eye03/ParkingLot/controllers"
	"github.com/hawk-eye03/ParkingLot/repo"
	"github.com/hawk-eye03/ParkingLot/services"
	"github.com/hawk-eye03/ParkingLot/strategies/spotassignmentstrategy"
)

func main() {
	gateRepo := repo.GateRepo{}
	vehicleRepo := repo.VehicleRepo{}
	spotAssignmentStrategy := &spotassignmentstrategy.RandomSpotAssignmentStrategy{}
	parkingLotRepo := repo.ParkingLotRepo{}
	ticketRepo := repo.TicketRepo{}

	ticketService := services.NewTicketService(gateRepo, vehicleRepo, spotAssignmentStrategy, parkingLotRepo, ticketRepo)
	ticketController := controllers.NewTicketController(*ticketService)

	fmt.Println("TicketController = ", ticketController)
	fmt.Println("Application has started on port: 8080")
}
