package exercises

import (
	"fmt"
	"math/rand"
)

func Ex1And2() {

	myInts := make([]int, 100)

	for i := 0; i < 100; i++ {
		myInts[i] = rand.Intn(100)
	}

	fmt.Println(myInts)
	fmt.Println(len(myInts))

	for v := range myInts {
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
