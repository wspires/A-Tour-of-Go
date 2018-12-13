// https://tour.golang.org/moretypes/18
/*
    Implement Pic.
    It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers.
    When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

    The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

    (You need to use a loop to allocate each []uint8 inside the [][]uint8.)

    (Use uint8(intValue) to convert between types.)
 */
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    //var p := [dx][dy]int8;
    p := make([][]uint8, dx, dx)
    for x := range p {
        p[x] = make([]uint8, dy, dy)
        for y := range p[x] {
            p[x][y] = uint8((x + y) / 2)
            //p[x][y] = uint8(x * y)
            p[x][y] = uint8(x ^ y)
        }
    }
    return p;
}

func main() {
    pic.Show(Pic)
}

