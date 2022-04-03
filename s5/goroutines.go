package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cs := map[string]int{"cashier": 0}
	mu := &sync.RWMutex{}
	// mu.RLock()
	// mu.RUnlock()

	for i := 0; i < 1000; i++ {
		go func(k int) {
			mu.Lock()
			defer mu.Unlock()
			cs["cashier"] += 1
		}(i)
	}

	time.Sleep(time.Second)

	fmt.Println(cs)
}
