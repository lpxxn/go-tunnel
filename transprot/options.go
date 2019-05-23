package transprot

import (
	"context"
	"time"
)

type DialOptions struct {
	Timeout time.Duration
	Context context.Context
}

type ListenOptions struct {
	Timeout time.Duration
	Context context.Context
}
