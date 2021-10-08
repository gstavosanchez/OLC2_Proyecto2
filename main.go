package main

import "fmt"

var stack [30101999]float64
var heap [30101999]float64
var t1, t2, t3, t4 float32

func main() {
	// -----------------------------------------------------------
	// INICIO EXPRESION LOGICA
	// -----------------------------------------------------------
	// INICIO EXPRESION LOGICA
	// -----------------------------------------------------------
	// INICIO EXPRESION RELACIONAL
	if 55 == 55 {
		goto L1
	}
	goto L2
	// FIN EXPRESION RELACIONAL
	// -----------------------------------------------------------
	// FIN EXPRESION LOGICA
	// -----------------------------------------------------------
L2:
	// -----------------------------------------------------------
	// INICIO EXPRESION RELACIONAL
	if 10 > 2 {
		goto L0
	}
	goto L1
	// FIN EXPRESION RELACIONAL
	// -----------------------------------------------------------
	// FIN EXPRESION LOGICA
	// -----------------------------------------------------------
L0:
	fmt.Printf("%c", int(116))
	fmt.Printf("%c", int(114))
	fmt.Printf("%c", int(117))
	fmt.Printf("%c", int(101))
	goto L3
L1:
	fmt.Printf("%c", int(102))
	fmt.Printf("%c", int(97))
	fmt.Printf("%c", int(108))
	fmt.Printf("%c", int(115))
	fmt.Printf("%c", int(101))
L3:
	fmt.Printf("%c", int(32))
	fmt.Printf("%c", int(10))

}
