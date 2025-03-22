package main

import (
	"fmt"
	"log"

	"github.com/Zhima-Mochi/mailmole"
)

func main() {
	// Create a tunnel agent using the factory pattern
	options := &mailmole.TunnelOptions{
		BrowserOptions: &mailmole.BrowserOptions{
			Headless: false,
		},
	}

	// Create a SmailPro tunnel agent
	agent, err := mailmole.CreateTunnel(mailmole.SmailProTunnel, options)
	if err != nil {
		log.Fatalf("Failed to create tunnel agent: %v", err)
	}

	// Initialize the tunnel agent
	err = agent.Init()
	if err != nil {
		log.Fatalf("Failed to initialize tunnel agent: %v", err)
	}

	// Ensure resources are cleaned up when done
	defer agent.Close()

	// Get the temporary email address
	email := agent.EmailAddress()
	fmt.Println("Temporary email address:", email)

	fmt.Println("Checking for verification email...")

	// Try checking for the verification email a few times
	const maxAttempts = 5
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		fmt.Printf("Attempt %d/%d\n", attempt, maxAttempts)

		code, err := agent.GetVerificationCode()
		if err != nil {
			log.Printf("Failed to extract verification code: %v", err)
			continue
		}

		fmt.Println("Verification code:", code)
		return
	}

	fmt.Println("No verification email found after", maxAttempts, "attempts")
}
