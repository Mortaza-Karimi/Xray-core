package extension

import (
	"context"

	"github.com/Mortaza-Karimi/Xray-core/blob/main/features"
	"google.golang.org/protobuf/proto"
)

type Observatory interface {
	features.Feature

	GetObservation(ctx context.Context) (proto.Message, error)
}

func ObservatoryType() interface{} {
	return (*Observatory)(nil)
}
