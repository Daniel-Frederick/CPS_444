package main

import (
    "fmt"
    "math"
    "os"
    "strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please enter exactly two numbers!")
		return
	}

	num1, err1 := strconv.Atoi(os.Args[1])
	num2, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil {
		fmt.Println("Please enter valid integers!")
		return
	}

	fmt.Println("sum:",        sum(num1, num2))
	fmt.Println("difference:", difference(num1, num2))
	fmt.Println("product:",    product(num1, num2))
	fmt.Println("average:",    average(num1, num2))
	fmt.Println("distance:",   distance(num1, num2))
	fmt.Println("maximum:",    maximum(num1, num2))
	fmt.Println("minimum:",    minimum(num1, num2))
}

func sum(a, b int) int {
	return a + b
}

func difference(a, b int) int {
	return a - b
}

func product(a, b int) int {
	return a * b
}

func average(a, b int) float64 {
	return float64((a + b) / 2)
}

func distance(a, b int) float64 {
    return math.Abs(float64(a - b))
}

func maximum(a, b int) float64 {
    return math.Max(float64(a), float64(b))
}

func minimum(a, b int) float64 {
    return math.Min(float64(a), float64(b))
}
