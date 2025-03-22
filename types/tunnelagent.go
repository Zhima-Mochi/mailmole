package types

// TunnelOptions holds options for creating a tunnel agent
type TunnelOptions struct {
	// Email to use (optional, for Yopmail)
	Email string
	// Browser options
	BrowserOptions *BrowserOptions
}

// BrowserOptions holds options for creating a browser
type BrowserOptions struct {
	Headless bool
	URL      string
}

// TunnelAgent represents an interface for email verification services
type TunnelAgent interface {
	// Init initializes the tunnel agent (login, open browser, generate email, etc.)
	Init() error

	// EmailAddress returns the current temporary email address
	EmailAddress() string

	// GetVerificationCode parses and returns the verification code
	// (e.g., 6-digit number)
	GetVerificationCode() (string, error)

	// SetCodeSelector sets the selector for the verification code
	SetCodeSelector(selector string)

	// Close cleans up resources
	Close() error
}
