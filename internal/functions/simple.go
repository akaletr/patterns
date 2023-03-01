package functions

import (
	"context"
	"errors"
	"math/rand"
	"time"
)

type Function func(...interface{}) (interface{}, error)

func NewSimpleFunc(timeout time.Duration, errorPercent int) Function {
	return func(...interface{}) (interface{}, error) {
		time.Sleep(timeout)
		result := rand.Intn(100)

		if result < errorPercent {
			return 0, errors.New("error: something went wrong")
		}

		return result, nil
	}
}

type FunctionWithContext func(context.Context, ...interface{}) (interface{}, error)

func NewSimpleFuncWithContext(timeout time.Duration, errorPercent int) FunctionWithContext {
	return func(ctx context.Context, i ...interface{}) (interface{}, error) {
		time.Sleep(timeout)
		result := rand.Intn(100)

		if result < errorPercent {
			return 0, errors.New("error: something went wrong")
		}

		return result, nil
	}
}
