// https://tour.golang.org/methods/22
// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (MyReader) Read(b []byte) (int, error) {
    b[0] = 'A'
    return 1, nil
}

func main() {
    reader.Validate(MyReader{})
}

