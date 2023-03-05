package gpskit

import (
	"context"
)

type SpeedUpdaterFunc func(context.Context, *Speed) error

func (p SpeedUpdaterFunc) UpdateSpeed(ctx context.Context, pos *Speed) error {
	return p(ctx, pos)
}

type SpeedUpdater interface {
	UpdateSpeed(ctx context.Context, p *Speed) error
}
