package exercises

import "fmt"

func Ex1() {
	greetings := []string{
		"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт",
	}

	first_two := greetings[:2]
	two_three_four := greetings[1:4]
	fourth_fifth := greetings[3:]

	fmt.Println(greetings)
	fmt.Println(first_two)
	fmt.Println(two_three_four)
	fmt.Println(fourth_fifth)

}
