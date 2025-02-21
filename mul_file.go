package gopkg

import "fmt"

func init() {
	fmt.Println("mul package: Package initialization...")
}

func MyMul(a, b int) int {
	return a * b
}
