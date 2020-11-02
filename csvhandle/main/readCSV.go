//package created to place CSV data into structs
//to make the data more amenable to manipulation,
//specifically via the Matrices and linear algebra
//packages

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	//"strconv"
)

type myData interface{}

type DataFrame struct { //struct for holding CSV data

	X, Y int

	Columns []string

	Data [][]myData
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}

}

func MakeDataFrame(x, y int) *DataFrame { //creates an empty DataFrame from the dimensions given
	x-- //subtract number of rows by 1 for Column names

	a := make([]string, 0)   //make a slice for column names
	b := make([][]myData, x) //make slice of slices to hold data

	for i := range a {
		b[i] = make([]myData, y)
	}

	frame := &DataFrame{x, y, a, b} //create variable to hold dataframe

	return frame
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

func LoadCsv(file string) *DataFrame { //opens a csv file and stores the information in a string matrix

	data, err := os.Open(file)

	if err != nil {
		log.Fatalln("Unable to load csv", err)
	}

	x := findRows(data)    //finds number of rows
	data.Seek(0, 0)        //returns to beginning
	y := findColumns(data) //finds number of columns
	data.Seek(0, 0)

	df := MakeDataFrame(x, y) //creates a blank DataFrame to hold the data

	df.Columns = LabelColumns(data)

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
			continue //skip header
		}

		for column := range record { //iterates over the column items for a given row
			df.Data[row-1] = append(df.Data[row-1], record[column]) //places csv data into DataFrame
		}

	}

	return df
}

func main() {

	file := "src/learning/csvHandle/main/testData/weightData.csv"

	df := LoadCsv(file)

	fmt.Println(df.Columns)
	fmt.Println(df.Data)

}
