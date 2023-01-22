package main

import (
	"cmd/main/main.go/internal/functions"
	"fmt"
	"time"

	"cmd/main/main.go/internal/patterns/breaker"
)

func main() {
	f := breaker.Breaker(3, time.Millisecond*200, functions.NewSimpleFunc(time.Millisecond*100, 50))

	for i := 0; i < 1000; i++ {
		fmt.Println(f())
	}

	fmt.Println("end of function")
}
