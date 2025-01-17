//go:build !windows && !android
// +build !windows,!android

package domainsocket_test

import (
	"context"
	"runtime"
	"testing"

	"github.com/Mortaza-Karimi/Xray-core/blob/main/common"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/common/buf"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/common/net"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet"
	. "github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/domainsocket"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/stat"
)

func TestListen(t *testing.T) {
	ctx := context.Background()
	streamSettings := &internet.MemoryStreamConfig{
		ProtocolName: "domainsocket",
		ProtocolSettings: &Config{
			Path: "/tmp/ts3",
		},
	}
	listener, err := Listen(ctx, nil, net.Port(0), streamSettings, func(conn stat.Connection) {
		defer conn.Close()

		b := buf.New()
		defer b.Release()
		common.Must2(b.ReadFrom(conn))
		b.WriteString("Response")

		common.Must2(conn.Write(b.Bytes()))
	})
	common.Must(err)
	defer listener.Close()

	conn, err := Dial(ctx, net.Destination{}, streamSettings)
	common.Must(err)
	defer conn.Close()

	common.Must2(conn.Write([]byte("Request")))

	b := buf.New()
	defer b.Release()
	common.Must2(b.ReadFrom(conn))

	if b.String() != "RequestResponse" {
		t.Error("expected response as 'RequestResponse' but got ", b.String())
	}
}

func TestListenAbstract(t *testing.T) {
	if runtime.GOOS != "linux" {
		return
	}

	ctx := context.Background()
	streamSettings := &internet.MemoryStreamConfig{
		ProtocolName: "domainsocket",
		ProtocolSettings: &Config{
			Path:     "/tmp/ts3",
			Abstract: true,
		},
	}
	listener, err := Listen(ctx, nil, net.Port(0), streamSettings, func(conn stat.Connection) {
		defer conn.Close()

		b := buf.New()
		defer b.Release()
		common.Must2(b.ReadFrom(conn))
		b.WriteString("Response")

		common.Must2(conn.Write(b.Bytes()))
	})
	common.Must(err)
	defer listener.Close()

	conn, err := Dial(ctx, net.Destination{}, streamSettings)
	common.Must(err)
	defer conn.Close()

	common.Must2(conn.Write([]byte("Request")))

	b := buf.New()
	defer b.Release()
	common.Must2(b.ReadFrom(conn))

	if b.String() != "RequestResponse" {
		t.Error("expected response as 'RequestResponse' but got ", b.String())
	}
}
