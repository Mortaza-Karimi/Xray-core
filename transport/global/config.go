package global

import (
	"github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet"
)

// Apply applies this Config.
func (c *Config) Apply() error {
	if c == nil {
		return nil
	}
	return internet.ApplyGlobalTransportSettings(c.TransportSettings)
}
