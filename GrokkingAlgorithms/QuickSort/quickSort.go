package main

import (
	"fmt"
)

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := len(arr) / 2
	lessArr := []int{}
	moreArr := []int{}
	resultArr := []int{}

	for i := 0; i < len(arr); i++ {
		if i == pivot {
			continue
		}
		if arr[i] <= arr[pivot] {
			lessArr = append(lessArr, arr[i])
		} else {
			moreArr = append(moreArr, arr[i])
		}
	}

	resultArr = append(resultArr, QuickSort(lessArr)...)
	resultArr = append(resultArr, arr[pivot])
	resultArr = append(resultArr, QuickSort(moreArr)...)

	return resultArr
}

func main() {
	arr := []int{1, -5, 6, 4, -100, 17, 3}
	arr = QuickSort(arr)
	fmt.Println(arr)
}
