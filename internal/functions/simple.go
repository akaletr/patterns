package functions

import (
	"errors"
	"math/rand"
	"time"
)

type Function func() (interface{}, error)

func NewSimpleFunc(timeout time.Duration, errorPercent int) Function {
	return func() (interface{}, error) {
		time.Sleep(timeout)
		result := rand.Intn(100)

		if result < errorPercent {
			return 0, errors.New("error: something went wrong")
		}

		return result, nil
	}
}
