package main

func gcdEuclidRecursive(a, b int64) int64 {
	if b == 0 {
		return a
	}

	r := a % b
	return gcdEuclidRecursive(b, r)
}

func gcdEuclidIterative(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func main() {

}
