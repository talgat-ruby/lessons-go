package payment

import (
	"fmt"
)

type Payment struct {
}

func NewPayment() *Payment {
	return &Payment{}
}

type PayReq struct {
	Amount float64
	Good   string
}

type PayResp struct{}

func (p *Payment) Pay(req *PayReq, _ string) (PayResp, error) {
	if req.Good == "flowers" {
		return PayResp{}, nil
	}
	return PayResp{}, fmt.Errorf("invalid good")
}
