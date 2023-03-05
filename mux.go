package gpskit

import (
	"context"
	"errors"
	"sync"
)

type Mux struct {
	mutext sync.Mutex

	entries []*Codec
}

func (m *Mux) Register(codec *Codec) {
	m.mutext.Lock()
	defer m.mutext.Unlock()

	if m.entries == nil {
		m.entries = make([]*Codec, 0)
	}

	m.entries = append(m.entries, codec)
}

func (m *Mux) Discovery(ctx context.Context, data []byte) (*Codec, error) {
	for _, v := range m.entries {
		if v.Enroller.Enroll(ctx, data) {
			return v, nil
		}
	}

	return nil, errors.New("no codec found")
}
