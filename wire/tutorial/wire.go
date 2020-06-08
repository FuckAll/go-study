//+build wireinject

package main

import (
	"context"
	"github.com/google/wire"
)

func InitializeEvent(phrase string, ctx context.Context, good string) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}
