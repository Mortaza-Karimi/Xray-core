package drain

import "io"

//go:generate go run github.com/Mortaza-Karimi/Xray-core/blob/main/common/errors/errorgen

type Drainer interface {
	AcknowledgeReceive(size int)
	Drain(reader io.Reader) error
}
