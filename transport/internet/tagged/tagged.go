package tagged

import (
	"context"

	"github.com/Mortaza-Karimi/Xray-core/blob/main/common/net"
)

type DialFunc func(ctx context.Context, dest net.Destination, tag string) (net.Conn, error)

var Dialer DialFunc
