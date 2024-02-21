package main

import "fmt"

func main() {
	var sli_1 []int //nil 切片
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_1), cap(sli_1), sli_1)

	var sli_2 = []int{} //空切片
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_1), cap(sli_2), sli_2)

	var sli_3 = []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_3), cap(sli_3), sli_3)

	sli_4 := []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_4), cap(sli_4), sli_4)

	var sli_5 []int = make([]int, 5, 8)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_5), cap(sli_5), sli_5)

	sli_6 := make([]int, 5, 9)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_6), cap(sli_6), sli_6)

	sli_7 := make([]int, 0, 9)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_7), cap(sli_7), sli_7)

	sli_8 := make([]int, 0)
	sli_8 = append(sli_8, 1)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_8), cap(sli_8), sli_8)

}
