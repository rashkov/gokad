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
	tests()
	// let's generate some nodes
	var neighborhood [10]node
	for i := 0; i < 10; i++ {
		n := node{}
		neighborhood[i] = n.new()
	}

	// var m []byte = neighborhood[0].nodeid
	// var n []byte = neighborhood[2].nodeid
	//fmt.Println("m,n,kxor", m, n, kxor(m, n))
}

// kxor returns p, the place of the binary digit where their nodeids differ
// node m files node n into place p of its kbucket array
func kxor(m []byte, n []byte) (kbucket_indx int) {
	var address_section int
	var xored uint8
	var differs = false
	for address_section = 0; address_section < len(m); address_section++ {
		o := m[address_section]
		y := n[address_section]
		xored = o ^ y
		fmt.Println("address section", address_section)
		fmt.Println("xoring", o, "^", y, "=", xored)
		if xored > 0 {
			differs = true
			break
		}
	}

	if !differs {
		return 1
	}

	var f uint8 = 0xFF
	var place int
	for place = 0; place < 7; f = f >> 1 {
		// "Sieve" thing to get the bit # of first differing bit
		//fmt.Println("xored,f", xored, f)
		if xored > f {
			place -= 1 // Went too far, shift back one
			break
		}
		place += 1
	}

	fmt.Println("address_section, place", address_section, place)
	kbucket_indx = address_section*8 + place
	return kbucket_indx
}

func tests() {
	var m []byte = []byte{0x1, 1, 3, 4}
	var n []byte = []byte{0x2, 1, 3, 4}
	// var o []byte = []byte{0x3, 1, 3, 4}
	// var p []byte = []byte{0x4, 1, 3, 4}
	// fmt.Println("expect one", m, m, kxor(m, m))
	fmt.Println("expect one", m, n, kxor(m, n))
	// fmt.Println("expect two", n, o, kxor(n, o))
	// fmt.Println("expect two", o, p, kxor(o, p))
	// fmt.Println("__TEST KXOR__:", "m: ", m, "n: ", n, kxor(m, n))
	// fmt.Println("__TEST KXOR__:", "o: ", o, "p: ", p, kxor(o, p))
}
