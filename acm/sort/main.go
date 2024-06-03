package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 1, 3, 5, 4}
	fmt.Println(arr)
	mergeSort(arr, 0, len(arr)-1, make([]int, len(arr)))
	fmt.Println(arr)

}

/*
1. 取第一个节点为哨兵x
2. 左右指针 分别等于 i = l 和 j = r
3. 从右边开始，找到第一个小于x 的值的索引 j
4. 哨兵处左指针 i 处被覆盖为 arr[j]的值, arr[i] = arr[j]
5. 左指针i向右找，找到大于x的值, arr[j] = arr[i]
6. 当 i = j 时，这个 索引肯定是被覆盖过的或者是哨兵的原位置r , 将其覆盖为 值 x， 完成一层循环
*/
func quickSort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	x := arr[l]
	i := l
	j := r
	for i < j {
		for i < j && arr[j] > x {
			j--
		}
		if i < j {
			arr[i] = arr[j]
		}

		for i < j && arr[i] < x {
			i++
		}
		if i < j {
			arr[j] = arr[i]
		}
	}
	arr[i] = x
	quickSort(arr, l, i-1)
	quickSort(arr, i+1, r)
}

/*
1. 取第一个节点为哨兵x
2. 左右指针 分别 i = l 和 j = r+1
3. 从第一个开始
*/
func quickSortV2(arr []int, l int, r int) {
	if l >= r {
		return
	}
	x := arr[l]
	i := l
	j := r
	for i < j {
		for i < j {
			for i < j && arr[i] < x {
				i++
			}
			for i < j && arr[j] > x {
				j--
			}
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}
	quickSort(arr, l, i-1)
	quickSort(arr, i, r)
}

func mergeSort(arr []int, l, r int, temp []int) {
	if l < r {
		mid := (l + r) / 2
		mergeSort(arr, l, mid, temp)
		mergeSort(arr, mid+1, r, temp)
		merge(arr, l, mid, r, temp)
	}
}

func merge(arr []int, l, mid, r int, temp []int) {
	i := l
	j := mid + 1
	k := 0
	for i <= mid && j <= r {
		if arr[i] < arr[j] {
			temp[k] = arr[i]
			k++
			i++
		} else {
			temp[k] = arr[j]
			k++
			j++
		}
	}
	for i <= mid {
		temp[k] = arr[i]
		k++
		i++
	}
	for j <= r {
		temp[k] = arr[j]
		k++
		j++
	}

	k = 0
	for l <= r {
		arr[l] = temp[k]
		l++
		k++
	}
}
