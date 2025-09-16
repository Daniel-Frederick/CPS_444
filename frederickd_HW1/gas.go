package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println(`Must enter 3 parameters:
1. The number of gallons of gas in the tank
2. The fuel efficiency in miles per gallon
3. The price of gas per gallon`)
		return
	}

	gallons, err1 := strconv.ParseFloat(os.Args[1], 64)
	mpg, err2 := strconv.ParseFloat(os.Args[2], 64)
	pricePerGallon, err3 := strconv.ParseFloat(os.Args[3], 64)

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Error: All parameters must be numbers.")
		return
	}

	fmt.Printf("Cost Per 100 Miles: $%.2f\n", costPer100miles(mpg, pricePerGallon))
	fmt.Printf("How far the car can go with fuel: %.2f miles\n", howFarTheCarGo(gallons, mpg))
}

func costPer100miles(mpg, pricePerGallon float64) float64 {
	// cost = (100 miles / mpg) * price per gallon
	return (100 / mpg) * pricePerGallon
}

func howFarTheCarGo(gallons, mpg float64) float64 {
	// distance = gallons * mpg
	return math.Max(gallons*mpg, 0)
}
