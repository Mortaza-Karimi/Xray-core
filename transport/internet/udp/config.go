package udp

import (
	"github.com/Mortaza-Karimi/xray-core/common"
	"github.com/Mortaza-Karimi/xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
