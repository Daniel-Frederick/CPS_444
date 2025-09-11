package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Must enter one number in meters!")
		return
	}

	meters, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Please enter a valid number!")
		return
	}

	fmt.Println("Miles:  ", miles(meters))
	fmt.Println("Feet:   ", feet(meters))
	fmt.Println("Inches: ", inches(meters))
}

func miles(m float64) float64 {
	return m * 0.000621371
}

func feet(m float64) float64 {
	return m * 3.28084
}

func inches(m float64) float64 {
	return m * 39.3701
}
