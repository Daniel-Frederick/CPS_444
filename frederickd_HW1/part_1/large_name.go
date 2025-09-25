// Part 1: Question 3 package main
package main

import (
	"fmt"
)

func main() {
	D := []string{
		"*****  ",
		"*    * ",
		"*     *",
		"*     *",
		"*    * ",
		"*****  ",
	}

	A := []string{
		"   *   ",
		"  * *  ",
		" *   * ",
		"*******",
		"*     *",
		"*     *",
	}

	N := []string{
		"*     *",
		"**    *",
		"* *   *",
		"*  *  *",
		"*   * *",
		"*    **",
	}

	I := []string{
		"*******",
		"   *   ",
		"   *   ",
		"   *   ",
		"   *   ",
		"*******",
	}

	E := []string{
		"*******",
		"*      ",
		"*      ",
		"*****  ",
		"*      ",
		"*******",
	}

	L := []string{
		"*      ",
		"*      ",
		"*      ",
		"*      ",
		"*      ",
		"*******",
	}

	letters := [][]string{D, A, N, I, E, L}

	for row := 0; row < len(D); row++ {
		for _, letter := range letters {
			fmt.Print(letter[row] + "   ")
		}
		fmt.Println()
	}
}

