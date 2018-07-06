package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	ch := make(chan int, 1000)
	var wg sync.WaitGroup

	wg.Add(1)
	// goroutine1
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		wg.Done()
	}()

	// goroutine2
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println(ok, "a: ", a)
		}
	}()

	wg.Wait()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}