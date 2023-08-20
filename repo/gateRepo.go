package repo

import "github.com/hawk-eye03/ParkingLot/models"

type GateRepo struct {
	gateMap map[int64]models.Gate
}

func (g *GateRepo) FindGateByID(id int64) models.Gate {
	if _, ok := g.gateMap[id]; !ok {
		return models.Gate{}
	}
	return g.gateMap[id]
}
