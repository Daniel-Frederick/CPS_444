// Part 1: Question 5
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <name>")
		return
	}

	name := os.Args[1]

	fmt.Println("        -----")
	fmt.Printf("       / Hello \\ \n")
	fmt.Printf("      <  %s!  |\n", name)
	fmt.Println("       \\       /")
	fmt.Println("        -----")

	fmt.Println("  /\\_/\\  ")
	fmt.Println(" ( o.o ) ")
	fmt.Println(" > ^ ^ < ")
}

