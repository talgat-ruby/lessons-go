package main

func main() {
	// Load configuration
	cfg := LoadConfig()

	// Create and start email service
	emailService := NewEmailService(cfg)
	defer emailService.Close()

	emailService.StartListening()
}
