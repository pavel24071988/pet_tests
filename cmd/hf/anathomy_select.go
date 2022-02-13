package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func main() {
	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string, 2)
	chan2 := make(chan string, 2)

	chan1 <- "Value 1"
	chan1 <- "Value 2"
	chan2 <- "Value 1"
	chan2 <- "Value 2"

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Respnse from service 2", res, time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
}
