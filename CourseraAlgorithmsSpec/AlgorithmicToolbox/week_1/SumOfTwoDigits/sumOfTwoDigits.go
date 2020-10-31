package main

import "fmt"

func sumOfTwoDigits(firstDigit, secondDigit int) int {
	return firstDigit + secondDigit
}

func main() {
	var firstDigit, secondDigit int

	fmt.Scan(&firstDigit, &secondDigit)

	fmt.Println(sumOfTwoDigits(firstDigit, secondDigit))
}
