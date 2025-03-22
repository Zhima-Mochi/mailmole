// Package mailmole provides a unified interface for handling email verification
// through temporary email services.
package mailmole

import (
	"github.com/Zhima-Mochi/mailmole/types"
)

// TunnelAgent is the interface that all email verification services implement
type TunnelAgent = types.TunnelAgent

// TunnelOptions holds options for creating a tunnel agent
type TunnelOptions = types.TunnelOptions

// BrowserOptions holds options for creating a browser
type BrowserOptions = types.BrowserOptions
