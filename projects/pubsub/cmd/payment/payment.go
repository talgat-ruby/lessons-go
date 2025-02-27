package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/redis/go-redis/v9"
)

type PaymentService struct {
	redisClient *redis.Client
	config      *Config
}

func NewPaymentService(cfg *Config) *PaymentService {
	client := NewClient(cfg)

	return &PaymentService{
		redisClient: client,
		config:      cfg,
	}
}

func (ps *PaymentService) ProcessPayment(amount float64, customerID, paymentMethod string) (*Transaction, error) {
	fmt.Printf("Processing payment of $%.2f for customer %s...\n", amount, customerID)
	time.Sleep(1 * time.Second) // Simulate processing time

	// Create transaction record
	transaction := NewTransaction(amount, customerID, paymentMethod)

	// Publish transaction completed event
	jsonMsg, err := transaction.ToJSON()
	if err != nil {
		return nil, fmt.Errorf("failed to serialize transaction: %w", err)
	}

	ctx := context.Background()
	count, err := ps.redisClient.Publish(ctx, ps.config.RedisChannel, jsonMsg).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to publish message: %w", err)
	}

	fmt.Printf("Transaction %s completed\n", transaction.TransactionID)
	fmt.Printf("Message published to %d subscribers\n", count)

	return transaction, nil
}

func (ps *PaymentService) SimulatePayments(count int) {
	paymentMethods := []string{"credit_card", "debit_card", "bank_transfer", "wallet"}

	for i := 0; i < count; i++ {
		amount := roundToTwoDecimals(rand.Float64()*490 + 10) // Random amount between 10 and 500
		customerID := fmt.Sprintf("cust_%d", rand.Intn(9000)+1000)
		paymentMethod := paymentMethods[rand.Intn(len(paymentMethods))]

		_, err := ps.ProcessPayment(amount, customerID, paymentMethod)
		if err != nil {
			log.Printf("Error processing payment: %v", err)
		}

		time.Sleep(2 * time.Second) // Wait between payments
	}
}

func (ps *PaymentService) Close() {
	ps.redisClient.Close()
}

func roundToTwoDecimals(num float64) float64 {
	return float64(int(num*100)) / 100
}
