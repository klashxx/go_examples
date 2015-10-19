package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X int
	Y int
}

type Vertex2 struct {
	Lat, Long float64
}

var (
	v1  = Vertex{1, 2}  // has type Vertex
	v2  = Vertex{X: 1}  // Y:0 is implicit
	v3  = Vertex{}      // X:0 and Y:0
	p   = &Vertex{1, 2} // has type *Vertex
	pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	m   = map[string]Vertex2{
		"Bell Labs": Vertex2{
			40.68433, -74.39967,
		},
		// If the top-level type is just a type name,
		// you can omit it from the elements of the literal.
		"Google": {
			37.42202, -122.08408,
		},
	}
)

func main() {
	i, j := 42, 2701

	// A pointer holds the memory address of a variable.

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	fmt.Println(Vertex{1, 2})

	// Struct fields are accessed using a dot.
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// Struct fields can be accessed through a struct pointer
	z := &v
	z.X = 1e9
	fmt.Println(z)

	fmt.Println(v1, p, v2, v3)

	// The type [n]T is an array of n values of type T.

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// A slice points to an array of values and also includes a length.

	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}

	// Slicing slices
	fmt.Println("s ==", s)
	fmt.Println("s[1:4] ==", s[1:4])

	// missing low index implies 0
	fmt.Println("s[:3] ==", s[:3])

	// missing high index implies len(s)
	fmt.Println("s[4:] ==", s[4:])

	// Slices are created with the make function.
	// It works by allocating a zeroed array and returning
	// a slice that refers to that array

	a2 := make([]int, 5)
	printSlice("a2", a2)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

	// The zero value of a slice is nil.

	var z2 []int
	fmt.Println(z2, len(z2), cap(z2))
	if z2 == nil {
		fmt.Println("nil!")
	}

	// It is common to append new elements to a slice,
	// and so Go provides a built-in append function

	var a3 []int
	printSlice("a3", a3)

	// append works on nil slices.
	a3 = append(a3, 0)
	printSlice("a3", a3)

	// the slice grows as needed.
	a3 = append(a3, 1)
	printSlice("a3", a3)

	// we can add more than one element at a time.
	a3 = append(a3, 2, 3, 4)
	printSlice("a3", a3)

	// The range form of the for loop iterates over a slice or map.

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// You can skip the index or value by assigning to _.
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	// A map maps keys to values.
	// Maps must be created with make (not new) before use;
	// the nil map is empty and cannot be assigned to.

	// OR m = make(map[string]Vertex2)

	m["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(m)

	// Mutating Maps

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	val, ok := m["Answer"]
	fmt.Println("The value:", val, "Present?", ok)

	// Functions are values too.

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))

	// Function closures

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

/* Go functions may be closures.
A closure is a function value that references variables
from outside its body. The function may access and assign
to the referenced variables; in this sense the function is
"bound" to the variables.

For example, the adder function returns a closure.
Each closure is bound to its own sum variable. */

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
