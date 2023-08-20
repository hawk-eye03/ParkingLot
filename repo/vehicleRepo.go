package repo

import "github.com/hawk-eye03/ParkingLot/models"

type VehicleRepo struct {
}

func (v *VehicleRepo) GetVehicleByNumber(number string) models.Vehicle {
	return models.Vehicle{}
}

func (v *VehicleRepo) CreateVehcile(vehicle models.Vehicle) {

}
