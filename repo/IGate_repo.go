package repo

import "github.com/hawk-eye03/ParkingLot/models"

type ITicketRepo interface {
	FindGateByID(id int64) models.Gate
}
