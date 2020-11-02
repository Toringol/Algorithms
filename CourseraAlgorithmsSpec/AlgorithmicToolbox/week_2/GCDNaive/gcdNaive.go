package main

func gcdNaive(a, b int64) int64 {
	result := int64(0)
	for i := int64(1); i < a+b; i++ {
		if a%i == 0 && b%i == 0 {
			result = i
		}
	}

	return result
}

func main() {

}
