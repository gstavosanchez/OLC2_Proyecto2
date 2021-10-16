package main

import "fmt"

var P, H float64
var stack [30101999]float64
var heap [30101999]float64
var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11 float64

// ===========================================================
// MAIN
// ===========================================================
func main() {
	P = 0 // Puntero Stack
	H = 0 //Puntero Heap
	// PRIMITIVO: 1
	// -----------------------------------------------------------
	// ASIGNACION: A
	stack[int(0)] = 1
	// FIN ASIGNACION
	// -----------------------------------------------------------

	// PRIMITIVO: 1
	// PRIMITIVO: 4
	// -----------------------------------------------------------
	// INICIO FOR
	t0 = P + 1
	stack[int(t0)] = 1

L0:
	t2 = P + 1 //Posicionar el puntero
	t1 = stack[int(t2)]
	if t1 > 4 {
		goto L1
	}
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: A
	t3 = stack[int(0)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: I
	t5 = P + 1 //Posicionar el puntero
	t4 = stack[int(t5)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t6 = t3 + t4

	// -----------------------------------------------------------
	// ASIGNACION: A
	t7 = P + 2
	stack[int(t7)] = t6
	// FIN ASIGNACION
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: A
	t9 = P + 2 //Posicionar el puntero
	t8 = stack[int(t9)]
	// FIN ACCESO
	// -----------------------------------------------------------

	fmt.Printf("%g", t8)
	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------
	goto L2

L2:
	t11 = P + 1 //Posicionar el puntero
	t10 = stack[int(t11)]
	t10 = t10 + 1
	stack[int(t11)] = t10
	goto L0
L1:
	// FIN FOR
	// -----------------------------------------------------------

}
