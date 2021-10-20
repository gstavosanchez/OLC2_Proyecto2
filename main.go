package main

import (
	"fmt"
)

var P, H float64
var stack [30101999]float64
var heap [30101999]float64
var t0, t1, t2, t3, t4, t5, t6 float64

// -----------------------------------------------------------
// FUNCIONES
func sumar() {
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: A
	t1 = P + 1 //Posicionar el puntero
	t0 = stack[int(t1)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: B
	t3 = P + 2 //Posicionar el puntero
	t2 = stack[int(t3)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t4 = t0 + t2

	stack[int(P)] = t4
	goto L0
L0:
	return
}

// ===========================================================
// MAIN
// ===========================================================
func main() {
	P = 0 // Puntero Stack
	H = 0 //Puntero Heap

	// -----------------------------------------------------------
	// LLAMADO A FUNCION: SUMAR
	// PRIMITIVO: 5
	// PRIMITIVO: 5
	t5 = P + 1
	stack[int(t5)] = 5
	t5 = t5 + 1
	stack[int(t5)] = 5
	P = P + 0
	sumar()
	t5 = stack[int(P)]
	P = P - 0
	// FIN LLAMADO A FUNCION
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ASIGNACION: A
	stack[int(0)] = t5
	// FIN ASIGNACION
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: A
	t6 = stack[int(0)]
	// FIN ACCESO
	// -----------------------------------------------------------

	fmt.Printf("%g", t6)
	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

}
