package delivery

import (
	"fmt"

	"github.com/talgat-ruby/lessons-go/lesson23/error-handling/internal/pkg/my_error"
)

type Delivery struct {
}

func NewDelivery() *Delivery {
	return &Delivery{}
}

type SendToClientReq struct {
	Address string
	Method  string
}

type SendToClientResp struct{}

func (d *Delivery) SendToClient(req SendToClientReq) (SendToClientResp, error) {
	if req.Method == "train" {
		return SendToClientResp{}, nil
	}
	return SendToClientResp{}, my_error.NewDeliveryError(fmt.Errorf("we don't support this method"), req.Method)
}
