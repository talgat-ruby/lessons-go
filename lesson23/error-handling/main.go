package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"github.com/talgat-ruby/lessons-go/lesson23/error-handling/internal/delivery"
	"github.com/talgat-ruby/lessons-go/lesson23/error-handling/internal/notification"
	"github.com/talgat-ruby/lessons-go/lesson23/error-handling/internal/payment"
	"github.com/talgat-ruby/lessons-go/lesson23/error-handling/pkg/httputils/request"
	"github.com/talgat-ruby/lessons-go/lesson23/error-handling/pkg/httputils/response"
)

func main() {
	ctx := context.Background()

	mux := http.NewServeMux()

	api := API{
		payment:      payment.NewPayment(),
		notification: notification.NewNotification(),
		delivery:     delivery.NewDelivery(),
	}

	mux.HandleFunc("POST /product", api.handlePurchase)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 5523),
		Handler: mux,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

type API struct {
	payment      *payment.Payment
	notification *notification.Notification
	delivery     *delivery.Delivery
}

func (a *API) handlePurchase(w http.ResponseWriter, r *http.Request) {
	reqBody := &purchaseReq{}
	if err := request.JSON(w, r, reqBody); err != nil {
		slog.Error(
			"failed to parse request body",
			slog.Any("error", err),
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	paymentS := payment.NewPayment()
	var payReq *payment.PayReq

	_, err := paymentS.Pay(payReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("error from pay %v", err), http.StatusBadRequest)
		return
	}

	// db product
	//dbResp := &struct {
	//	ID            string
	//	Name          string
	//	Price         float64
	//	Total         int
	//	ClientAddress string
	//}{
	//	ID:            reqBody.Data.ID,
	//	Name:          "ps",
	//	Price:         3000,
	//	Total:         3000,
	//	ClientAddress: "Narnia",
	//}

	//paymentResp, err := a.payment.Pay(payment.PayReq{Amount: dbResp.Price, Good: dbResp.Name})
	//if err != nil {
	//	http.Error(w, "failed to pay the payment", http.StatusPaymentRequired)
	//}
	//
	//notificationResp, err := a.notification.NotifyClient(notification.NotifyClientReq{Description: "purchase goods", Method: "plane"})
	//if err != nil {
	//	delErr := &my_error.DeliveryError{}
	//	switch {
	//	case errors.As(err, &delErr):
	//		err := err.(*my_error.DeliveryError)
	//		http.Error(w, fmt.Sprintf("error from delivery %v, delivery method: %s", err, err.Method), http.StatusBadRequest)
	//		return
	//	}
	//	http.Error(w, "failed to pay the payment", http.StatusPaymentRequired)
	//	return
	//}

	if err := response.JSON(
		w,
		http.StatusCreated,
		nil,
	); err != nil {
		slog.Error(
			"fail json",
			slog.Any("error", err),
		)
		return
	}

	slog.Info(
		"success",
	)
}

type purchaseReq struct {
	Data *purchaseReqData `json:"data"`
}

type purchaseReqData struct {
	ID string `json:"id"`
}
