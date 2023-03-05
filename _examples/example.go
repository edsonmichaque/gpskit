package example

import (
	"context"
	"errors"

	"github.com/edsonmichaque/gpskit"
)

func main() {
	mux := gpskit.Mux{}

	codec := &gpskit.Codec{
		Enroller: gpskit.EnrollerFunc(func(ctx context.Context, b []byte) bool {
			return true
		}),
		Decoder: gpskit.DecoderFunc(func(ctx context.Context, b []byte) (*gpskit.Command, error) {
			return nil, errors.New("not implemented")
		}),
		Encoder: gpskit.EncoderFunc(func(ctx context.Context, c *gpskit.Command) ([]byte, error) {
			return nil, errors.New("not implemented")
		}),
	}

	mux.Register(codec)

}