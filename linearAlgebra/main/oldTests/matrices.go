package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix struct { //this type creates a 2 dimensional matrix object for integers
	X, Y int
	Mat  [][]int
}

func makeMat(x, y int) *Matrix {

	a := make([][]int, y)

	for i := range a {
		a[i] = make([]int, x)
	}

	mat := &Matrix{x, y, a}

	return mat
}

func StringMatrix(x, y int) [][]string { //converts an integer matrix to a string
	a := make([][]string, x)

	for i := range a {
		a[i] = make([]string, y)
	}
	return a
}

func (m *Matrix) PrintMat() [][]string {
	a := StringMatrix(m.Y, m.X)
	for i := 0; i < len(m.Mat); i++ {
		for j := 0; j < len(m.Mat[i]); j++ {
			a[i][j] = strconv.Itoa(m.Mat[i][j])
		}
		b := strings.Join(a[i], " ")
		fmt.Println(b)

	}
	fmt.Println("")
	return a
}

func (m *Matrix) AddColumn() *Matrix {
	for i := range m.Mat {
		m.Mat[i] = append(m.Mat[i], 0)
	}
	m.X++
	return m
}

func (m *Matrix) AddRow() *Matrix {
	newRow := make([]int, m.X)
	m.Mat = append(m.Mat, newRow)
	m.Y++
	return m
}

func main() {

	m := makeMat(4, 4)

	m.PrintMat()
	m.AddColumn()
	m.PrintMat()
	m.AddRow()
	m.PrintMat()

}
