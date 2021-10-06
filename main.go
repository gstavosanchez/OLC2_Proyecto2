package main

import "fmt"

// var t0 float64

func main() {
	goto L0
	goto L1
L0:
	fmt.Printf("%c", 116) // T
	fmt.Printf("%c", 114) // R
	fmt.Printf("%c", 117) // U
	fmt.Printf("%c", 101) // E
	goto L2
L1:
	fmt.Printf("%c", 116) // F
	fmt.Printf("%c", 97)  // A
	fmt.Printf("%c", 108) // L
	fmt.Printf("%c", 115) // S
	fmt.Printf("%c", 101) // E
L2:
	fmt.Printf("%c", 10) // '\n'

}
