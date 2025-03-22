package mailmole

import (
	"fmt"

	"github.com/Zhima-Mochi/mailmole/tunnels"
)

// TunnelType represents the type of email tunnel agent to create
type TunnelType string

const (
	// SmailPro tunnel type
	SmailProTunnel TunnelType = "smailpro"
	// Yopmail tunnel type
	YopmailTunnel TunnelType = "yopmail"
)

// CreateTunnel is a factory function that creates and returns a TunnelAgent implementation
// based on the specified type and options
func CreateTunnel(tunnelType TunnelType, options *TunnelOptions) (TunnelAgent, error) {
	switch tunnelType {
	case SmailProTunnel:
		return tunnels.NewSmailProTunnel(options), nil
	case YopmailTunnel:
		return tunnels.NewYopmailTunnel(options), nil
	}
	return nil, fmt.Errorf("unsupported tunnel type: %s", tunnelType)
}

// RegisterTunnelType could be added in the future to allow registration of custom tunnel types
