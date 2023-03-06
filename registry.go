package gpskit

import (
	"context"
	"errors"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func NewDispatcher(listener net.Listener) *Dispatcher {
	return &Dispatcher{
		entries:  make([]Entry, 0),
		listener: listener,
	}
}

type Dispatcher struct {
	mutext sync.Mutex

	entries  []Entry
	listener net.Listener
}

func (m *Dispatcher) Register(matcher Matcher, codec *Codec) {
	m.mutext.Lock()
	defer m.mutext.Unlock()

	if m.entries == nil {
		m.entries = make([]Entry, 0)
	}

	m.entries = append(m.entries, Entry{
		Codec:   codec,
		Matcher: matcher,
	})
}

func (m *Dispatcher) RegisterFunc(matcher MatcherFunc, codec *Codec) {
	m.Register(matcher, codec)
}

func (m *Dispatcher) find(req *Request) (*Codec, error) {
	for _, e := range m.entries {
		if e.Matcher.Match(req) {
		}
	}

	return nil, errors.New("no handler found")
}

func (m *Dispatcher) Dispatch() error {
	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGINT)

	var stop bool

	go func() {
		<-sig
		stop = true
	}()

	for {
		conn, err := m.listener.Accept()
		if err != nil {
			return err
		}

		if stop {
			break
		}

		req := &Request{
			reader: conn,
		}

		codec, err := m.find(req)
		if err != nil {
			conn.Close()
			continue
		}

		go func(conn net.Conn) {
			defer func() {
				conn.Close()
			}()

			for {
				buf := make([]byte, 512)
				if _, err := conn.Read(buf); err != nil {
					return
				}

				cmd, err := codec.Decoder.Decode(context.Background(), buf)
				if err != nil {
					return
				}

				if cmd.Position != nil {
					codec.Handler.PositionUpdater.UpdatePosition(context.Background(), cmd.Position)
				}
			}
		}(conn)
	}

	return errors.New("no codec matcher")
}

type Entry struct {
	Matcher Matcher
	Codec   *Codec
}

type Matcher interface {
	Match(req *Request) bool
}

type Request struct {
	reader io.Reader
}

type MatcherFunc func(req *Request) bool

func (m MatcherFunc) Match(req *Request) bool {
	return m(req)
}

func MatchAny(_ *Request) bool {
	return true
}
