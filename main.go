package main

import (
	"fmt"
)

var P, H float64
var stack [30101999]float64
var heap [30101999]float64
var t0, t1, t2, t3 float64

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
	// PRIMITIVO: 100
	heap[int(t1)] = 100
	t1 = t1 + 1
	// PRIMITIVO: 200
	heap[int(t1)] = 200
	t1 = t1 + 1
	// PRIMITIVO: 300
	heap[int(t1)] = 300
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
	t2 = stack[int(0)]
	t2 = t2 + 1 //saltar len
	// PRIMITIVO: 2
	t2 = t2 + 1
	t3 = heap[int(t2)]
	fmt.Printf("%g", t3)
	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

}
