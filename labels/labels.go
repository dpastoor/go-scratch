package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	tick := time.NewTicker(time.Millisecond * 500)
	count := 0

	defer tick.Stop()

	go func() {
		time.Sleep(time.Second * 2)
		done <- struct{}{}
	}()

OUTER:

	for {
		select {
		case <-tick.C:
			count++
			fmt.Println(count)
		case <-done:
			fmt.Println("quitting")
			// if just do break
			// will only break out of the select statement, not the for loop, so need to use label
			break OUTER
		}
	}

	fmt.Println("done")
}
