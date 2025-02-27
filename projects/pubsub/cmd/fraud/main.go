package main

func main() {
	// Load configuration
	cfg := LoadConfig()

	// Create and start fraud detection service
	fraudService := NewFraudDetectionService(cfg)
	defer fraudService.Close()

	fraudService.StartListening()
}
