package main

import (
	"fmt"
	"sort"
)

/*
* 合并区间 : https://leetcode.cn/problems/merge-intervals/
 */

func main() {
	var intervals = [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	fmt.Println(intervals, &intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals, &intervals)

	result := merge(intervals)
	fmt.Println(result, &result)
}

func merge(s [][]int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < len(s); i++ {
		l, r := s[i][0], s[i][1]
		if len(result) == 0 || result[len(result)-1][1] < l {
			result = append(result, []int{l, r})
		} else {
			result[len(result)-1][1] = max(r, result[len(result)-1][1])
		}
	}
	return result
}
