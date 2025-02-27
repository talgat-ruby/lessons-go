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

type EmailService struct {
	redisClient *redis.Client
	config      *Config
}

func NewEmailService(cfg *Config) *EmailService {
	client := NewClient(cfg)

	return &EmailService{
		redisClient: client,
		config:      cfg,
	}
}

func (es *EmailService) SendReceiptEmail(transaction *Transaction) {
	fmt.Printf("\n[EMAIL SERVICE] Sending receipt email for transaction %s\n", transaction.TransactionID)
	fmt.Printf("To: Customer %s\n", transaction.CustomerID)
	fmt.Printf("Subject: Receipt for your payment of $%.2f\n", transaction.Amount)
	fmt.Printf("Thank you for your payment of $%.2f via %s.\n", transaction.Amount, transaction.PaymentMethod)
	fmt.Printf("Transaction ID: %s\n", transaction.TransactionID)
	fmt.Printf("Timestamp: %s\n\n", transaction.Timestamp)
}

func (es *EmailService) StartListening() {
	ctx := context.Background()
	pubsub := es.redisClient.Subscribe(ctx, es.config.RedisChannel)
	defer pubsub.Close()

	fmt.Println("Email Service started. Listening for transactions. Press Ctrl+C to exit.")

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
			es.SendReceiptEmail(transaction)

		case <-sigChan:
			fmt.Println("Email Service shutting down...")
			return
		}
	}
}

func (es *EmailService) Close() {
	es.redisClient.Close()
}
