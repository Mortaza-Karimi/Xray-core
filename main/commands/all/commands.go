package all

import (
	"github.com/Mortaza-Karimi/xray-core/main/commands/all/api"
	"github.com/Mortaza-Karimi/xray-core/main/commands/all/tls"
	"github.com/Mortaza-Karimi/xray-core/main/commands/base"
)

// go:generate go run github.com/Mortaza-Karimi/xray-core/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		// cmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
	)
}
