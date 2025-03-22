package mailmole

import (
	"fmt"

	"github.com/Zhima-Mochi/mailmole/tunnels"
	"github.com/Zhima-Mochi/mailmole/types"
)

// TunnelType represents the type of email tunnel agent to create
type TunnelType string

const (
	// SmailPro tunnel type
	SmailPro TunnelType = "smailpro"
	// Yopmail tunnel type
	Yopmail TunnelType = "yopmail"
)

// CreateTunnel is a factory function that creates and returns a TunnelAgent implementation
// based on the specified type and options
func CreateTunnel(tunnelType TunnelType, options TunnelOptions) (types.TunnelAgent, error) {
	switch tunnelType {
	case SmailPro:
		return tunnels.NewSmailProTunnel(options), nil
	case Yopmail:
		return tunnels.NewYopmailTunnel(options), nil
	}
	return nil, fmt.Errorf("unsupported tunnel type: %s", tunnelType)
}

// RegisterTunnelType could be added in the future to allow registration of custom tunnel types
