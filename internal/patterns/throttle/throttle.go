package throttle

import (
	"context"
	"errors"
	"sync"
	"time"

	"cmd/main/main.go/internal/functions"
)

func ReturnsError(callback functions.FunctionWithContext, max, refill uint, duration time.Duration) functions.FunctionWithContext {
	tokens := max
	once := sync.Once{}
	return func(ctx context.Context, i ...interface{}) (interface{}, error) {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		once.Do(func() {
			ticker := time.NewTicker(duration)
			go func() {
				defer ticker.Stop()
				for {
					select {
					case <-ctx.Done():
						return
					case <-ticker.C:
						t := tokens + refill
						if t > max {
							t = max
						}
						tokens = t
					}
				}
			}()
		})

		if tokens <= 0 {
			return nil, errors.New("error: too many calls")
		}

		tokens--
		return callback(ctx)
	}
}
