package main

import "fmt"
import "strings"
import "strconv"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ipAddr IPAddr) String() string {
	ipParts := make([] string, 4, 4)
	for index, val := range ipAddr {
		ipParts[index] = strconv.Itoa(int(val))
	}
	return strings.Join(ipParts, ".")
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
