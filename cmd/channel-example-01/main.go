package main

import "fmt"

func main() {
	canal := make(chan string)

	go func() {
		canal <- "apple"
	}()

	result := <-canal

	fmt.Println(result)
}
