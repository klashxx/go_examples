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

type VertexPointer struct {
	X, Y float64
}

func (v *VertexPointer) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *VertexPointer) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// Methods can be associated with a named type
	// or a pointer to a named type.

	vp := &VertexPointer{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", vp, vp.Abs())
	vp.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", vp, vp.Abs())
}
