package main

import (
	"fmt"
)

var P, H float64
var stack [30101999]float64
var heap [30101999]float64
var t0, t1, t2, t3, t4 float64

// ===========================================================
// MAIN
// ===========================================================
func main() {
	P = 0 // Puntero Stack
	H = 0 //Puntero Heap

	// -----------------------------------------------------------
	// DECLARAR STRUCT PERSONA()
	t0 = H
	t1 = t0
	H = H + 2 //Almacenar espacio en heap del struct Persona
	// PRIMITIVO: 23
	heap[int(t1)] = 23
	t1 = t1 + 1
	// PRIMITIVO: GUS

	t2 = H
	heap[int(H)] = 71 // G
	H = H + 1
	heap[int(H)] = 117 // u
	H = H + 1
	heap[int(H)] = 115 // s
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t1)] = t2
	t1 = t1 + 1
	// -----------------------------------------------------------
	// ASIGNACION: A
	stack[int(0)] = t0
	// FIN ASIGNACION DE A
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	t3 = stack[int(0)]
	t3 = t3 + 0
	t4 = heap[int(t3)]
	fmt.Printf("%g", t4)
	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

}
