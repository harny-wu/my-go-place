package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		c <- struct{}{}
	}()
	for {
		select {
		case <-c:
			fmt.Println("done")
			return
		}
	}
}
