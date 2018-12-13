// https://tour.golang.org/methods/23
// A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.
//
// For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).
//
// Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.
//
// The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.
package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (reader *rot13Reader) Read(b []byte) (int, error) {
	n, err := reader.r.Read(b)
    for i := 0; i != n; i++ {
		c := int(b[i])

        offset := 13
        a := 97
        z := a + 25
        A := 65
        Z := A + 25
		if c >= a && c <= z {
			c += offset
			dist := c - z
			if dist > 0 {
			    c = a + dist - 1
			}
		} else if c >= A && c <= Z {
			c += offset
			dist := c - Z
			if dist > 0 {
			    c = A + dist - 1
			}
		}

	    b[i] = byte(c)
	}
	return n, err
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}

