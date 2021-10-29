package main

import (
	"fmt"
)

var P, H float64
var stack [30101999]float64
var heap [30101999]float64
var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10 float64
var t11, t12 float64

// -----------------------------------------------------------
// FUNCIONES
func suma() {
	// -----------------------------------------------------------
	// INICIO IF
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: N
	t1 = P + 1 //Posicionar el puntero
	t0 = stack[int(t1)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// PRIMITIVO: 1
	if t0 == 1 {
		goto L1
	}
	goto L2
L1:
	// PRIMITIVO: 1
	stack[int(P)] = 1
	goto L0
	goto L3
L2:
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: N
	t3 = P + 1 //Posicionar el puntero
	t2 = stack[int(t3)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// LLAMADO A FUNCION: SUMA
	// GUARDAR TEMPORALES

	t4 = P + 2
	stack[int(t4)] = t0
	t4 = t4 + 1
	stack[int(t4)] = t2

	// FIN DE GUARDAR TEMPORALES
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: N
	t6 = P + 1 //Posicionar el puntero
	t5 = stack[int(t6)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// PRIMITIVO: 1
	t7 = t5 - 1

	t8 = P + 5
	stack[int(t8)] = t7
	P = P + 4
	suma()
	t8 = stack[int(P)]
	P = P - 4
	// RECUPERAR TEMPORALES

	t9 = P + 2
	t0 = stack[int(t9)]
	t9 = t9 + 1
	t2 = stack[int(t9)]
	t9 = t9 + 1
	t5 = stack[int(t9)]
	t9 = t9 + 1
	t7 = stack[int(t9)]

	// FIN DE RECUPERACION
	// FIN LLAMADO A FUNCION
	// -----------------------------------------------------------

	t10 = t2 + t8

	stack[int(P)] = t10
	goto L0
L3:
	// FIN IF
	// -----------------------------------------------------------

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
	// LLAMADO A FUNCION: SUMA
	// PRIMITIVO: 10
	t11 = P + 1
	stack[int(t11)] = 10
	P = P + 0
	suma()
	t11 = stack[int(P)]
	P = P - 0
	// FIN LLAMADO A FUNCION
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ASIGNACION: RES
	stack[int(0)] = t11
	// FIN ASIGNACION DE RES
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: RES
	t12 = stack[int(0)]
	// FIN ACCESO
	// -----------------------------------------------------------

	fmt.Printf("%d", int(t12))
	// FIN PRINT
	// -----------------------------------------------------------

}
