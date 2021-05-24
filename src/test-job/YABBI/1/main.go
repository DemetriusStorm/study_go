// 1. Измените программу так, чтобы цифры от 1 до 9 печатались в консоль по порядку. 
// source: https://play.golang.org/p/lYQo-iGx2E
// Если я правильно понял, то задача на каналы.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	// redact
	ch1 := make(chan bool, 1)
	ch2 := make(chan bool, 1)
	ch3 := make(chan bool, 1)

    defer close(ch1)
    defer close(ch2)
    defer close(ch3)

    ch1 <- true

	go func() {
		for _, value := range []int{1, 4, 7} {
			// redact
			<- ch1
			fmt.Println(value)
			// redact
			ch2 <- true
		}

		wg.Done()
	}()

	go func() {
		for _, value := range []int{2, 5, 8} {
			// redact
			<- ch2
			fmt.Println(value)
			// redact
			ch3 <- true
		}

		wg.Done()
	}()

	go func() {
		for _, value := range []int{3, 6, 9} {
			// redact
			<- ch3
			fmt.Println(value)
			// redact
			ch1 <- true
		}

		wg.Done()
	}()

	wg.Wait()
}