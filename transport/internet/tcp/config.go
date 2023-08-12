package tcp

import (
	"github.com/Mortaza-Karimi/xray-core/common"
	"github.com/Mortaza-Karimi/xray-core/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
