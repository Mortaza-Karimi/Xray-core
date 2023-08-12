//go:build !linux && !freebsd && !darwin
// +build !linux,!freebsd,!darwin

package tcp

import (
	"github.com/Mortaza-Karimi/Xray-core/blob/main/common/net"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/stat"
)

func GetOriginalDestination(conn stat.Connection) (net.Destination, error) {
	return net.Destination{}, nil
}
