package quic

import (
	"sync"

	"github.com/Mortaza-Karimi/Xray-core/blob/main/common/bytespool"
)

var pool *sync.Pool

func init() {
	pool = bytespool.GetPool(2048)
}

func getBuffer() []byte {
	return pool.Get().([]byte)
}

func putBuffer(p []byte) {
	pool.Put(p)
}
