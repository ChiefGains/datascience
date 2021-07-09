//package for creating and doing math with matrices
//KNOWN ISSUES:
package datascience

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

//Matrix stores coordinates for a 2d matrix. Support may be added in the future for n dimensional
//matrices, probalby just by embedding multiple 2d matrix structures into a superstruct as a slice
//of matrices. Who knows, I haven't really thought about the implementation yet
type Matrix struct { //data type stores coordinates for a 2D matrix
	X, Y        int      //dimesions of the matrix
	ColumnNames []string //this is for tables that have columns, for use in later functions
	Matrix      [][]float64
	Cof         [][]float64
	Adj         [][]float64
}

//stringMat is for storing a string version of a matrix for printing and visualization
type stringMat struct {
	X, Y int        //dimensions of the matrix
	Mat  [][]string //this is where string versions of the float64 values will be held
}

//MakeMatrix takes dimensions as input and returns a pointer to a matrix struct
func MakeMatrix(x, y int) *Matrix {

	matrix := make([][]float64, x)

	for i := range matrix {
		matrix[i] = make([]float64, y)
	}

	columns := make([]string, x)

	var cof [][]float64

	var adj [][]float64

	mat := &Matrix{
		X:           x,
		Y:           y,
		ColumnNames: columns,
		Matrix:      matrix,
		Cof:         cof,
		Adj:         adj,
	}

	return mat
}

//StringMatrix converets type Matrix to type StringMat, mainly for use in PrintMat function
func StringMatrix(x, y int) *stringMat {
	a := make([][]string, x)

	for i := range a {
		a[i] = make([]string, y)
	}

	mat := &stringMat{x, y, a}

	return mat
}

//PrintMat converts a float64 matrix to string and prints
func (m *Matrix) PrintMat() *stringMat {
	a := StringMatrix(m.X, m.Y)
	for i := 0; i < len(m.Matrix); i++ {
		for j := 0; j < len(m.Matrix[i]); j++ {
			a.Mat[i][j] = strconv.FormatFloat(m.Matrix[i][j], 'f', 2, 64)
		}
		b := strings.Join(a.Mat[i], " ")
		fmt.Println(b)

	}
	fmt.Println("")
	return a
}

//RandFill fills an empty matrix with random values
func (m *Matrix) RandFill() *Matrix {
	for i := 0; i < m.X; i++ {
		for j := 0; j < m.Y; j++ {
			m.Matrix[i][j] = rand.NormFloat64()
		}
	}
	return m
}

//MatMult takes two matrices as inputs and multiplies them if they can be multiplied, or returns an error
func MatMult(first, second *Matrix) (*Matrix, error) {
	res := MakeMatrix(first.X, second.Y)
	a := first.Matrix
	b := second.Matrix
	c := res.Matrix
	if first.Y != second.X { //check that # of columns in first matches # of rows in second
		return res, fmt.Errorf("Matrices cannot be multiplied.")
	}
	for i := 0; i < res.X; i++ {
		for j := 0; j < second.Y; j++ {
			for k := 0; k < first.Y; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return res, nil
}

//IdentityMatrix is a method that returns a pointer to an identity matrix of the given matrix
//If no identity matrix can be found, it returns an error
func (m *Matrix) IdentityMatrix() (*Matrix, error) {

	var res *Matrix

	if m.X != m.Y {
		return &Matrix{}, fmt.Errorf("No identity matrix found")
	}

	res = MakeMatrix(m.X, m.Y)

	for i := range res.Matrix {
		res.Matrix[i][i] = 1
	}

	return res, nil
}

//Invert is a method which finds the inverse of a given matrix
//Currently only works for 2x2 matrices
func (m *Matrix) Invert() (*Matrix, error) {
	res := MakeMatrix(m.X, m.Y)

	if m.X != m.Y {
		return res, fmt.Errorf("Only square matrices may be inverted")
	}

	//I don't think this step is necessary. Like, at all
	for i := range m.Matrix {
		copy(res.Matrix[i], m.Matrix[i])
	}

	switch m.X { //ew. I'm not sure a switch statement is necessary here, just write a better algorithm
	case 2:
		switch m.Y {
		case 2:
			det := 1 / ((m.Matrix[0][0] * m.Matrix[1][1]) - (m.Matrix[0][1] * m.Matrix[1][0]))
			for i := range res.Matrix {
				for j := range res.Matrix[i] {
					val := (res.Matrix[i][j] * det)
					res.Matrix[i][j] = val
				}
			}
			a, b := res.Matrix[0][0], res.Matrix[1][1] //seems unnecessary to declare them in their own line
			res.Matrix[0][0] = b
			res.Matrix[1][1] = a
			res.Matrix[0][1] *= -1
			res.Matrix[1][0] *= -1
		default:
			return res, fmt.Errorf("Still working on this feature")
		}
	default:
		return res, fmt.Errorf("Still working on this feature")
	}
	return res, nil
}
