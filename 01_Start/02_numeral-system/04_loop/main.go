package main

import "fmt"

func main() {

	for x := 1; x < 150; x++ {
		fmt.Printf("%d - %b - %q\n", x, x, x)
	}

}
