package grpc

import (
	"net/url"
	"strings"

	"github.com/Mortaza-Karimi/Xray-core/blob/main/common"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet"
)

const protocolName = "grpc"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}

func (c *Config) getServiceName() string {
	// Normal old school config
	if !strings.HasPrefix(c.ServiceName, "/") {
		return url.PathEscape(c.ServiceName)
	}
	// Otherwise new custom paths
	rawServiceName := c.ServiceName[1:strings.LastIndex(c.ServiceName, "/")] // trim from first to last '/'
	serviceNameParts := strings.Split(rawServiceName, "/")
	for i := range serviceNameParts {
		serviceNameParts[i] = url.PathEscape(serviceNameParts[i])
	}
	return strings.Join(serviceNameParts, "/")
}

func (c *Config) getTunStreamName() string {
	// Normal old school config
	if !strings.HasPrefix(c.ServiceName, "/") {
		return "Tun"
	}
	// Otherwise new custom paths
	endingPath := c.ServiceName[strings.LastIndex(c.ServiceName, "/")+1:] // from the last '/' to end of string
	return url.PathEscape(strings.Split(endingPath, "|")[0])
}

func (c *Config) getTunMultiStreamName() string {
	// Normal old school config
	if !strings.HasPrefix(c.ServiceName, "/") {
		return "TunMulti"
	}
	// Otherwise new custom paths
	endingPath := c.ServiceName[strings.LastIndex(c.ServiceName, "/")+1:] // from the last '/' to end of string
	streamNames := strings.Split(endingPath, "|")
	if len(streamNames) == 1 { // client side. Service name is the full path to multi tun
		return url.PathEscape(streamNames[0])
	} else { // server side. The second part is the path to multi tun
		return url.PathEscape(streamNames[1])
	}
}
