package retry

import (
	"context"
	"time"

	"cmd/main/main.go/internal/functions"
)

func Retry(callback functions.FunctionWithContext, retryCount int, delay time.Duration) functions.FunctionWithContext {
	return func(ctx context.Context, options ...interface{}) (interface{}, error) {
		for r := 0; ; r++ {
			result, err := callback(ctx)
			if err == nil || r >= retryCount {
				return result, err
			}
			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}
	}
}
