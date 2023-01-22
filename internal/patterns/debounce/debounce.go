package debounce

import (
	"sync"
	"time"

	"cmd/main/main.go/internal/functions"
)

// First ограничивает вызов функции callback, вызывает первую из серии в интервале duration
func First(duration time.Duration, callback functions.Function) functions.Function {
	var (
		lastCall time.Time
		result   interface{}
		err      error
	)

	mu := sync.RWMutex{}

	return func(options ...interface{}) (interface{}, error) {
		mu.Lock()
		defer func() {
			lastCall = time.Now()
			mu.Unlock()
		}()

		if time.Now().Before(lastCall.Add(duration)) {
			return result, nil
		}

		result, err = callback()
		return result, err
	}
}

// Last ограничивает вызов функции callback, вызывает последнюю из серии после интервала duration
func Last(duration time.Duration, callback functions.Function) functions.Function {

	return func(...interface{}) (interface{}, error) {
		return callback()
	}
}
