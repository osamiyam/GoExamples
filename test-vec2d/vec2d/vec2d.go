package vec2d

import "fmt"
import "math"

type Vec struct {
	X float64
	Y float64
}

func (v Vec) String() string {
	return fmt.Sprintf("(%f, %f)", v.X, v.Y)
}

func (v Vec) decomp() (float64, float64) {
	return v.X, v.Y
}

func (v Vec) Length() float64 {
	return math.Sqrt(v.InnerProd(v))
}

func (v Vec) Add(v1 Vec) Vec {
	var x, y = v.decomp()
	var x1, y1 = v1.decomp()
	return Vec{x + x1, y + y1}
}

func (v Vec) InnerProd(v1 Vec) float64 {
	var x, y = v.decomp()
	var x1, y1 = v1.decomp()
	return x * x1 + y * y1
}

