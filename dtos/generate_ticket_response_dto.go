package dtos

type GenerateTicketResponseDto struct {
	ticketID       int64
	operatorName   string
	spotNumber     int
	responseStatus ResponseStatus
}

func (r GenerateTicketResponseDto) GetTicketID() int64 {
	return r.ticketID
}

func (r GenerateTicketResponseDto) GetOperatorName() string {
	return r.operatorName
}

func (r GenerateTicketResponseDto) GetSpotNumber() int {
	return r.spotNumber
}

func (r GenerateTicketResponseDto) GetResponseStatus() ResponseStatus {
	return r.responseStatus
}

// Setter methods

func (r *GenerateTicketResponseDto) SetTicketID(id int64) {
	r.ticketID = id
}

func (r *GenerateTicketResponseDto) SetOperatorName(name string) {
	r.operatorName = name
}

func (r *GenerateTicketResponseDto) SetSpotNumber(number int) {
	r.spotNumber = number
}

func (r *GenerateTicketResponseDto) SetResponseStatus(status ResponseStatus) {
	r.responseStatus = status
}
