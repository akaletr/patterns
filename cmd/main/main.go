package main

import (
	"cmd/main/main.go/internal/patterns/debounce"
	"fmt"
	"time"

	"cmd/main/main.go/internal/functions"
)

func main() {
	f := debounce.First(time.Millisecond*200, functions.NewSimpleFunc(time.Millisecond*100, 50))

	for i := 0; i < 1000; i++ {
		fmt.Println(f())
	}

	fmt.Println("end of function")
}
