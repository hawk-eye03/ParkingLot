package spotassignmentstrategy

import "github.com/hawk-eye03/ParkingLot/models"

type RandomSpotAssignmentStrategy struct{}

func (r *RandomSpotAssignmentStrategy) FindSpot(vehicleType models.VehicleType, parkingLot models.ParkingLot) *models.ParkingSpot {
	for _, parkingFloor := range parkingLot.GetFloors() {
		for _, parkingSpot := range parkingFloor.GetParkingSlots() {
			for _, vehicleSpotType := range parkingSpot.GetVehicleTypes() {
				if vehicleType == vehicleSpotType && parkingSpot.GetStatus() == models.AVAILABLE {
					return &parkingSpot
				}
			}
		}
	}
	return nil
}
