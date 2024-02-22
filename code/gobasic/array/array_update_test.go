package main

import (
	"fmt"
	"testing"
)

func TestArrayUpdate(t *testing.T) {
	arr := [1]int{1}
	fmt.Printf("arr %p \n", &arr)
	update1(arr)
	fmt.Println(arr)
	update2(&arr)
	fmt.Println(arr)
}
func update1(arr [1]int) {
	fmt.Printf("arr1 %p \n", &arr)
	arr[0] = 100
}

func update2(arr *[1]int) {
	fmt.Printf("arr2 %p  \n", arr)
	arr[0] = 100
}
