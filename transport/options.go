package transport

import (
	"context"
	"time"
)

var defaultTimeout = time.Second * 15

type DialOptions struct {
	Timeout time.Duration
	Context context.Context
}

type ListenOptions struct {
	Timeout time.Duration
	Context context.Context
}
