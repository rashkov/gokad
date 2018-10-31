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
var N int = 32
var K int = 20
var A int = 3

type node struct {
	nodeid   []byte
	kbuckets [][][]byte
}

func (n node) new() node {
	// var a byte
	var a, b, c, d byte
	a = byte(rand.Intn(math.MaxUint8))
	b = byte(rand.Intn(math.MaxUint8))
	c = byte(rand.Intn(math.MaxUint8))
	d = byte(rand.Intn(math.MaxUint8))
	n.nodeid = []byte{a, b, c, d}
	n.kbuckets = make([][][]byte, N)
	return n
}

func (self node) kbucket_insert(neighbor_nodeid []byte) bool {
	kbucket := self.kbuckets[kxor(self.nodeid, neighbor_nodeid)]
	fmt.Println(len(kbucket), cap(kbucket))
	for i := 0; i < len(kbucket); i++ {
		fmt.Println(kbucket[i])
	}
	return true
}

func main() {
	//tests()
	// let's generate some nodes
	var neighborhood [10]node
	for i := 0; i < 10; i++ {
		n := node{}
		neighborhood[i] = n.new()
	}

	var m node = neighborhood[0]
	var n node = neighborhood[1]
	//fmt.Println("m,n,kxor", m, n, kxor(m.nodeid, n.nodeid))
	// Bootstrap m
	m.kbucket_insert(n.nodeid)
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
		//fmt.Println("address section", address_section)
		//fmt.Println("xoring", o, "^", y, "=", xored)
		if xored > 0 {
			differs = true
			break
		}
	}

	if !differs {
		return 0
	}

	var f uint8 = 1 << 7
	//fmt.Printf("F (binary): %b\n", f)
	var place int
	for place = 7; place >= 0; f = f >> 1 {
		// "Sieve" thing to get the bit # of first differing bit
		//fmt.Println("xored,f", xored, f)
		if xored >= f {
			break
		}
		place -= 1
	}

	//fmt.Println("address_section, place", address_section, place)
	kbucket_indx = address_section*8 + place
	return kbucket_indx
}

func tests() {
	var l []byte = []byte{0x0, 1, 3, 4}
	var m []byte = []byte{0x1, 1, 3, 4}
	var n []byte = []byte{0x2, 1, 3, 4}
	var o []byte = []byte{0x3, 1, 3, 4}
	var p []byte = []byte{0x4, 1, 3, 4}
	var q []byte = []byte{0x5, 1, 3, 4}
	fmt.Println("expect zero", l, m, kxor(l, m))
	fmt.Println("expect zero", m, m, kxor(m, m))
	fmt.Println("expect one", m, n, kxor(m, n))
	fmt.Println("expect one", n, o, kxor(n, o))
	fmt.Println("expect two", o, p, kxor(o, p))
	fmt.Println("expect zero", p, q, kxor(p, q))
	// fmt.Println("__TEST KXOR__:", "m: ", m, "n: ", n, kxor(m, n))
	// fmt.Println("__TEST KXOR__:", "o: ", o, "p: ", p, kxor(o, p))
}
