package tunnels

import (
	"errors"
	"log"
	"time"

	"github.com/Zhima-Mochi/mailmole/types"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

var (
	defaultTunnelOptions = &types.TunnelOptions{
		BrowserOptions: &types.BrowserOptions{
			Headless: true,
		},
	}
)

// Custom code extractor for different email formats
func customCodeExtractor(page *rod.Page, matcher string) (string, error) {
	if matcher == "" {
		matcher = `^\d{6}$`
	}

	iframeElement, err := page.Timeout(10 * time.Second).Element("iframe")
	if err != nil {
		return "", errors.New("failed to find iframe")
	}

	iframePage, err := iframeElement.Frame()
	if err != nil {
		return "", errors.New("failed to get iframe page")
	}

	defer func() {
		_ = page.Mouse.MoveTo(proto.Point{X: 10, Y: 10})
		_ = page.Mouse.Click(proto.InputMouseButtonLeft, 1)
	}()

	codeElement, err := iframePage.ElementR("*", matcher)
	if err != nil {
		return "", errors.New("failed to find verification code")
	}

	code, err := codeElement.Text()
	if err != nil {
		return "", errors.New("failed to get verification code")
	}

	return code, nil
}

func getBrowser(browserOptions *types.BrowserOptions) *rod.Browser {
	browser := rod.New()

	if browserOptions.URL != "" {
		browser.ControlURL(browserOptions.URL).MustConnect()

		return browser
	}

	if !browserOptions.Headless {
		launcher := launcher.New().Headless(false)
		if err := browser.ControlURL(launcher.MustLaunch()).Connect(); err != nil {
			log.Fatalf("Failed to connect to browser: %v", err)
		}
	}

	return browser
}
