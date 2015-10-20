package main

import (
	"fmt"
	"math"
	"os"
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

// An interface type is defined by a set of methods
type Abser interface {
	Abs() float64
}

// A value of interface type can hold any
// value that implements those methods.

func (v *VertexPointer) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *VertexPointer) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Reader-Writer example
// A type implements an interface by implementing the
// methods. There is no explicit declaration of intent.
// NO "implements" keyword.

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
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

	var a Abser

	a = f // a MyFloat implements Abser
	a = v // a *Vertex implements Abser

	fmt.Println(a.Abs())

	var w Writer

	// os.Stdout implements Writer
	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}
