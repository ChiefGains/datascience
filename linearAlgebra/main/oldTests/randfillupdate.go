//update to allow random filling of a matrix with different value types
package main

import("fmt"
		"src\learning\linearAlgebra\matrices")

func (m *Matrix) RandFill(t string) *Matrix { //fills a matrix struct with empty values

	//THIS DOES NOT WORK CURRENTLY
	//Must either move type to MakeMat and perform type switch,
	//or find a way to fill empty mat with interface. Or,
	//manually convert each cell to the correct type, which sounds
	//like a stupid way to handle that

	switch t {
	case "float64":
		for i := 0; i < m.X; i++ {
			for j := 0; j < m.Y; j++ {
				m.Mat[i][j] = rand.NormFloat64()
			}
		}
	case "float32":
		for i := 0; i < m.X; i++ {
			for j := 0; j < m.Y; j++ {
				m.Mat[i][j] = rand.NormFloat32()
			}
		}
	case "int":
		for i := 0; i < m.X; i++ {
			for j := 0; j < m.Y; j++ {
				m.Mat[i][j] = rand.Intn(20)
			}
		}
	default:
		fmt.Println("Please enter argument float64, float32, or int")
	}
	return m
}

func main(){

}