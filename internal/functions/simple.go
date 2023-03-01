package functions

import (
	"context"
	"errors"
	"fmt"
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
	count := 1
	return func(ctx context.Context, i ...interface{}) (interface{}, error) {
		fmt.Println("NewSimpleFuncWithContext started, #", count)
		count++
		time.Sleep(timeout)
		rand.Seed(time.Now().UnixNano())
		result := rand.Intn(100)

		if result < errorPercent {
			fmt.Println("NewSimpleFuncWithContext error")
			return 0, errors.New("error: something went wrong")
		}
		fmt.Println("NewSimpleFuncWithContext done")
		return result, nil
	}
}
