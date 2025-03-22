package tunnels

import (
	"errors"
	"strings"
	"time"

	"github.com/Zhima-Mochi/mailmole/types"
	"github.com/go-rod/rod"
)

// YopmailTunnel implements the TunnelAgent interface for yopmail.com
type YopmailTunnel struct {
	email         string
	page          *rod.Page
	browser       *rod.Browser
	codeSelector  string
	tunnelOptions *types.TunnelOptions
}

// NewYopmailTunnel creates a new YopmailTunnel
func NewYopmailTunnel(options *types.TunnelOptions) types.TunnelAgent {
	if options == nil {
		options = defaultTunnelOptions
	}

	return &YopmailTunnel{
		codeSelector:  `\b\d{6}\b`,
		tunnelOptions: options,
	}
}

// Init initializes the YopmailTunnel
func (v *YopmailTunnel) Init() error {
	// Launch a new browser
	v.browser = getBrowser(v.tunnelOptions.BrowserOptions)

	// Create a new page
	page := v.browser.MustPage("https://yopmail.com/")
	v.page = page

	// Wait for page to load
	page.MustWaitLoad()

	// If no email was provided, generate a random one
	if v.tunnelOptions.Email == "" {
		// Click on random email button
		randomButton, err := page.Element("a.genrnd")
		if err != nil {
			return errors.New("failed to find random button")
		}
		randomButton.MustClick()
		time.Sleep(time.Second)

		// Get the generated email address
		emailInput, err := page.Element("#login")
		if err != nil {
			return errors.New("failed to find email input")
		}

		v.email = emailInput.MustProperty("value").String() + "@yopmail.com"
	} else {
		// Enter the provided email
		emailInput, err := page.Element("#login")
		if err != nil {
			return errors.New("failed to find email input")
		}

		// Extract username if a full email was provided
		username := v.tunnelOptions.Email
		if strings.Contains(username, "@") {
			username = strings.Split(username, "@")[0]
		}

		emailInput.MustInput(username)

		// Click the check button
		checkButton, err := page.Element("button.material-icons-outlined.f36")
		if err != nil {
			return errors.New("failed to find check button")
		}
		checkButton.MustClick()

		// Ensure the email is complete with domain
		if !strings.Contains(v.email, "@") {
			v.email = username + "@yopmail.com"
		}
	}

	time.Sleep(2 * time.Second)
	return nil
}

// EmailAddress returns the current email address
func (v *YopmailTunnel) EmailAddress() string {
	return v.email
}

// SetCodeSelector sets the selector for the verification code
func (v *YopmailTunnel) SetCodeSelector(selector string) {
	v.codeSelector = selector
}

// GetVerificationCode extracts the verification code from the email
func (v *YopmailTunnel) GetVerificationCode() (string, error) {
	// Refresh inbox
	if err := v.refreshInbox(); err != nil {
		return "", err
	}

	// Extract verification code
	return customCodeExtractor(v.page, v.codeSelector)
}

// Close cleans up resources
func (v *YopmailTunnel) Close() error {
	if v.browser != nil {
		v.browser.MustClose()
	}
	return nil
}

// refreshInbox refreshes the inbox
func (v *YopmailTunnel) refreshInbox() error {
	// Refresh inbox
	refreshButton, err := v.page.Element("#refresh")
	if err != nil {
		return errors.New("failed to find refresh button")
	}
	refreshButton.MustClick()
	time.Sleep(2 * time.Second)

	return nil
}
