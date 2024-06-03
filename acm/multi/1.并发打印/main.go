package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

/*
* 交替打印 1 2 3 4
 */
func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	ch1 <- 1
	go echo(wg, ch1, ch2)
	go echo(wg, ch2, ch1)
	wg.Wait()
}

func echo(wg *sync.WaitGroup, curr chan int, next chan int) {
	for {
		select {
		case i := <-curr:
			if i <= 100 {
				fmt.Printf("gid : %v , print: %v \n", GetGoid(), i)
				next <- i + 1
			} else {
				next <- i + 1
				wg.Done()
			}
		}
	}
}

func GetGoid() int64 {
	var (
		buf [64]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine")
	)

	idField := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Errorf("can not get goroutine id: %v", err))
	}

	return int64(id)
}
