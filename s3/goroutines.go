package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cs := map[string]int{"cashier": 0} // not intended to be used with goroutines
	mu := &sync.Mutex{}                // always hold reference to a mutex

	for i := 0; i < 1000; i++ {
		go func(k int) {
			// mu.Unlock() // panic
			mu.Lock()
			// mu.Lock() // deadlock
			defer mu.Unlock()
			cs["cashier"] += 1
		}(i)
	}

	time.Sleep(time.Second)

	fmt.Println(cs)
}
