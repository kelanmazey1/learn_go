package exercises

import "fmt"

func Ex3() {

	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	person_one := Employee{"Jane", "Doe", 1}

	person_two := Employee{
		firstName: "Joe",
		lastName:  "Bloggs",
		id:        2,
	}
	var person_three = Employee{}
	person_three.firstName = "Kelan"
	person_three.lastName = "Mazey"
	person_three.id = 3

	fmt.Println(person_one)
	fmt.Println(person_two)
	fmt.Println(person_three)
}
