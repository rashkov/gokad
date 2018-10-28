package main

import (
	_ "bytes"
	"fmt"
	"math"
	"math/rand"
)

// If we have 8 bits of address space,
//   then we have an array of 8 kbuckets,
//   one for each bit of difference.
// N is the address space in bits
// K is the max length of a k-bucket
// A is the # of simultaneous async requests to send
var N int = 8
var K int = 10
var A int = 3

type node struct {
	nodeid  []byte
	kbucket [][]byte
}

func (n node) new() node {
	// var a byte
	var a, b, c, d byte
	a = byte(rand.Intn(math.MaxUint8))
	b = byte(rand.Intn(math.MaxUint8))
	c = byte(rand.Intn(math.MaxUint8))
	d = byte(rand.Intn(math.MaxUint8))
	n.nodeid = []byte{a, b, c, d}
	return n
}

func main() {
	// let's generate some nodes
	var neighborhood [10]node
	for i := 0; i < 10; i++ {
		n := node{}
		neighborhood[i] = n.new()
	}
	fmt.Println(neighborhood)

	var f uint8 = 0x01
	var bucket_indx_map [8]uint8
	for i := 0; i < 7; f = f << 1 {
		i += 1
		fmt.Printf("%8b\n", f)
		bucket_indx_map[i] = f
	}
	fmt.Println(bucket_indx_map)

	var m []byte = neighborhood[0].nodeid
	var n []byte = neighborhood[2].nodeid
	fmt.Println(kxor(m, n))
	// for i := 1; i < 10; i++ {
	// 	m[i]
	// }
}

func kxor(m []byte, n []byte) (kbucket_indx uint8, place uint8) {
	fmt.Println(m, n)
	fmt.Println(len(m), len(n))
	for i := 0; i < len(m); i++ {
		o := m[i]
		y := n[i]
		// fmt.Println("o:", o)
		// for j := 0; j < len(n); j++ {
		// y := n[j]
		// fmt.Println("y:", y)
		xr := o ^ y
		if xr > 0 {
			fmt.Println(xr)
		}
		// }
	}
	return 0, 0
}
