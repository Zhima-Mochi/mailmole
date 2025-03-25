package tunnels

import (
	"errors"
	"strings"
	"time"

	"github.com/Zhima-Mochi/mailmole/types"
	"github.com/go-rod/rod"
)

// SmailProTunnel implements the TunnelAgent interface for smailpro.com
type SmailProTunnel struct {
	email         string
	page          *rod.Page
	browser       *rod.Browser
	codeMatcher   string
	tunnelOptions *types.TunnelOptions
}

// NewSmailProTunnel creates a new SmailProTunnel
func NewSmailProTunnel(options *types.TunnelOptions) types.TunnelAgent {
	if options == nil {
		options = defaultTunnelOptions
	}

	return &SmailProTunnel{
		tunnelOptions: options,
	}
}

// Init initializes the SmailProTunnel
func (v *SmailProTunnel) Init() error {
	// Launch a new browser
	v.browser = getBrowser(v.tunnelOptions.BrowserOptions)

	pages := v.browser.MustPages()
	for _, p := range pages {
		if strings.Contains(p.MustHTML(), "https://smailpro.com/") {
			v.page = p
			break
		}
	}

	// Create a new page
	if v.page == nil {
		page := v.browser.MustPage("https://smailpro.com/")
		v.page = page
	}

	// Wait for page to load
	v.page.MustWaitLoad()

	return nil
}

func (v *SmailProTunnel) RenewEmail() error {
	// Click create email button
	createButton, err := v.page.Element(`button[title="Create temporary email"]`)
	if err != nil {
		return errors.New("failed to find create email button")
	}
	createButton.MustClick()
	time.Sleep(time.Second)

	// Click modal create button
	modalCreateButton, err := v.page.ElementR("button", "Create")
	if err != nil {
		return errors.New("failed to find modal create button")
	}
	modalCreateButton.MustClick()
	time.Sleep(2 * time.Second)

	return nil
}

// EmailAddress returns the current email address
func (v *SmailProTunnel) EmailAddress() (string, error) {
	// Get the email address
	emailElement, err := v.page.Element(`div[class="text-base sm:text-lg md:text-xl text-gray-700"]`)
	if err != nil {
		return "", errors.New("failed to find email element")
	}

	v.email = emailElement.MustText()
	if !strings.Contains(v.email, "@") {
		return "", errors.New("invalid email address")
	}

	return v.email, nil
}

// SetCodeExtractor sets a custom function to extract verification codes
func (v *SmailProTunnel) SetCodeMatcher(selector string) {
	v.codeMatcher = selector
}

// GetVerificationCode extracts the verification code from the email
func (v *SmailProTunnel) GetVerificationCode() (string, error) {
	// Refresh inbox
	if err := v.refreshInbox(); err != nil {
		return "", err
	}

	// Click on the first message if found
	found := v.page.MustEval(`() => {
		const elements = document.querySelectorAll("div.cursor-pointer");
		for (const el of elements) {
			const clickAttr = el.getAttribute("@click");
			const xClickAttr = el.getAttribute("x-on:click");
			if (clickAttr === "message(getTemporaryEmailAddress(), mes)" || 
				xClickAttr === "message(getTemporaryEmailAddress(), mes)") {
				el.click();
				return true;
			}
		}
		return false;
	}`).Bool()

	if !found {
		return "", errors.New("no message found to click")
	}

	// Wait a bit for the message to load
	time.Sleep(1 * time.Second)

	// Extract verification code after clicking
	return customCodeExtractor(v.page, v.codeMatcher)
}

// Close cleans up resources
func (v *SmailProTunnel) Close() error {
	if v.browser != nil {
		v.browser.MustClose()
	}
	return nil
}

// refreshInbox refreshes the inbox
func (v *SmailProTunnel) refreshInbox() error {
	// Refresh inbox
	refreshButton, err := v.page.Element(`button[id="refresh"]`)
	if err != nil {
		return errors.New("failed to find refresh button")
	}
	refreshButton.MustClick()
	time.Sleep(2 * time.Second)

	return nil
}
