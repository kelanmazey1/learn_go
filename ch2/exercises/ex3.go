package exercises

import "fmt"

func Ex3() {
	fmt.Println("Hello")

	var b byte
	var smallI int32
	var bigI uint64

	b = 255
	smallI = 2147483647
	bigI = 18446744073709551615

	b += 1
	smallI += 1
	bigI += 1

	fmt.Printf("b: %d\n", b)
	fmt.Printf("smallI: %d\n", smallI)
	fmt.Printf("bigI: %d\n", bigI)
}
