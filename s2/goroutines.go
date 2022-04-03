package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for x := 0; x < 10; x++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Printf("goroutine %v\n", x)
		}(x)
	}

	time.Sleep(time.Second)
}
