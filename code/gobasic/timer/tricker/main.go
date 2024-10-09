package main

import (
	"fmt"
	"time"
)

func main() {
	tricker := time.NewTicker(2 * time.Second)

	for {
		select {
		case t := <-tricker.C:
			{
				fmt.Println("定时器触发: " + t.String())
			}
		}
	}
}
