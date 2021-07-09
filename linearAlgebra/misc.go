package datascience

//obviously this will be renamed at some point, but right now it's just a catch-all for things that
//don't have a neat category yet

//LineIntercept takes two lines and determines where they intersect
func LineIntercept(a, b *Line) float64 {
	if a.Intercept == b.Intercept {
		return a.Intercept
	}
	return (a.Slope - b.Slope) / (a.Intercept - b.Intercept)
}
