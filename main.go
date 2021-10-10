package main

import "fmt"

var P, H float64
var stack [30101999]float64
var heap [30101999]float64
var t0, t1, t2, t3, t4, t5, t6, t7, t8 float64

// -----------------------------------------------------------
// FUNCIONES NATIVAS
func printStr() {
	t1 = P + 1 // Puntero del stack // parametro 1
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

// -----------------------------------------------------------
func concacString() {

}

func main() {
	P = 0 // Puntero Stack
	H = 0 //Puntero Heap

	t0 = H
	heap[int(H)] = 104 // h
	H = H + 1
	heap[int(H)] = 111 // o
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 97 // a
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

	t6 = H
	heap[int(H)] = 109 // m
	H = H + 1
	heap[int(H)] = 117 // u
	H = H + 1
	heap[int(H)] = 110 // n
	H = H + 1
	heap[int(H)] = 100 // d
	H = H + 1
	heap[int(H)] = 111 // o
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t7 = P + 0
	t7 = t7 + 1
	stack[int(t7)] = t6
	// -----------------------------------------------------------

	P = P + 0
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t8 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 0

	fmt.Printf("%c", int(10))

}
