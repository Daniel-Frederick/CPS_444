// Part 2: question 4

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 6 {
		fmt.Println(`Must enter these 5 parameters:
1. The cost of a new car
2. The estimated miles driven per year
3. The estimated gas price (per gallon)
4. The efficiency in miles per gallon
5. The estimated resale value after 5 years`)
		return
	}

	// Parse command-line arguments
	cost, err1 := strconv.ParseFloat(os.Args[1], 64)
	milesPerYear, err2 := strconv.ParseFloat(os.Args[2], 64)
	gasPrice, err3 := strconv.ParseFloat(os.Args[3], 64)
	mpg, err4 := strconv.ParseFloat(os.Args[4], 64)
	resaleValue, err5 := strconv.ParseFloat(os.Args[5], 64)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		fmt.Println("Error: All parameters must be numbers.")
		return
	}

	totalCost := ownCarForFiveYears(cost, milesPerYear, gasPrice, mpg, resaleValue)
	fmt.Printf("Total cost of owning the car for five years: $%.2f\n", totalCost)
}

func ownCarForFiveYears(cost, milesPerYear, gasPrice, mpg, resaleValue float64) float64 {
	// Calculate fuel cost over 5 years
	fuelCost := ((milesPerYear * 5) / mpg) * gasPrice

	// Total cost = purchase cost + fuel cost - resale value
	total := cost + fuelCost - resaleValue

	// Make sure result isn't negative
	return math.Max(total, 0)
}
