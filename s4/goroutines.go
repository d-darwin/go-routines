package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cs := map[string]int{"cashier1": 0, "cashier2": 0} // not intended to be used with goroutines
	mu := &sync.Mutex{}

	for i := 0; i < 100; i++ {
		go func(k int) {
			mu.Lock()
			defer mu.Unlock()
			cs["cashier1"] += 1
		}(i)
	}

	for i := 0; i < 100; i++ {
		go func(k int) {
			mu.Lock()
			defer mu.Unlock()
			cs["cashier2"] += 1
		}(i)
	}

	time.Sleep(time.Second)

	fmt.Println(cs)
}
