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

type address []byte

type node struct {
	nodeid   address
	kbuckets [][]address
}

func (n node) new() node {
	// var a byte
	var a, b, c, d byte
	a = byte(rand.Intn(math.MaxUint8))
	b = byte(rand.Intn(math.MaxUint8))
	c = byte(rand.Intn(math.MaxUint8))
	d = byte(rand.Intn(math.MaxUint8))
	n.nodeid = address{a, b, c, d}
	n.kbuckets = make([][]address, N)
	fmt.Println("printed from A", n.nodeid)
	return n
}

func (self node) kbucket_insert(neighbor_nodeid address) bool {
	kdistance := kxor(self.nodeid, neighbor_nodeid)
	kbucket := self.kbuckets[kdistance]

	if len(kbucket) <= 20 {
		self.kbuckets[kdistance] = append(kbucket, neighbor_nodeid)
		kbucket = self.kbuckets[kdistance]
		return true
	} else {
		// Out of space, handle this as an LRU & keep a cache of substitute nodes ready
		return false
	}
}

func (self node) find_node(target_nodeid address) bool {
	fmt.Println(self.get_k_nearest(target_nodeid))
	return true
}

func (self node) get_k_nearest(target_nodeid address) []address {
	var k_nearest []address
	distance_to_target := kxor(self.nodeid, target_nodeid) + 1
	index := distance_to_target
	lower_bound := distance_to_target - 1
	upper_bound := distance_to_target + 1

	for _, b := range self.kbuckets[index] {
		if b != nil {
			k_nearest = append(k_nearest, b)
		}
	}
	for len(k_nearest) < 20 && (lower_bound >= 0 || upper_bound <= 19) {
		if lower_bound >= 0 {
			for _, b := range self.kbuckets[lower_bound] {
				if b != nil && len(k_nearest) < 20 {
					k_nearest = append(k_nearest, b)
				}
			}
		}
		if upper_bound < 20 {
			for _, b := range self.kbuckets[upper_bound] {
				if b != nil && len(k_nearest) < 20 {
					k_nearest = append(k_nearest, b)
				}
			}
		}
		lower_bound--
		upper_bound++
	}
	return k_nearest
}

func main() {
	//tests()
	// let's generate some nodes
	var neighborhood [10]node
	for i := 0; i < 10; i++ {

		// TODO: Figure out why the following weirdness happens.
		//       Probably has to do with lexical scope / closure
		// NOTE: When done this way, output looks like this
		// printed from A [86 132 122 254]
		// printed from B []
		// printed from A [151 153 115 155]
		// printed from B []
		// n := node{}
		// neighborhood[i] = n.new()

		// NOTE: When done this way, output looks like this:
		// printed from A [86 132 122 254]
		// printed from B [86 132 122 254]
		// printed from A [151 153 115 155]
		// printed from B [151 153 115 155]
		n := node{}.new()
		neighborhood[i] = n

		fmt.Println("printed from B", n.nodeid)
	}

	var m node = neighborhood[0]
	var n node = neighborhood[1]
	// Bootstrap m
	m.kbucket_insert(n.nodeid)
	n.kbucket_insert(m.nodeid)
	n.find_node(m.nodeid) // have m search for itself
}

// kxor returns p, the place of the binary digit where their nodeids differ
// node m files node n into place p of its kbucket array
func kxor(m address, n address) (kbucket_indx int) {
	var address_section int
	var xored uint8
	var differs = false
	for address_section = 0; address_section < len(m); address_section++ {
		o := m[address_section]
		y := n[address_section]
		xored = o ^ y
		if xored > 0 {
			differs = true
			break
		}
	}

	if !differs {
		return 0
	}

	var f uint8 = 1 << 7
	var place int
	for place = 7; place >= 0; f = f >> 1 {
		// "Sieve" thing to get the bit # of first differing bit
		if xored >= f {
			break
		}
		place -= 1
	}

	kbucket_indx = address_section*8 + place
	return kbucket_indx
}

func tests() {
	var l address = address{0x0, 1, 3, 4}
	var m address = address{0x1, 1, 3, 4}
	var n address = address{0x2, 1, 3, 4}
	var o address = address{0x3, 1, 3, 4}
	var p address = address{0x4, 1, 3, 4}
	var q address = address{0x5, 1, 3, 4}
	fmt.Println("expect zero", l, m, kxor(l, m))
	fmt.Println("expect zero", m, m, kxor(m, m))
	fmt.Println("expect one", m, n, kxor(m, n))
	fmt.Println("expect one", n, o, kxor(n, o))
	fmt.Println("expect two", o, p, kxor(o, p))
	fmt.Println("expect zero", p, q, kxor(p, q))
}
