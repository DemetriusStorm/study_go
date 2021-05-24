// 2. Объяснить причину постоянно разных значений в переменной «counter» при выполнении программы более чем на одном ядре.
// source: https://play.golang.org/p/r8BqUOaW-l
package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	wg.Add(1000)

	// redact
	m := new(sync.Mutex)

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				// redact
				m.Lock()
				counter++
				// redact
				m.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
