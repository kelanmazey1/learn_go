package exercises

import "fmt"

func Ex2() {
	message := "Hi 👩 and 👨"
	runes := []rune(message)

	fmt.Println(message[3])       // first byte of the lady
	fmt.Println(message[3:8])     // full emoji (4 bytes long)
	fmt.Println(runes[3])         // unicode value
	fmt.Println(string(runes[3])) // actual emoji
}
