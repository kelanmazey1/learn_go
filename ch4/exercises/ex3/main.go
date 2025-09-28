package main

import "fmt"

func main() {
	var total int

	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Println(total)
	}

	fmt.Printf("Total in outer block: %d\n", total)

	// Most likely bug here is that you'd expect total in 'main' function block to have increased
}
