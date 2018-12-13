// https://tour.golang.org/methods/18
// Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.
//
// For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".
package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ipaddr IPAddr) String() string {
    return fmt.Sprintf("%v.%v.%v.%v", ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3])
}

func main() {
    hosts := map[string]IPAddr{
        "loopback":  {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    for name, ip := range hosts {
        fmt.Printf("%v: %v\n", name, ip)
    }
}

