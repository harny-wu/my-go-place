package main

import "fmt"

func main() {
	a := 1
	fmt.Println("函数外变量地址 : ", &a, a)
	add1(a)
	add2(&a)
}

func add1(x int) {
	fmt.Println("函数内变量地址 : ", &x, x)

}
func add2(x *int) {
	fmt.Println("函数内变量地址 : ", &x, x, *x)
}

func add3() {

}
