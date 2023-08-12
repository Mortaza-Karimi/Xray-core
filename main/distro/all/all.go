package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/app/dispatcher"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/app/proxyman/inbound"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/app/commander"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/app/log/command"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/app/proxyman/command"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/app/stats/command"

	// Developer preview services
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/app/observatory/command"

	// Other optional features.
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/app/dns"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/app/dns/fakedns"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/app/log"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/app/metrics"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/app/policy"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/app/reverse"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/app/router"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/tagged/taggedimpl"

	// Developer preview features
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/app/observatory"

	// Inbound and outbound proxies.
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/proxy/blackhole"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/proxy/dns"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/proxy/dokodemo"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/proxy/freedom"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/proxy/http"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/proxy/loopback"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/proxy/shadowsocks"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/proxy/socks"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/proxy/trojan"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/proxy/vless/inbound"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/proxy/vless/outbound"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/proxy/vmess/inbound"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/proxy/vmess/outbound"
	// _ "github.com/Mortaza-Karimi/Xray-core/blob/main/proxy/wireguard"

	// Transports
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/domainsocket"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/grpc"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/http"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/kcp"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/quic"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/reality"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/tcp"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/tls"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/udp"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/websocket"

	// Transport headers
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/headers/http"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/headers/noop"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/headers/srtp"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/headers/tls"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/headers/utp"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/headers/wechat"
	// _ "github.com/Mortaza-Karimi/Xray-core/blob/main/transport/internet/headers/wireguard"

	// JSON & TOML & YAML
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/main/json"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/main/toml"
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/main/yaml"

	// Load config from file or http(s)
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/main/confloader/external"

	// Commands
	_ "github.com/Mortaza-Karimi/Xray-core/blob/main/main/commands/all"
)
