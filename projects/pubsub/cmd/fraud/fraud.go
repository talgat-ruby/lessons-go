package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/redis/go-redis/v9"
)

type FraudDetectionService struct {
	redisClient          *redis.Client
	config               *Config
	suspiciousThresholds map[string]float64
}

func NewFraudDetectionService(cfg *Config) *FraudDetectionService {
	client := NewClient(cfg)

	thresholds := map[string]float64{
		"credit_card":   300.0,
		"debit_card":    250.0,
		"bank_transfer": 1000.0,
		"wallet":        150.0,
	}

	return &FraudDetectionService{
		redisClient:          client,
		config:               cfg,
		suspiciousThresholds: thresholds,
	}
}

func (fds *FraudDetectionService) AnalyzeTransaction(transaction *Transaction) bool {
	fmt.Printf("\n[FRAUD DETECTION] Analyzing transaction %s\n", transaction.TransactionID)

	threshold, exists := fds.suspiciousThresholds[transaction.PaymentMethod]
	if !exists {
		threshold = 200.0 // Default threshold
	}

	isSuspicious := transaction.Amount > threshold

	if isSuspicious {
		fmt.Println("ALERT: Suspicious transaction detected!")
		fmt.Printf("Amount $%.2f exceeds threshold of $%.2f for %s\n",
			transaction.Amount, threshold, transaction.PaymentMethod)
		fmt.Printf("Customer ID: %s\n", transaction.CustomerID)
		fmt.Println("Flagging for review...\n")
	} else {
		fmt.Println("Transaction appears legitimate. No action needed.\n")
	}

	return isSuspicious
}

func (fds *FraudDetectionService) StartListening() {
	ctx := context.Background()
	pubsub := fds.redisClient.Subscribe(ctx, fds.config.RedisChannel)
	defer pubsub.Close()

	fmt.Println("Fraud Detection Service started. Listening for transactions. Press Ctrl+C to exit.")

	// Channel to receive messages
	ch := pubsub.Channel()

	// Channel to handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Process messages until interrupted
	for {
		select {
		case msg := <-ch:
			transaction, err := TransactionFromJSON(msg.Payload)
			if err != nil {
				log.Printf("[ERROR] Failed to parse transaction: %v", err)
				continue
			}
			fds.AnalyzeTransaction(transaction)

		case <-sigChan:
			fmt.Println("Fraud Detection Service shutting down...")
			return
		}
	}
}

func (fds *FraudDetectionService) Close() {
	fds.redisClient.Close()
}
