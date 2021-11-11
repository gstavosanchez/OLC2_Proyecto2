package main

import (
	"fmt"
)

var P, H float64
var stack [30101999]float64
var heap [30101999]float64
var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10 float64
var t11, t12, t13, t14, t15, t16 float64

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
	H = H + 4         //Almacenar espacio en heap longitud + array
	heap[int(t1)] = 3 // Guadar longtud arreglo
	t1 = t1 + 1
	// GUARDAR VALORES
	// PRIMITIVO: EL PADRINO

	t2 = H
	heap[int(H)] = 69 // E
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 80 // P
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 100 // d
	H = H + 1
	heap[int(H)] = 114 // r
	H = H + 1
	heap[int(H)] = 105 // i
	H = H + 1
	heap[int(H)] = 110 // n
	H = H + 1
	heap[int(H)] = 111 // o
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t1)] = t2
	t1 = t1 + 1
	// PRIMITIVO: EL PADRINO 2

	t3 = H
	heap[int(H)] = 69 // E
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 80 // P
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 100 // d
	H = H + 1
	heap[int(H)] = 114 // r
	H = H + 1
	heap[int(H)] = 105 // i
	H = H + 1
	heap[int(H)] = 110 // n
	H = H + 1
	heap[int(H)] = 111 // o
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 50 // 2
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t1)] = t3
	t1 = t1 + 1
	// PRIMITIVO: EL PADRINO 3

	t4 = H
	heap[int(H)] = 69 // E
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 80 // P
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 100 // d
	H = H + 1
	heap[int(H)] = 114 // r
	H = H + 1
	heap[int(H)] = 105 // i
	H = H + 1
	heap[int(H)] = 110 // n
	H = H + 1
	heap[int(H)] = 111 // o
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 51 // 3
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t1)] = t4
	t1 = t1 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	// -----------------------------------------------------------
	// ASIGNACION: CINEPOLIS
	stack[int(0)] = t0
	// FIN ASIGNACION DE CINEPOLIS
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO FOR
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: CINEPOLIS
	t5 = stack[int(0)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t6 = P + 1
	stack[int(t6)] = t5
	t7 = t5
	t8 = stack[int(t6)]
	t9 = heap[int(t8)]
	t12 = 0
	t14 = t7
	t14 = t14 + 1
L0:
	t11 = P + t6
	t10 = stack[int(t11)]
	if t12 < t9 {
		goto L1
	}
	goto L2
L1:
	t13 = heap[int(t14)]
	stack[int(1)] = t13

	// -----------------------------------------------------------
	// INICIO PRINT
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: PELICULA
	t16 = P + 1 //Posicionar el puntero
	t15 = stack[int(t16)]
	// FIN ACCESO
	// -----------------------------------------------------------

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------
	t12 = t12 + 1
	t14 = t14 + 1
	goto L0
L2:
	// FIN FOR
	// -----------------------------------------------------------

}
