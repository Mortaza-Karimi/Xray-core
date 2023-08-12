package command_test

import (
	"context"
	"testing"

	"github.com/Mortaza-Karimi/Xray-core/blob/main/app/dispatcher"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/app/log"
	. "github.com/Mortaza-Karimi/Xray-core/blob/main/app/log/command"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/app/proxyman"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/app/proxyman/inbound"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/app/proxyman/outbound"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/common"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/common/serial"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/core"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
