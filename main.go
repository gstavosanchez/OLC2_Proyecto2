package main

import (
	"fmt"
)

var P, H float64
var stack [30101999]float64
var heap [30101999]float64
var t0, t1, t2, t3, t4, t5, t6, t7, t8 float64

// -----------------------------------------------------------
// FUNCIONES NATIVAS
func getLenStr() {
	t3 = P + 1
	t4 = stack[int(t3)]
	t6 = heap[int(t4)]
	t5 = 0
L1:
	if t6 == -1 {
		goto L0
	}
	t4 = t4 + 1
	t6 = heap[int(t4)]
	t5 = t5 + 1
	goto L1
L0:
	stack[int(P)] = t5
	return
}

// ===========================================================
// MAIN
// ===========================================================
func main() {
	P = 0 // Puntero Stack
	H = 0 //Puntero Heap

	// -----------------------------------------------------------
	// ARREGLO
	t0 = H            // Apuntador donde se guarda el array
	t1 = H            // Contador del heap
	H = H + 6         //Almacenar espacio en heap longitud + array
	heap[int(t1)] = 5 // Guadar longtud arreglo
	t1 = t1 + 1
	// GUARDAR VALORES
	// PRIMITIVO: 1
	heap[int(t1)] = 1
	t1 = t1 + 1
	// PRIMITIVO: 2
	heap[int(t1)] = 2
	t1 = t1 + 1
	// PRIMITIVO: 3
	heap[int(t1)] = 3
	t1 = t1 + 1
	// PRIMITIVO: 4
	heap[int(t1)] = 4
	t1 = t1 + 1
	// PRIMITIVO: 5
	heap[int(t1)] = 5
	t1 = t1 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	// -----------------------------------------------------------
	// ASIGNACION: A
	stack[int(0)] = t0
	// FIN ASIGNACION DE A
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: HOLA

	t2 = H
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
	t7 = P + 1
	t7 = t7 + 1
	stack[int(t7)] = t2
	P = P + 1
	getLenStr()
	t8 = stack[int(P)]
	P = P - 1
	// -----------------------------------------------------------

	fmt.Printf("%d", int(t8))
	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

}
