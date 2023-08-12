package conf

import (
	"github.com/Mortaza-Karimi/Xray-core/blob/main/proxy/loopback"
	"google.golang.org/protobuf/proto"
)

type LoopbackConfig struct {
	InboundTag string `json:"inboundTag"`
}

func (l LoopbackConfig) Build() (proto.Message, error) {
	return &loopback.Config{InboundTag: l.InboundTag}, nil
}
