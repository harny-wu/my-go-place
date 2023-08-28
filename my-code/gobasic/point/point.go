package main

import "fmt"

func main() {
	a := 1
	b := &a
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of a:%T , value of a:%v\n ", a, a)
	fmt.Printf("type of b:%T , value of b:%v\n ", b, b)
	fmt.Printf("type of c:%T , value of c:%v\n ", c, c)
}
