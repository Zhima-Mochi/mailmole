# Mailmole

![Mailmole Logo](./docs/mailmole.jpeg)


Mailmole is a Go package that provides a unified interface for handling email verification through temporary email services, we call them "**tunnels**". It simplifies the process of receiving verification codes and managing temporary email addresses for testing and automation purposes.

## Features

- Unified interface for multiple temporary email services
- Browser automation support
- Configurable options for both email and browser settings
- Easy-to-use API for email verification workflows
- Support for headless browser mode

## Installation

```bash
go get github.com/Zhima-Mochi/mailmole
```

## Quick Start

Here's a simple example of how to use Mailmole:

```go
package main

import (
    "fmt"
    "log"
    "github.com/Zhima-Mochi/mailmole"
)

func main() {
    // Create a tunnel agent with options
    options := mailmole.TunnelOptions{
        BrowserOptions: &mailmole.BrowserOptions{
            Headless: false,
        },
    }

    // Create a SmailPro tunnel agent
    agent, err := mailmole.CreateTunnel(mailmole.SmailPro, options)
    if err != nil {
        log.Fatalf("Failed to create tunnel agent: %v", err)
    }

    // Initialize the tunnel agent
    err = agent.Init()
    if err != nil {
        log.Fatalf("Failed to initialize tunnel agent: %v", err)
    }

    // Clean up resources when done
    defer agent.Close()

    // Get the temporary email address
    email := agent.EmailAddress()
    fmt.Println("Temporary email address:", email)

    // Get verification code
    code, err := agent.GetVerificationCode()
    if err != nil {
        log.Fatalf("Failed to get verification code: %v", err)
    }
    fmt.Println("Verification code:", code)
}
```

## Configuration

### TunnelOptions

```go
type TunnelOptions struct {
    // Email to use (optional, for Yopmail)
    Email string
    // Browser options
    BrowserOptions *BrowserOptions
}
```

### BrowserOptions

```go
type BrowserOptions struct {
    Headless bool   // Whether to run browser in headless mode
    URL      string // Custom URL to navigate to
}
```

## Supported Services

Currently supported temporary email services:
- SmailPro
- Yopmail

Welcome to contribute more services!

## Dependencies

- Go 1.22 or later
- github.com/go-rod/rod (for browser automation)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. 