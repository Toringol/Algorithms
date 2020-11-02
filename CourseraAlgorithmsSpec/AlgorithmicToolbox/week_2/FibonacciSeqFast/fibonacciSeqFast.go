package main

func fibonacciSeqFast(number int64) int64 {
	if number <= 1 {
		return number
	}

	result := int64(0)
	prevVal := int64(1)
	for i := int64(0); i < number; i++ {
		result, prevVal = prevVal, result+prevVal
	}

	return result
}

func main() {

}
