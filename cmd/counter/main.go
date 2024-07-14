package main

import (
	"fmt"
	"time"
)

func count(item string) {
	for i := 0; i < 10; i++ {
		fmt.Println(item, i)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go count("apple")
	go count("orange")
	go count("grape")
	go count("melon")
	count("banana")
}
