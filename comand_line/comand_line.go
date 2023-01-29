package main

import (
	"fmt"
	"os"
)

func main() {

	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s = s + sep + os.Args[i]

		sep = " "
	}

	fmt.Println(s)
	test_main()

}

// The other way

func test_main() {

	s := ""
	var i int
	for j, arg := range os.Args[1:] {
		i = i + 1
		fmt.Println(j)
		s = string(j) + arg

		fmt.Println(s)
	}

	fmt.Println(s)
}
