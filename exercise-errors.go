// https://tour.golang.org/methods/20
// Copy your Sqrt function from the earlier exercise and modify it to return an error value.
//
// Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.
//
// Create a new type
//
// type ErrNegativeSqrt float64
// and make it an error by giving it a
//
// func (e ErrNegativeSqrt) Error() string
// method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".
//
// Note: a call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop. You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?
//
// Change your Sqrt function to return an ErrNegativeSqrt value when given a negative number.
package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, ErrNegativeSqrt(x)
    }

    if x == 0 {
        return 0, nil
    }

    // Initial guess.
    //z := 1.0
    z := x / 2

    threshold := .00001
    max_iter := 10
    for i := 0; i != max_iter; i++ {
        // z -= (z*z - x) / (2*z)
        new_z := z - (z*z - x) / (2*z)

        diff := math.Abs(z - new_z)
        //fmt.Println(i, new_z, diff)
        if diff < threshold {
            break
        }
        z = new_z
    }
    return z, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}

