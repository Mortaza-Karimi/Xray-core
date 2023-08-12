package all

import (
	"github.com/Mortaza-Karimi/Xray-core/blob/main/main/commands/all/api"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/main/commands/all/tls"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/main/commands/base"
)

// go:generate go run github.com/Mortaza-Karimi/Xray-core/blob/main/common/errors/errorgen

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
