/*Package regression Version 0.2 of my regression package. So far just does some very basic linear
regression,and isn't even able to auto-fill data. Eventually, I'll be able to load a CSV and take a
column of the dataframe, and temporarily convert it into a line for use in this package. Hopefully
this package will also eventually be able to work in tandem with the Matrices package. Would also
like to be able to marshal and unmarshal json, which should be at least as easy as csv, if not
significantly easier.*/
package regression

import (
	"fmt"
	"math"
	"math/rand"
)

//Line defines a fit line and its values.
//One weakness here is that this line
//cannot exist in n number of dimensions.
//We'll work on that.
type Line struct {
	Dimensions int
	X          [][]float64 //keeps a slice of x values for each dimension
	Y          []float64
	XAvg       []float64 //indexed slice of x averages for each dimesion
	YAvg       float64
	Slope      []float64 //I *think* we need a different slope for each dimension?
	Intercept  float64
	Length     float64 //number of values, not true distance of line, which is theoretically infinite
}

func makeLine(dimensions, points int) *Line {
	x := make([][]float64, dimensions)

	for i := 0; i < points; i++ {
		x[i] = make([]float64, points)
	}

	y := make([]float64, points)

	xavg := make([]float64, dimensions)
	var yavg float64
	slope := make([]float64, dimensions)
	var intercept float64
	length := float64(points)

	line := &Line{dimensions, x, y, xavg, yavg, slope, intercept, length}

	return line

}

//randFill fills a line with random points, for testing purposes
func (l *Line) randFill(limit int) {
	for i := range l.X {
		for j := range l.X[i] {
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
func (l *Line) fitLine() {
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

		l.Slope[i] = num / den
	}

	//l.Intercept = l.YAvg - (l.Slope * l.XAvg) *I DON'T KNOW HOW TO DO THIS FOR MULTIPLE REGRESSION*

}

/*predict still needs updated to handle multiple, don't know how to do that yet*/
//predict looks at a fit line, looks at an input variable, and predicts the outcome based on the fit line
func (l *Line) predict(i int) float64 {
	x := float64(i)

	y := l.Intercept + (l.Slope * x)

	fmt.Println(y)

	return y
}
