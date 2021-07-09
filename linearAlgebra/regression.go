/*Package regression Version 0.2 of my regression package. So far just does some very basic linear
regression,and isn't even able to auto-fill data. Eventually, I'll be able to load a CSV and take a
column of the dataframe, and temporarily convert it into a line for use in this package. Hopefully
this package will also eventually be able to work in tandem with the Matrices package. Would also
like to be able to marshal and unmarshal json, which should be at least as easy as csv, if not
significantly easier.*/
package datascience

import (
	"fmt"
	"math"
	"math/rand"
)

//Line defines a fit line and its values.
//Working on making a line work in n number of dimensions
//Also, should dimensional attributes be kept in a map in order to label them? like instead of just
//"dimension 3", do we need things like "heigh" and "weight" for easier manipulation?
type Line struct {
	Dimensions int
	X          [][]float64 //keeps a slice of x values for each dimension. Should be indexed in order of dimension
	Y          []float64
	XAvg       []float64 //indexed slice of x averages for each dimesion
	YAvg       float64
	Slope      float64 //Slope is the same in every dimension, right? Gotta double check that
	Intercept  float64
	Length     float64 //number of values, not true distance of line, which is theoretically infinite
}

//MakeLine an integer "dimensions" as an input and returns a pointer to a line
func MakeLine(dimensions, points int) *Line {
	x := make([][]float64, dimensions)

	for i := 0; i < points; i++ {
		x[i] = make([]float64, points)
	}

	y := make([]float64, points)

	xavg := make([]float64, dimensions)
	var yavg float64
	var slope float64
	var intercept float64
	length := float64(points)

	line := &Line{dimensions, x, y, xavg, yavg, slope, intercept, length}

	return line

}

//randFill fills a line with random points, for testing purposes
func (l *Line) RandFill(limit int) {
	for i := range l.X {
		for j := range l.X[i] { //need to update this, obviously, since it's broken
			l.X[i][j] = float64(rand.Intn(limit))
		}
		l.Y[i] = float64(rand.Intn(limit))
	}

	var ySum float64 //I should probably add xSum and ySum to the line struct, but that sounds like a headache
	ySum = 0

	for i := range l.X {
		var xSum float64
		xSum = 0

		for j := range l.X[i] {
			xSum += l.X[i][j]
		}

		l.XAvg[i] = xSum / l.Length

		ySum += l.Y[i]
	}

	l.YAvg = ySum / l.Length
}

/*fitLine like *super* needs updated to handle multiple regression. Should
be as simple as an iterative statement going through and doing least squares
for each dimension*/

//fitLine takes data points and fits them to a regression line
func (l *Line) FitLine() {
	for i := range l.X {
		var num float64
		var den float64

		num, den = 0, 0

		for j := range l.X {
			a := l.X[i][j] - l.XAvg[i]
			b := l.Y[j] - l.YAvg

			num += a * b
			den += math.Pow(a, 2)
		}

		l.Slope = num / den //this is useless inside of the loop. I'm still unclear if slope is constant in all dimensions
	}

	//l.Intercept = l.YAvg - (l.Slope * l.XAvg) *I DON'T KNOW HOW TO DO THIS FOR MULTIPLE REGRESSION*

}

//THIS FUNCTION IS BEING UPDATED TO HANDLE MULTIPLE REGRESSION. WE APPRECIATE YOUR PATIENCE
//predict looks at a fit line, looks at an input variable, and predicts the outcome based on the fit line
func (l *Line) Predict(variables ...int) []float64 {
	res := make([]float64, len(variables))
	for _, d := range variables {
		x := float64(d)

		res = append(res, l.Intercept+(l.Slope*x))

		fmt.Println(res)
	}

	return res
}
