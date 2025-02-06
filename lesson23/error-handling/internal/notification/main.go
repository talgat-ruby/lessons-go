package notification

import (
	"fmt"
)

type Notification struct {
}

func NewNotification() *Notification {
	return &Notification{}
}

type NotifyClientReq struct {
	Description string
	Method      string
}

type NotifyClientResp struct{}

func (n *Notification) NotifyClient(req NotifyClientReq) (NotifyClientResp, error) {
	if req.Method == "email" {
		return NotifyClientResp{}, nil
	}
	return NotifyClientResp{}, fmt.Errorf("invalid notification method")
}
