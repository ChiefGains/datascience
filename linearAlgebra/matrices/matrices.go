//package for creating and doing math with matrices
//KNOWN ISSUES:
package matrices

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Matrix struct { //data type stores coordinates for a 2D matrix
	X, Y int
	Mat  [][]float64
	Cof  [][]float64
	Adj  [][]float64
}

type StringMat struct { //data type holds values for a 2D matrix of string values
	X, Y int
	Mat  [][]string
}

func MakeMat(x, y int) *Matrix { //Instantiates and creates a pointer to an empty 2D matrix

	a := make([][]float64, x)

	var b [][]float64

	var c [][]float64

	for i := range a {
		a[i] = make([]float64, y)
	}

	mat := &Matrix{x, y, a, b, c}

	return mat
}

func StringMatrix(x, y int) *StringMat { //converets type Matrix to type StringMat, mainly for use in PrintMat function
	a := make([][]string, x)

	for i := range a {
		a[i] = make([]string, y)
	}

	mat := &StringMat{x, y, a}

	return mat
}

func (m *Matrix) PrintMat() *StringMat { //Converts an float64 matrix to string and prints
	a := StringMatrix(m.X, m.Y)
	for i := 0; i < len(m.Mat); i++ {
		for j := 0; j < len(m.Mat[i]); j++ {
			a.Mat[i][j] = strconv.FormatFloat(m.Mat[i][j], 'f', 2, 64)
		}
		b := strings.Join(a.Mat[i], " ")
		fmt.Println(b)

	}
	fmt.Println("")
	return a
}

func (m *Matrix) RandFill() *Matrix { //fills a matrix struct with empty values
	for i := 0; i < m.X; i++ {
		for j := 0; j < m.Y; j++ {
			m.Mat[i][j] = rand.NormFloat64()
		}
	}
	return m
}

func MatMult(uno, dos *Matrix) *Matrix { //multiplies 2 matrices, if they can be multiplied
	res := MakeMat(uno.X, dos.Y)
	a := uno.Mat
	b := dos.Mat
	c := res.Mat
	switch {
	case uno.Y != dos.X: //check that # of columns in uno matches # of rows in dos
		fmt.Println("These matrices cannot be multiplied")
	default:
		for i := 0; i < res.X; i++ {
			for j := 0; j < dos.Y; j++ {
				for k := 0; k < uno.Y; k++ {
					c[i][j] += a[i][k] * b[k][j]
				}
			}
		}

	}

	return res
}

func (m *Matrix) IdentityMatrix() *Matrix { //creates an identity matrix for a given input matrix

	var res *Matrix

	switch {
	case m.X != m.Y:
		fmt.Println("Nah. Not today")
		return m
	default:
		res = MakeMat(m.X, m.Y)
	}

	for i := range res.Mat {
		res.Mat[i][i] = 1
	}

	return res
}

func (m *Matrix) Invert() *Matrix { //finds the inverse of a matrix. Currently only works for 2x2
	I := MakeMat(m.X, m.Y)

	if m.X != m.Y {
		fmt.Println("Only square matrices can be inverted")
		return m
	}

	for i := range m.Mat {
		copy(I.Mat[i], m.Mat[i])
	}

	switch m.X {
	case 2:
		switch m.Y {
		case 2:
			det := 1 / ((m.Mat[0][0] * m.Mat[1][1]) - (m.Mat[0][1] * m.Mat[1][0]))
			for i := range I.Mat {
				for j := range I.Mat[i] {
					val := (I.Mat[i][j] * det)
					I.Mat[i][j] = val
				}
			}
			a, b := I.Mat[0][0], I.Mat[1][1]
			I.Mat[0][0] = b
			I.Mat[1][1] = a
			I.Mat[0][1] *= -1
			I.Mat[1][0] *= -1
		default:
			fmt.Println("I can't do that yet")
		}
	default:
		fmt.Println("I can't do that yet")
	}
	return I
}
