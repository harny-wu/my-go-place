package main

import "fmt"

func quickSort(arr []int, low, high int) {
	if low < high {
		// 基准数
		pivot := arr[low]
		i, j := low, high
		for i < j {
			for arr[j] >= pivot && i < j {
				j--
			}
			for arr[i] <= pivot && i < j {
				i++
			}
			if i < j {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		arr[low] = arr[i]
		arr[i] = pivot
		quickSort(arr, low, i-1)
		quickSort(arr, i+1, high)
	}

}

func main() {
	array := []int{2, 34, 12, 3, 4, 42, 312}
	quickSort(array, 0, len(array)-1)
	fmt.Println(array)
}
