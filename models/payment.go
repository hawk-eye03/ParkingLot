// models/payment.go
package models

type Payment struct {
	Base
	mode   PaymentMode   // ONLINE, OFFLINE
	status PaymentStatus // FAILED, SUCCESS
	refID  string
	amount int
}

// func NewPayment(id int, mode, status, refID string, amount int) *Payment {
// 	return &Payment{
// 		id:     id,
// 		mode:   mode,
// 		status: status,
// 		refID:  refID,
// 		amount: amount,
// 	}
// }

func (p *Payment) GetID() int64 {
	return p.id
}

func (p *Payment) SetID(id int64) {
	p.id = id
}

func (p *Payment) GetMode() PaymentMode {
	return p.mode
}

func (p *Payment) SetMode(mode PaymentMode) {
	p.mode = mode
}

func (p *Payment) GetStatus() PaymentStatus {
	return p.status
}

func (p *Payment) SetStatus(status PaymentStatus) {
	p.status = status
}

func (p *Payment) GetRefID() string {
	return p.refID
}

func (p *Payment) SetRefID(refID string) {
	p.refID = refID
}

func (p *Payment) GetAmount() int {
	return p.amount
}

func (p *Payment) SetAmount(amount int) {
	p.amount = amount
}
