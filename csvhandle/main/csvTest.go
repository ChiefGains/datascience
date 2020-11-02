//test file to use the csv package

package main

import (
	"fmt"
	"learning/csvhandle/csvmat"
)

func main() {
	file := "src/learning/csvHandle/main/testData/weightData.csv"

	df := csvmat.LoadCsv(file)

	fmt.Println(df.Columns)
	fmt.Println(df.DF)
}
