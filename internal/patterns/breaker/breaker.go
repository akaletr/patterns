package breaker

import (
	"errors"
	"sync"
	"time"

	"cmd/main/main.go/internal/functions"
)

// Breaker при количестве ошибок больше errorCount подряд размыкает цепь на duration секунд
func Breaker(errorCount int, duration time.Duration, f functions.Function) functions.Function {
	mu := sync.RWMutex{}
	count := 0
	lastBreakTime := time.Now()

	return func() (interface{}, error) {
		mu.RLock()
		if count >= errorCount {
			if time.Since(lastBreakTime) < duration {
				mu.RUnlock()
				return 0, errors.New("error: function rejected")
			}
		}
		mu.RUnlock()

		res, err := f()
		if err != nil {
			mu.Lock()
			count++
			lastBreakTime = time.Now()
			mu.Unlock()
			return res, err
		}

		mu.Lock()
		count = 0
		mu.Unlock()

		_ = lastBreakTime
		return res, err
	}
}
