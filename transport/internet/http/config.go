package http

import (
	"github.com/Mortaza-Karimi/Xray-core/blob/main/common"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/common/dice"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet"
)

const protocolName = "http"

func (c *Config) getHosts() []string {
	if len(c.Host) == 0 {
		return []string{"www.example.com"}
	}
	return c.Host
}

func (c *Config) isValidHost(host string) bool {
	hosts := c.getHosts()
	for _, h := range hosts {
		if h == host {
			return true
		}
	}
	return false
}

func (c *Config) getRandomHost() string {
	hosts := c.getHosts()
	return hosts[dice.Roll(len(hosts))]
}

func (c *Config) getNormalizedPath() string {
	if c.Path == "" {
		return "/"
	}
	if c.Path[0] != '/' {
		return "/" + c.Path
	}
	return c.Path
}

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
