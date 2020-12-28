/*Package regression Version 0.1 of my regression package. So far just does some very basic linear
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
	//Xpoints []float64 (would be just to have a list for iterating through)
	X         []float64 //maybe make this a map? X  map[string][]float64 = X[x1: [4, 2, 8], x2: [3, 7, 20]] etc
	Y         []float64
	XAvg      float64 //XAvg map[string]float64 = XAvg[x1: 4.67, x2: 10]
	YAvg      float64
	Slope     float64 //I *think* for multiple regression, we're going to need a different slope for each dimension, yeah?
	Intercept float64
	Length    float64
}

func makeLine(x int) *Line {
	a := make([]float64, x)
	b := make([]float64, x)

	var c float64
	var d float64
	var e float64
	var f float64

	g := float64(x)

	line := &Line{a, b, c, d, e, f, g}

	return line

}

func (l *Line) randFill() *Line {
	for i := range l.X {
		l.X[i] = float64(rand.Intn(5))
		l.Y[i] = float64(rand.Intn(20))
	}

	var a float64
	var b float64
	a, b = 0, 0

	for i := range l.X {
		a += l.X[i]
		b += l.Y[i]
	}

	a = a / l.Length
	b = b / l.Length

	l.XAvg = a
	l.YAvg = b

	return l

}

func (l *Line) findRegression() *Line { //should really be something like "fitLine", but whatevs
	var num float64
	var den float64

	num, den = 0, 0

	for i := range l.X {
		a := l.X[i] - l.XAvg
		b := l.Y[i] - l.YAvg

		num += a * b
		den += math.Pow(a, 2)
	}

	l.Slope = num / den

	l.Intercept = l.YAvg - (l.Slope * l.XAvg)

	return l
}

/*func (l *Line) multipleFit() {

	for _, i := range l.X{
		var numerator, denominator float64 = 0, 0
		for j := range l.X[i]{
			a := l.X[i] - l.XAvg[i]
			b := l.Y[i] - l.YAvg

			numerator += a * b
			denominator += math.Pow(a, 2)
		}
	}
}
*/

func (l *Line) predict(i int) *Line {
	x := float64(i)

	y := l.Intercept + (l.Slope * x)

	fmt.Println(y)

	return l
}
