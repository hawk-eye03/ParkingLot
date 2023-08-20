package spotassignmentstrategy

import "github.com/hawk-eye03/ParkingLot/models"

type SpotAssignmentStrategy interface {
	FindSpot(vehicleType models.VehicleType, parkingLot models.ParkingLot) *models.ParkingSpot
}
