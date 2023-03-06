package example

import (
	"context"
	"errors"
	"net"

	"github.com/edsonmichaque/gpskit"
)

func main() {
	ln, err := net.Listen("tcp", ":12345")
	if err != nil {
		panic(err)
	}

	mux := gpskit.NewDispatcher(ln)

	codec := &gpskit.Codec{
		Decoder: gpskit.DecoderFunc(func(ctx context.Context, b []byte) (*gpskit.Command, error) {
			return nil, errors.New("not implemented")
		}),
		Encoder: gpskit.EncoderFunc(func(ctx context.Context, c *gpskit.Command) ([]byte, error) {
			return nil, errors.New("not implemented")
		}),
	}

	mux.RegisterFunc(gpskit.MatchAny, codec)
	mux.Dispatch()
}
