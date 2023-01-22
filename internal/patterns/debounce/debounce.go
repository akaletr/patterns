package debounce

import (
	"sync"
	"time"

	"cmd/main/main.go/internal/functions"
)

// First ограничивает вызов функции callback, вызывает первую из серии в интервале duration
func First(duration time.Duration, callback functions.Function) functions.Function {
	var lastCall time.Time
	mu := sync.RWMutex{}

	_ = lastCall
	return func(options ...interface{}) (interface{}, error) {
		mu.Lock()
		lastCall = time.Now()
		mu.Unlock()

		res, err := callback()

		if err != nil {

		}

		return res, err
	}
}

// Last ограничивает вызов функции callback, вызывает последнюю из серии после интервала duration
func Last(duration time.Duration, callback functions.Function) functions.Function {

	return func(...interface{}) (interface{}, error) {
		return callback()
	}
}
