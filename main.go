package main

import (
	"fmt"
)

var P, H float64
var stack [30101999]float64
var heap [30101999]float64
var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10 float64
var t11, t12 float64

// ===========================================================
// MAIN
// ===========================================================
func main() {
	P = 0 // Puntero Stack
	H = 0 //Puntero Heap
	// -----------------------------------------------------------
	// INICIO FOR

	// -----------------------------------------------------------
	// ARREGLO
	t0 = H            // Apuntador donde se guarda el array
	t1 = H            // Contador del heap
	H = H + 6         //Almacenar espacio en heap longitud + array
	heap[int(t1)] = 5 // Guadar longtud arreglo
	t1 = t1 + 1
	// GUARDAR VALORES
	// PRIMITIVO: 100
	heap[int(t1)] = 100
	t1 = t1 + 1
	// PRIMITIVO: 2
	heap[int(t1)] = 2
	t1 = t1 + 1
	// PRIMITIVO: 4
	heap[int(t1)] = 4
	t1 = t1 + 1
	// PRIMITIVO: 44
	heap[int(t1)] = 44
	t1 = t1 + 1
	// PRIMITIVO: 200
	heap[int(t1)] = 200
	t1 = t1 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	t2 = P + 0
	stack[int(t2)] = t0
	t3 = t0
	t4 = stack[int(t2)]
	t5 = heap[int(t4)]
	t8 = 0
	t10 = t3
	t10 = t10 + 1
L0:
	t7 = P + t2
	t6 = stack[int(t7)]
	if t8 < t5 {
		goto L1
	}
	goto L2
L1:
	t9 = heap[int(t10)]
	stack[int(0)] = t9

	// -----------------------------------------------------------
	// INICIO PRINT
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: I
	t12 = P + 0 //Posicionar el puntero
	t11 = stack[int(t12)]
	// FIN ACCESO
	// -----------------------------------------------------------

	fmt.Printf("%d", int(t11))
	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------
	t8 = t8 + 1
	t10 = t10 + 1
	goto L0
L2:
	// FIN FOR
	// -----------------------------------------------------------

}
