package main

import "fmt"

func main() {
	arr := [2]int{1, 2}
	fmt.Println("====== arr 函数外 ====== ")
	fmt.Println(arr)
	fmt.Println(&arr, &arr[0], &arr[1])
	add1(arr)

	slic := []int{1, 2}
	fmt.Println("====== slice 函数外 ====== ")
	fmt.Println(slic)
	fmt.Println(&slic, &slic[0], &slic[1])
	add2(slic)

}

func add1(x [2]int) {
	fmt.Println("====== arr 函数内 ====== ")
	fmt.Println(x)
	fmt.Println(&x, &x[0], &x[1])
}

func add2(x []int) {
	fmt.Println("====== slice 函数内 ====== ")
	fmt.Println(x)
	fmt.Println(&x, &x[0], &x[1])
}
