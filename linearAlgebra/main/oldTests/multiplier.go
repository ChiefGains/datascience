//matrix multiplier version 0.2.0
//generates random matrices and multiplies them

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func RandMatrix(x, y int) [][]int { //generates a random matrix of n height and length
	a := make([][]int, x)

	for i := range a {
		a[i] = make([]int, y)
		for j := range a[i] {
			a[i][j] = rand.Intn(10)
		}
	}
	return a
}

func BlankMatrix(x, y int) [][]int { //generates an empty matrix
	a := make([][]int, x)

	for i := range a { 
		a[i] = make([]int, y)
		for j := range a[i] { //I'm sure this loop is unnecessary, since the null value of a slice is already 0
			a[i][j] = 0
		}
	}
	return a
}

func StringMatrix(x, y int) [][]string { //converts an integer matrix to a string
	a := make([][]string, x)

	for i := range a {
		a[i] = make([]string, y)
	}
	return a
}

func MatMult(a, b [][]int) [][]int { //this algorithm multiplies matrices of any size, provided they can be multipied
	c := BlankMatrix(len(a), len(b[0]))
	switch {
	case len(a[0]) != len(b): //verify that the matrices in question can, in fact, be multiplied
		fmt.Println("These matrices cannot be multiplied")
	default:
		switch { //the original algorithm worked fine for square matrices, but I had to write separate algorithms for
			 //matrices that were not square to correctly multiply
		case len(a[0]) < len(b[0]):
			for i := 0; i < len(a); i++ {
				for j := 0; j < len(b[0]); j++ {
					for k := 0; k < len(b); k++ {
						c[i][j] += a[i][k] * b[k][j]
					}
				}
			}
		default:
			for i := 0; i < len(a); i++ {
				for j := 0; j < len(b[0]); j++ {
					for k := 0; k < len(a[0]); k++ {
						c[i][j] += a[i][k] * b[k][j]
					}
				}
			}

		}

	}
	return c
}

func PrintMat(mat [][]int) [][]string { //takes an integer matrix and outputs a formatted string
	a := StringMatrix(len(mat), len(mat[0]))
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			a[i][j] = strconv.Itoa(mat[i][j])
		}
		b := strings.Join(a[i], " ")
		fmt.Println(b)

	}
	fmt.Println("")
	return a
}

func main() {

	a := RandMatrix(2, 3) //generate/initialize two random matrices of certain height and length
	b := RandMatrix(3, 3)

	PrintMat(a) //print these matrix values
	PrintMat(b)

	ab := MatMult(a, b) //create a new matrix by multiplying matrices a and b

	PrintMat(ab) //print the values for matrix ab

}
