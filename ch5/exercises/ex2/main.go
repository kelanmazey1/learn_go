package main

import (
	"fmt"
	"os"
)

func main() {
	f_len, err := fileLen("test.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(f_len)
}

func fileLen(name string) (int, error) {
	f, err := os.Open(name)
	if err != nil {
		return 0, err
	}

	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return 0, err
	}

	return int(info.Size()), nil

}
