package gpskit

import (
	"context"
)

type Codec struct {
	Enroller Enroller
	Encoder  Encoder
	Decoder  Decoder
	Handler  *Handler
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

type EnrollerFunc func(context.Context, []byte) bool

func (e EnrollerFunc) Enroll(ctx context.Context, data []byte) bool {
	return e(ctx, data)
}

type Enroller interface {
	Enroll(context.Context, []byte) bool
}
