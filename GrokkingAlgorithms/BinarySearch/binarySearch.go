package main

import (
	"errors"
	"fmt"
)

func BinarySearch(arr []int, aim int) (int, error) {
	if len(arr) < 1 {
		return -1, errors.New("Empty array")
	}

	start := 0
	end := len(arr) - 1

	for start <= end {
		middle := (start + end) / 2
		if arr[middle] < aim {
			start = middle + 1
		} else if arr[middle] > aim {
			end = middle - 1
		} else {
			return middle, nil
		}
	}

	return -1, errors.New("No such element in array")
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	aim := 2

	pos, err := BinarySearch(arr, aim)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pos)
}
