package main

import (
	"fmt"
	"time"
)

func fillData(c chan int) {
	i := 0
	for {
		if i > 10 {
			close(c)
			break
		}

		c <- i
		i++
		time.Sleep(time.Second * 1)
	}
}

func main() {
	canal := make(chan int)

	go fillData(canal)

	for x := range canal {
		fmt.Println(x)
	}
}
