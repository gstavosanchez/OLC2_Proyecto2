package main

import (
	"fmt"
)

var P, H float64
var stack [30101999]float64
var heap [30101999]float64
var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11 float64

// -----------------------------------------------------------
// FUNCIONES NATIVAS
func printStr() {
	t1 = P + 1 // Puntero del parametro
	t2 = stack[int(t1)]
L1:
	t3 = heap[int(t2)]
	if t3 == -1 {
		goto L0
	}
	fmt.Printf("%c", int(t3))
	t2 = t2 + 1 //aumentar el contador del heap
	goto L1
L0:
	return
}
func divValidate() {
	t6 = P + 1
	t7 = stack[int(t6)]
	t6 = t6 + 1
	t8 = stack[int(t6)]
	if t8 != 0 {
		goto L3
	}
	fmt.Printf("%c", int(77))
	fmt.Printf("%c", int(97))
	fmt.Printf("%c", int(116))
	fmt.Printf("%c", int(104))
	fmt.Printf("%c", int(69))
	fmt.Printf("%c", int(114))
	fmt.Printf("%c", int(114))
	fmt.Printf("%c", int(111))
	fmt.Printf("%c", int(114))
	fmt.Printf("%c", int(10))
	t9 = 0 // resultado incorrecto
	goto L2
L3:
	t9 = t7 / t8
L2:
	stack[int(P)] = t9 // gurdar resultado
	return
}

// ===========================================================
// MAIN
// ===========================================================
func main() {
	P = 0 // Puntero Stack
	H = 0 //Puntero Heap

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: G

	t0 = H
	heap[int(H)] = 103 // g
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t4 = P + 0
	t4 = t4 + 1
	stack[int(t4)] = t0
	// -----------------------------------------------------------

	P = P + 0
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t5 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 0

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: 5
	// PRIMITIVO: 0
	// -----------------------------------------------------------
	// PASO DE PARAMETROS
	t10 = P + 0
	// 1ER PARAMETRO
	t10 = t10 + 1
	stack[int(t10)] = 5
	// 2DO PARAMETRO
	t10 = t10 + 1
	stack[int(t10)] = 0
	// FIN PASO DE PARAMETROS
	// -----------------------------------------------------------
	// CAMBIO DE ENTORNO
	P = P + 0
	divValidate()
	// GUARDAR EL RETURN DE LA FUNCION
	t11 = stack[int(P)]
	P = P - 0

	fmt.Printf("%g", t11)
	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

}
