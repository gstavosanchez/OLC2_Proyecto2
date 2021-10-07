package main

import "fmt"

var t0, t1 float64

func main() {
	// -----------------------------------------------------------
	// INICIO EXPRESION RELACIONAL
	// -----------------------------------------------------------
	goto L0
	goto L1
L0:
	t0 = 1
	goto L2
L1:
	t0 = 0
L2:
	goto L3
	goto L4
L3:
	t1 = 1
	goto L5
L4:
	t1 = 0
L5:
	if t0 == t1 {
		goto L6
	}
	goto L7
	// -----------------------------------------------------------
	// FIN EXPRESION RELACIONAL
	// -----------------------------------------------------------
L6:
	fmt.Printf("%c", int(116))
	fmt.Printf("%c", int(114))
	fmt.Printf("%c", int(117))
	fmt.Printf("%c", int(101))
	goto L8
L7:
	fmt.Printf("%c", int(102))
	fmt.Printf("%c", int(97))
	fmt.Printf("%c", int(108))
	fmt.Printf("%c", int(115))
	fmt.Printf("%c", int(101))
L8:
	fmt.Printf("%c", int(32))
	fmt.Printf("%c", int(10))

}
