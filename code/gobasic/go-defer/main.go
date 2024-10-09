package main

import "fmt"

/*
*
方法中的参数a，有参函数中的参数v，会请求参数，直接把参数代入，所以输出的都是1。
i++变成2之后，3个defer语句以后声明先执行的顺序执行，无参函数中使用的a现在已经是2了，故输出2。
*/
func main() {
	i := 1
	defer fmt.Println("方法: ", i)
	defer func() { fmt.Println("无参函数: ", i) }()
	// 值传递
	defer func(v int) { fmt.Println("有参函数: ", v) }(i)
	i++
}
