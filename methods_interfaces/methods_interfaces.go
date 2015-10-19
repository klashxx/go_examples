package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Go does not have classes. However, you can
// define methods on struct types.

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// You can declare a method on any type that is
// declared in your package, not just struct types.

// However, you cannot define a method on a type
// from another package (including built in types).

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
