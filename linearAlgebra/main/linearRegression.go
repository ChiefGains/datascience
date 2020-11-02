//Test of program that will take data points and calculate a linear regression line
//This version will just create and use slices for now, but eventually we'll be using the
//"Matrix" struct from the "matrices" package. It's just too much to figure out how to make that
//work while I'm also still just workshopping the concept
package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Line struct {
	X         []float64
	Y         []float64
	XAvg      float64
	YAvg      float64
	Slope     float64
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

func (l *Line) findRegression() *Line {
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

func (l *Line) predict(i int) *Line {
	x := float64(i)

	y := l.Intercept + (l.Slope * x)

	fmt.Println(y)

	return l
}

func main() {
	//fmt.Println("Just doing this so 'fmt' doesn't disappear when I save")

	line := makeLine(5)

	line.randFill()

	line.findRegression()

	fmt.Println(line)

	line.predict(4)
	line.predict(7)
	line.predict(-2)

}
