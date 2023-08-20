package repo

import "github.com/hawk-eye03/ParkingLot/models"

type ParkingLotRepo struct {
	parkingLot map[int64]models.ParkingLot
}

func (pl *ParkingLotRepo) GetParkingLotByGate(gate models.Gate) *models.ParkingLot {
	for _, parkingLot := range pl.parkingLot {
		for _, plGate := range parkingLot.GetGates() {
			if plGate.GetID() == gate.GetID() {
				return &parkingLot
			}
		}
	}
	return nil
}
