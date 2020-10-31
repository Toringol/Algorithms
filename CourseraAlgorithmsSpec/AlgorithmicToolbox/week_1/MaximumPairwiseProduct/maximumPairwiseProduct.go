package main

import "fmt"

func MaximumPairwiseProduct(inputArr []int) int64 {
	var maxIndexFirst int = -1
	for i := range inputArr {
		if maxIndexFirst == -1 || inputArr[i] > inputArr[maxIndexFirst] {
			maxIndexFirst = i
		}
	}

	var maxIndexSecond int = -1
	for j := range inputArr {
		if j != maxIndexFirst && (maxIndexSecond == -1 || inputArr[j] > inputArr[maxIndexSecond]) {
			maxIndexSecond = j
		}
	}

	return int64(inputArr[maxIndexFirst]) * int64(inputArr[maxIndexSecond])
}

func main() {
	var arrLength int

	fmt.Scan(&arrLength)

	inputArr := make([]int, arrLength)
	for i := range inputArr {
		fmt.Scan(&inputArr[i])
	}

	fmt.Println(MaximumPairwiseProduct(inputArr))
}
