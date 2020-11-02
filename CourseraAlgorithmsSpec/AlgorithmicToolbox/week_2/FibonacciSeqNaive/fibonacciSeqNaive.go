package main

func fibonacciSeqNaive(number int) int64 {
	if number <= 1 {
		return int64(number)
	}
	return int64(fibonacciSeqNaive(number-1)) + fibonacciSeqNaive(number-2)
}

func main() {

}
