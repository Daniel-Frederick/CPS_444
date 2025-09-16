// Part 2: question 3

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Must enter one number as a radius!")
		return
	}

	radius, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Please enter a valid number!")
		return
	}

	fmt.Println("Area:          ", area(radius))
	fmt.Println("Circumference: ", circum(radius))
	fmt.Println("Volume:        ", volume(radius))
	fmt.Println("Surface Area:  ", surfaceArea(radius))
}

func area(r float64) float64 {
	return math.Pi * r * r
}

func circum(r float64) float64 {
	return 2 * math.Pi * r
}

func volume(r float64) float64 {
	return (4.0 / 3.0) * math.Pi * math.Pow(r, 3)
}

func surfaceArea(r float64) float64 {
	return 4 * math.Pi * r * r
}
