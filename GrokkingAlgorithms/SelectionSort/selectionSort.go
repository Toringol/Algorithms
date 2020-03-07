package main

import (
	"errors"
	"fmt"
)

func Swap(a, b *int) {
	*a, *b = *b, *a
}

func FindMinIndex(arr []int, startPos, endPos int) int {
	min := startPos
	for i := startPos; i <= endPos; i++ {
		if arr[i] < arr[min] {
			min = i
		}
	}
	return min
}

func SelectionSort(arr []int) ([]int, error) {
	if len(arr) < 1 {
		return []int{}, errors.New("No elements in arr")
	}

	for i := 0; i < len(arr)-1; i++ {
		Swap(&arr[i], &arr[FindMinIndex(arr, i, len(arr)-1)])
	}

	return arr, nil
}

func main() {
	arr := []int{1, -5, 6, 4, -100, 17, 3}
	SelectionSort(arr)
	fmt.Println(arr)
}
