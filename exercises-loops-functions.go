// https://tour.golang.org/flowcontrol/8
package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    if x == 0 {
        return 0
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
        fmt.Println(i, new_z, diff)
        if diff < threshold {
            break
        }
        z = new_z
    }
    return z
}

func main() {
    for i := 0; i != 10; i++ {
        x := float64(i)
        sqrt_x := Sqrt(x)
        math_sqrt_x := math.Sqrt(x)
        fmt.Println(x, sqrt_x, math_sqrt_x, sqrt_x - math_sqrt_x)
        fmt.Println("===")
    }
}

