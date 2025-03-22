package tunnels

import (
	"errors"
	"time"

	"github.com/Zhima-Mochi/mailmole/types"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

// Custom code extractor for different email formats
func customCodeExtractor(page *rod.Page, selector string) (string, error) {
	if selector == "" {
		selector = `\b\d{6}\b`
	}

	iframeElement, err := page.Timeout(10 * time.Second).Element("iframe")
	if err != nil {
		return "", errors.New("failed to find iframe")
	}

	iframePage, err := iframeElement.Frame()
	if err != nil {
		return "", errors.New("failed to get iframe page")
	}

	codeElement, err := iframePage.ElementR("*", selector)
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
	if browserOptions == nil {
		browserOptions = &types.BrowserOptions{
			Headless: true,
		}
	}

	if browserOptions.URL == "" {
		browserOptions.URL = "https://www.google.com"
	}

	var browser *rod.Browser
	if browserOptions.Headless {
		browser = rod.New().ControlURL(browserOptions.URL).MustConnect()
	} else {
		url := launcher.New().Headless(false).MustLaunch()
		browser = rod.New().ControlURL(url).MustConnect()
	}
	return browser
}
