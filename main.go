package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	counter := 0
 
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			counter++
			defer mu.Unlock()
		}()
	}

 	time.Sleep(1 * time.Second)
	fmt.Println(counter)
}