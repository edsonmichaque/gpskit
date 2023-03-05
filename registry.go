package gpskit

import (
	"context"
	"errors"
	"sync"
)

func NewRegistry() *Registry {
	return &Registry{
		entries: make([]*Codec, 0),
	}
}

type Registry struct {
	mutext sync.Mutex

	entries []*Codec
}

func (m *Registry) Register(codec *Codec) {
	m.mutext.Lock()
	defer m.mutext.Unlock()

	if m.entries == nil {
		m.entries = make([]*Codec, 0)
	}

	m.entries = append(m.entries, codec)
}

func (m *Registry) Discovery(ctx context.Context, data []byte) (*Codec, error) {
	for _, v := range m.entries {
		if v.Matcher.Match(ctx, data) {
			return v, nil
		}
	}

	return nil, errors.New("no codec found")
}
