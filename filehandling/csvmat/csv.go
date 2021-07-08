//Test to make struct that will hand csv data, and to
//see how well it interacts with functions from the
//matrices.go package. Some things will definitely not
//work, given that they are methods implemented on certain
//types within the matrices package

//IMPORTANT:
//it may be easiest to nest a Matrix struct *inside* a csv struct, so that it can use all the same operations
//in this case, basically the only two things a DataFrame would need to hold would be a Matrix, and Column Names
//The tricky part of this is that it will need to know if there are column names

//Other Note: this version will not automatically detect if there are column names, it
//will just assume that there are and fill them in. Later version will add sophistication

package datascience

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type dataFloat struct {
	X    float64
	Null bool
}

type DataFrame struct { //this type, or type test if you will, is going to just straight up stick a matrix in itself

	X, Y int

	Columns []string
	DF      [][]*dataFloat
}

func check(err error) { //well this is bad error handling. And by that I mean it isn't really even error handling
	if err != nil {
		log.Fatal(err)
	}

}

func findRows(file *os.File) int {
	r := csv.NewReader(file)
	x := 0

	for { //increase x by one for every new row
		_, err := r.Read()
		if err == io.EOF {
			break
		}

		check(err)

		x++
	}

	return x

}

func findColumns(file *os.File) int {
	r := csv.NewReader(file)
	y := 0

	for i := 0; i < 1; i++ { //increase y by one for every new column
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		check(err)

		for range record {
			y++
		}
	}

	return y
}

func LabelColumns(f *os.File) []string { //get Columns for struct
	r := csv.NewReader(f)

	columns := make([]string, 0)

	for i := 0; i < 1; i++ {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		check(err)

		for j := range record {
			columns = append(columns, record[j])
		}
	}
	return columns
}

func makeDF(f *os.File) *DataFrame { //instantiates a new DataFrame. Will need to check if data can be converted to Float. Top row, if all strings, will be placed in "Column Names". This function should be merged with "LoadCSV" from readCSV.g

	x := findRows(f)
	f.Seek(0, 0)
	y := findColumns(f)
	f.Seek(0, 0)

	a := make([]string, y)
	b := make([][]*dataFloat, x-1)

	for i := range b {
		b[i] = make([]*dataFloat, y)
	}

	frame := &DataFrame{x, y, a, b}

	return frame

}

func newDataFloat(x string) *dataFloat { //converts string data to dataFloat
	if s, err := strconv.ParseFloat(x, 64); err == nil { //if the value can be converted to a float, do so
		return &dataFloat{s, false}
	}

	return &dataFloat{0, true} //value cannot be converted, return nil = true
}

func LoadCsv(file string) *DataFrame {
	data, error := os.Open(file)

	if error != nil {
		log.Fatalln("Unable to load csv", error)
	}

	df := makeDF(data)

	r := csv.NewReader(data)

	data.Seek(0, 0)

	for row := 0; ; row++ { //iterate through the rows of the csv
		record, err := r.Read() //create reader for CSV
		if err == io.EOF {      //check for end of file
			break
		}

		if err != nil { //log errors
			log.Fatal(err)
		}

		if row == 0 {
			for column := range record {
				df.Columns[column] = record[column]
			}
		} else {
			for column := range record { //iterates over the column items for a given row
				df.DF[row-1][column] = newDataFloat(record[column]) //places csv data into DataFrame
			}
		}

	}

	return df

}
