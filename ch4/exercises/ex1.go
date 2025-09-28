package exercises

import (
	"fmt"
	"math/rand"
)

func Ex1And2() {

	my_ints := make([]int, 100)

	for i := 0; i < 100; i++ {
		my_ints[i] = rand.Intn(100)
	}

	fmt.Println(my_ints)
	fmt.Println(len(my_ints))

	for v := range my_ints {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six!")
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}

}
