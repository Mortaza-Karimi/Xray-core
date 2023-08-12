//go:build !windows && !wasm
// +build !windows,!wasm

package domainsocket

import (
	"context"

	"github.com/Mortaza-Karimi/Xray-core/blob/main/common"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/common/net"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/reality"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/stat"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/tls"
)

func Dial(ctx context.Context, dest net.Destination, streamSettings *internet.MemoryStreamConfig) (stat.Connection, error) {
	settings := streamSettings.ProtocolSettings.(*Config)
	addr, err := settings.GetUnixAddr()
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUnix("unix", nil, addr)
	if err != nil {
		return nil, newError("failed to dial unix: ", settings.Path).Base(err).AtWarning()
	}

	if config := tls.ConfigFromStreamSettings(streamSettings); config != nil {
		return tls.Client(conn, config.GetTLSConfig(tls.WithDestination(dest))), nil
	} else if config := reality.ConfigFromStreamSettings(streamSettings); config != nil {
		return reality.UClient(conn, config, ctx, dest)
	}

	return conn, nil
}

func init() {
	common.Must(internet.RegisterTransportDialer(protocolName, Dial))
}
