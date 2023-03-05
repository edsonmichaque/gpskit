package gpskit

import (
	"context"
)

type PositionUpdaterFunc func(context.Context, *Position) error

func (p PositionUpdaterFunc) UpdatePosition(ctx context.Context, pos *Position) error {
	return p(ctx, pos)
}

type PositionUpdater interface {
	UpdatePosition(ctx context.Context, p *Position) error
}
