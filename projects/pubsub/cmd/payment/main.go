package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Load configuration
	cfg := LoadConfig()

	// Create and start payment service
	paymentService := NewPaymentService(cfg)
	defer paymentService.Close()

	h := &Handler{payment: paymentService}

	fmt.Println("Payment Service started. Processing transactions...")
	http.HandleFunc("POST /api/process", h.ProcessAmount)

	// Start the server
	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Server starting on %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
	fmt.Println("Payment Service shutting down...")
}

type Handler struct {
	payment *PaymentService
}

type AmountRequest struct {
	Amount int `json:"amount"`
}

type AmountResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (h *Handler) ProcessAmount(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var request AmountRequest
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&request); err != nil {
		log.Printf("Error decoding request: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Process the amount using service
	h.payment.SimulatePayments(request.Amount)

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := AmountResponse{
		Success: true,
		Message: "Payment was sent successfully",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
