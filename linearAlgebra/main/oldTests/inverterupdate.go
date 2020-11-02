//update to the inverter, hopefully to be able to get it to invert matrices larger than 2x2
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	m := MakeMat(3, 3)
	m.RandFill()
	m.PrintMat()

	m.Update()

	m.PrintAll()

	m1 := MakeMat(3, 3)
	m1.Mat = m.Inv
	m1.PrintMat()

	res := MatMult(m, m1)

	res.PrintMat()
}

type Matrix struct { //data type stores coordinates for a 2D matrix
	X, Y int
	Mat  [][]float64
	Cof  [][]float64
	Adj  [][]float64
	Inv  [][]float64
	Det  float64
}

func (m *Matrix) findDet() float64 { //find the determinant of a given matrix
	var det float64

	switch {
	case m.X != m.Y:
		fmt.Println("Error: can only find determinant for square matrices") //yes, I know, this should be an actual error, but I don't know how to do that in go yet
		return 0
	case m.X == 2:
		det = ((m.Mat[0][0] * m.Mat[1][1]) - (m.Mat[0][1] * m.Mat[1][0]))
	default:
		det = 0
		for i := range m.Mat {
			det += m.Mat[0][i] * m.Cof[0][i]

		}
	}

	return det

}

func (m *Matrix) findCof() *Matrix { //finds cofactor of matrix
	//so far only working on 3x3 functionality
	if m.X != m.Y {
		fmt.Println("Can only find cofactors for square matrices")
		return m
	}

	temp := MakeMat(m.X-1, m.Y-1)

	var nums []int

	for i := range m.Mat {
		nums = append(nums, i)
	}

	fmt.Println("Nums:")
	fmt.Println(nums)

	for i := range nums {
		var nums1 []int
		for j := range nums {
			if nums[j] != i {
				nums1 = append(nums1, j)
			}
		}
		for j := range nums {
			var nums2 []int
			for k := range nums {
				if nums[k] != j {
					nums2 = append(nums2, k)
				}
			}

			for tempX, mX := range nums1 {
				for tempY, mY := range nums2 {
					temp.Mat[tempX][tempY] = m.Mat[mX][mY]
				}
			}
			m.Cof[i][j] = temp.findDet()
		}

	}

	m.changeSigns()

	return m

}

func (m *Matrix) changeSigns() *Matrix {
	var sign float64
	sign = 1
	for i := range m.Cof {
		for j := range m.Cof[i] {
			m.Cof[i][j] *= sign
			sign *= -1
		}
	}
	return m
}

func (m *Matrix) findAdj() *Matrix { //finds the adjugate of a matrix
	if m.X != m.Y {
		fmt.Println("Can only find the adjugate for a square matrix")
		return m
	}

	for i := range m.Cof {
		for j := range m.Cof[i] {
			m.Adj[j][i] = m.Cof[i][j]
		}
	}

	return m

}

func (m *Matrix) Update() *Matrix { //function to find the inverse of a matrix
	if m.X != m.Y {
		fmt.Println("Only square matrices can be inverted")
		return m
	}

	switch m.X {
	case 2:
		det := 1 / ((m.Mat[0][0] * m.Mat[1][1]) - (m.Mat[0][1] * m.Mat[1][0]))
		for i := range m.Mat {
			for j := range m.Mat[i] {
				m.Inv[i][j] = (m.Mat[i][j] * det)
			}
		}
		a, b := m.Inv[0][0], m.Inv[1][1]
		m.Inv[0][0] = b
		m.Inv[1][1] = a
		m.Inv[0][1] *= -1
		m.Inv[1][0] *= -1
	default:
		m.findCof()
		m.findAdj()
		m.Det = m.findDet()

		for i := range m.Adj {
			for j := range m.Adj {
				m.Inv[i][j] = m.Adj[i][j] * (1 / m.Det)
			}
		}
	}
	return m
}

//BELOW THIS POINT IS ALL THE FUNCTIONS I'M NOT WORKING ON THAT ARE NECESSARY TO MAKE THIS CODE WORK
//Ok but I lied and I made some changes to this code

func MakeMat(x, y int) *Matrix { //Instantiates and creates a pointer to an empty 2D matrix

	mat := make([][]float64, x)

	for i := range mat {
		mat[i] = make([]float64, y)
	}

	cof := make([][]float64, x)

	for i := range cof {
		cof[i] = make([]float64, y)
	}

	adj := make([][]float64, x)

	for i := range adj {
		adj[i] = make([]float64, y)
	}

	inv := make([][]float64, x)

	for i := range inv {
		inv[i] = make([]float64, y)
	}

	var det float64

	m := &Matrix{x, y, mat, cof, adj, inv, det}

	return m
}

func StringMatrix(x, y int) [][]string { //creates a blank matrix to hold a string (for the PrintMat func)
	a := make([][]string, x)

	for i := range a {
		a[i] = make([]string, y)
	}
	return a
}

func (m *Matrix) PrintMat() [][]string { //Converts an float64 matrix to string and prints
	a := StringMatrix(m.X, m.Y)
	for i := 0; i < len(m.Mat); i++ {
		for j := 0; j < len(m.Mat[i]); j++ {
			a[i][j] = strconv.FormatFloat(m.Mat[i][j], 'f', 2, 64)
		}
		b := strings.Join(a[i], " ")
		fmt.Println(b)

	}
	fmt.Println("")
	return a
}

func (m *Matrix) PrintCof() [][]string { //Converts an float64 matrix to string and prints
	a := StringMatrix(m.X, m.Y)
	for i := 0; i < len(m.Mat); i++ {
		for j := 0; j < len(m.Mat[i]); j++ {
			a[i][j] = strconv.FormatFloat(m.Cof[i][j], 'f', 2, 64)
		}
		b := strings.Join(a[i], " ")
		fmt.Println(b)

	}
	fmt.Println("")
	return a
}

func (m *Matrix) PrintAdj() [][]string { //Converts an float64 matrix to string and prints
	a := StringMatrix(m.X, m.Y)
	for i := 0; i < len(m.Mat); i++ {
		for j := 0; j < len(m.Mat[i]); j++ {
			a[i][j] = strconv.FormatFloat(m.Adj[i][j], 'f', 2, 64)
		}
		b := strings.Join(a[i], " ")
		fmt.Println(b)

	}
	fmt.Println("")
	return a
}

func (m *Matrix) PrintInv() [][]string { //Converts an float64 matrix to string and prints
	a := StringMatrix(m.X, m.Y)
	for i := 0; i < len(m.Mat); i++ {
		for j := 0; j < len(m.Mat[i]); j++ {
			a[i][j] = strconv.FormatFloat(m.Inv[i][j], 'f', 2, 64)
		}
		b := strings.Join(a[i], " ")
		fmt.Println(b)

	}
	fmt.Println("")
	return a
}

func (m *Matrix) PrintAll() {
	fmt.Println("Matrix:")
	m.PrintMat()
	fmt.Println("Cofactors:")
	m.PrintCof()
	fmt.Println("Adjugate:")
	m.PrintAdj()
	fmt.Println("Determinant:\n", m.Det)
	fmt.Println("Inverse:")
	m.PrintInv()
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
