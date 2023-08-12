package net_test

import (
	"testing"

	. "github.com/Mortaza-Karimi/Xray-core/blob/main/common/net"
)

func TestPortRangeContains(t *testing.T) {
	portRange := &PortRange{
		From: 53,
		To:   53,
	}

	if !portRange.Contains(Port(53)) {
		t.Error("expected port range containing 53, but actually not")
	}
}
