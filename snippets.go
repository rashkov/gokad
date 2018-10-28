package snippets

func main() {
	var kbucket [][]byte = [][]byte{{0xAA, 0xBB, 0xCC}}
	o := node{kbucket: kbucket}
	fmt.Println("HI", o)

	var f uint8 = 0x01
	for i := 0; i < 7; f = f << 1 {
		i += 1
		fmt.Printf("%8b\n", f)
	}

	for i := 0; i < 8; f = f >> 1 {
		i += 1
		fmt.Printf("%8b\n", f)
	}
	var g = 5 << f
	fmt.Printf("%d %d\n", f, g)
	fmt.Printf("%x %x\n", f, g)
	fmt.Printf("%b %b\n", f, g)
}
