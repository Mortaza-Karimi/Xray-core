package conf_test

import (
	"testing"

	"github.com/Mortaza-Karimi/Xray-core/blob/main/common/serial"
	. "github.com/Mortaza-Karimi/Xray-core/blob/main/infra/conf"
	"github.com/Mortaza-Karimi/Xray-core/blob/main/proxy/blackhole"
)

func TestHTTPResponseJSON(t *testing.T) {
	creator := func() Buildable {
		return new(BlackholeConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"response": {
					"type": "http"
				}
			}`,
			Parser: loadJSON(creator),
			Output: &blackhole.Config{
				Response: serial.ToTypedMessage(&blackhole.HTTPResponse{}),
			},
		},
		{
			Input:  `{}`,
			Parser: loadJSON(creator),
			Output: &blackhole.Config{},
		},
	})
}
