package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// As in C or Java, you can leave the pre and
	// post statements empty
	// As in C or Java, you can leave the pre
	// and post statements empty.

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// The if statement looks as it does in C or Java,
	// except that the ( ) are gone and the { } are required.
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// A case body breaks automatically, unless it ends
	// with a fallthrough statement.
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// A defer statement defers the execution of a
	// function until the surrounding function returns.

	// The deferred call's arguments are evaluated immediately,
	// but the function call is not executed until the
	// surrounding function returns.

	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("counting")

	// Deferred function calls are pushed onto a stack.
	// When a function returns, its deferred calls are
	// executed in last-in-first-out order.
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {

	// Like for, the if statement can start with a short
	// statement to execute before the condition.
	// Variables declared by the statement are only
	// in scope until the end of the if.

	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// Variables declared inside an if short statement are
		// also available inside any of the else blocks.
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
