package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(workerId int, msg chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for res := range msg {
		fmt.Println("Worker", workerId, ":", res)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	msg := make(chan string)
	wg := sync.WaitGroup{}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(i, msg, &wg)
	}

	for i := 0; i < 100; i++ {
		msg <- fmt.Sprintf("Message %d", i)
	}

	fmt.Println("Closing channel")
	close(msg)

	wg.Wait()
}
