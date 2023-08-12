package udp

import (
	"github.com/Mortaza-Karimi/Xray-core/blob/main/common"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
