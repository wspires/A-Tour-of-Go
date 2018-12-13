// https://tour.golang.org/moretypes/26
// Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
/*
func fibonacci() func() int {
    fib := 0
    prev_fib := 0
    return func() int {
        saved_fib := fib
        fib += prev_fib
		prev_fib = saved_fib
        if fib == 0 {
		    prev_fib = 1
		}
		return fib
    }
}
*/

func fibonacci() func() int {
    x, y := 0, 1
    return func() int {
        fib := x
        x, y = y, x + y
        return fib
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}

