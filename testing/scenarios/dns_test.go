package scenarios

import (
	"fmt"
	"testing"
	"time"

	"github.com/Mortaza-Karimi/Xray-core/blob/main/app/dns"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/app/proxyman"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/app/router"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/common"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/common/net"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/common/serial"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/core"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/proxy/blackhole"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/proxy/freedom"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/proxy/socks"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/testing/servers/tcp"
	xproxy "golang.org/x/net/proxy"
)

func TestResolveIP(t *testing.T) {
	tcpServer := tcp.Server{
		MsgProcessor: xor,
	}
	dest, err := tcpServer.Start()
	common.Must(err)
	defer tcpServer.Close()

	serverPort := tcp.PickPort()
	serverConfig := &core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&dns.Config{
				Hosts: map[string]*net.IPOrDomain{
					"google.com": net.NewIPOrDomain(dest.Address),
				},
			}),
			serial.ToTypedMessage(&router.Config{
				DomainStrategy: router.Config_IpIfNonMatch,
				Rule: []*router.RoutingRule{
					{
						Cidr: []*router.CIDR{
							{
								Ip:     []byte{127, 0, 0, 0},
								Prefix: 8,
							},
						},
						TargetTag: &router.RoutingRule_Tag{
							Tag: "direct",
						},
					},
				},
			}),
		},
		Inbound: []*core.InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortList: &net.PortList{Range: []*net.PortRange{net.SinglePortRange(serverPort)}},
					Listen:   net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&socks.ServerConfig{
					AuthType: socks.AuthType_NO_AUTH,
					Accounts: map[string]string{
						"Test Account": "Test Password",
					},
					Address:    net.NewIPOrDomain(net.LocalHostIP),
					UdpEnabled: false,
				}),
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&blackhole.Config{}),
			},
			{
				Tag: "direct",
				ProxySettings: serial.ToTypedMessage(&freedom.Config{
					DomainStrategy: freedom.Config_USE_IP,
				}),
			},
		},
	}

	servers, err := InitializeServerConfigs(serverConfig)
	common.Must(err)
	defer CloseAllServers(servers)

	{
		noAuthDialer, err := xproxy.SOCKS5("tcp", net.TCPDestination(net.LocalHostIP, serverPort).NetAddr(), nil, xproxy.Direct)
		common.Must(err)
		conn, err := noAuthDialer.Dial("tcp", fmt.Sprintf("google.com:%d", dest.Port))
		common.Must(err)
		defer conn.Close()

		if err := testTCPConn2(conn, 1024, time.Second*5)(); err != nil {
			t.Error(err)
		}
	}
}
