// Part 1: Question 2

package main

import (
	"fmt"
	"os"
)

func main() {
	if (len(os.Args) < 2) {
		fmt.Println("No Name given!")
		return;
	}
	if (len(os.Args) >2) {
	  fmt.Println("Only enter one name!")
		return;
	}

	name := os.Args[1]

	var box string = "+"

	// Top box line
	for i := 0; i < len(name) + 2; i++ {
		box += "-"
	}
	box += "+\n"

	// Middle box line
	box += "| " + name + " |\n"

	// Bottom box line
	box += "+"
	for i := 0; i < len(name) + 2; i++ {
		box += "-"
	}
	box += "+"

	fmt.Println(box)
}
