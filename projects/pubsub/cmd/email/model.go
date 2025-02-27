package main

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	TransactionID string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
	CustomerID    string  `json:"customer_id"`
	PaymentMethod string  `json:"payment_method"`
	Timestamp     string  `json:"timestamp"`
}

func NewTransaction(amount float64, customerID, paymentMethod string) *Transaction {
	return &Transaction{
		TransactionID: uuid.New().String(),
		Amount:        amount,
		CustomerID:    customerID,
		PaymentMethod: paymentMethod,
		Timestamp:     time.Now().Format(time.RFC3339),
	}
}

func (t *Transaction) ToJSON() (string, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func TransactionFromJSON(jsonStr string) (*Transaction, error) {
	var transaction Transaction
	err := json.Unmarshal([]byte(jsonStr), &transaction)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}
