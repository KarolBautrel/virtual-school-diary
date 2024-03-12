package globalutils

import (
	"context"
	"time"
)

func NewTimeoutContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)

}
