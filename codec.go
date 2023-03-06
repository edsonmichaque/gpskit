package gpskit

import (
	"context"
)

type Codec struct {
	Encoder Encoder
	Decoder Decoder
	Handler *Handler
}

type EncoderFunc func(context.Context, *Command) ([]byte, error)

func (e EncoderFunc) Encode(ctx context.Context, cmd *Command) ([]byte, error) {
	return e(ctx, cmd)
}

type Encoder interface {
	Encode(context.Context, *Command) ([]byte, error)
}

type DecoderFunc func(context.Context, []byte) (*Command, error)

func (d DecoderFunc) Decode(ctx context.Context, data []byte) (*Command, error) {
	return d(ctx, data)
}

type Decoder interface {
	Decode(context.Context, []byte) (*Command, error)
}
