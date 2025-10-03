package main

import "fmt"

func prefixer(pfx string) func(string) string {
	return func(s string) string {
		return pfx + s
	}

}

func main() {
	helloPrefix := prefixer("Hello ")
	fmt.Println(helloPrefix("Kel"))
	fmt.Println(helloPrefix("El"))
}
