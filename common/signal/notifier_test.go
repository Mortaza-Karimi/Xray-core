package signal_test

import (
	"testing"

	. "github.com/Mortaza-Karimi/Xray-core/blob/main/common/signal"
)

func TestNotifierSignal(t *testing.T) {
	n := NewNotifier()

	w := n.Wait()
	n.Signal()

	select {
	case <-w:
	default:
		t.Fail()
	}
}
