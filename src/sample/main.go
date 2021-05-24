package main

import (
	"fmt"
	"sync"

	//"runtime"
	"time"
)

func main() {
	m := new(sync.Mutex)
	start := time.Now()

	go func() {
		for i := 0; i <= 3; i++ {
			m.Lock()
			fmt.Println("First loop ^", i)
			m.Unlock()
		}
	}()

	go func() {
		for i := 0; i <= 6; i++ {
			m.Lock()
			fmt.Println("Second loop ^", i)
			m.Unlock()
		}
	}()

	elapsedTime := time.Since(start)

	fmt.Println("Total Time For Execution: " + elapsedTime.String())

	time.Sleep(time.Second)
}
