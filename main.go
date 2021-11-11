package main

import (
	"fmt"
)

var P, H float64
var stack [30101999]float64
var heap [30101999]float64
var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10 float64
var t11, t12, t13, t14, t15, t16, t17, t18, t19, t20 float64
var t21, t22, t23, t24, t25, t26, t27, t28 float64

// -----------------------------------------------------------
// FUNCIONES NATIVAS
func printStr() {
	t24 = P + 1 // Puntero del parametro
	t25 = stack[int(t24)]
L4:
	t26 = heap[int(t25)]
	if t26 == -1 {
		goto L3
	}
	fmt.Printf("%c", int(t26))
	t25 = t25 + 1 //aumentar el contador del heap
	goto L4
L3:
	return
}

// ===========================================================
// MAIN
// ===========================================================
func main() {
	P = 0 // Puntero Stack
	H = 0 //Puntero Heap

	// -----------------------------------------------------------
	// DECLARAR STRUCT CARTELERA()
	t0 = H
	t1 = t0
	H = H + 1 //Almacenar espacio en heap del struct Cartelera

	// -----------------------------------------------------------
	// ARREGLO
	t2 = H            // Apuntador donde se guarda el array
	t3 = H            // Contador del heap
	H = H + 4         //Almacenar espacio en heap longitud + array
	heap[int(t3)] = 3 // Guadar longtud arreglo
	t3 = t3 + 1
	// GUARDAR VALORES
	// PRIMITIVO: PELICULA 1

	t4 = H
	heap[int(H)] = 80 // P
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 105 // i
	H = H + 1
	heap[int(H)] = 99 // c
	H = H + 1
	heap[int(H)] = 117 // u
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 49 // 1
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t3)] = t4
	t3 = t3 + 1
	// PRIMITIVO: PELICULA 2

	t5 = H
	heap[int(H)] = 80 // P
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 105 // i
	H = H + 1
	heap[int(H)] = 99 // c
	H = H + 1
	heap[int(H)] = 117 // u
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 50 // 2
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t3)] = t5
	t3 = t3 + 1
	// PRIMITIVO: PELICULA 3

	t6 = H
	heap[int(H)] = 80 // P
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 105 // i
	H = H + 1
	heap[int(H)] = 99 // c
	H = H + 1
	heap[int(H)] = 117 // u
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 51 // 3
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t3)] = t6
	t3 = t3 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	heap[int(t1)] = t2
	t1 = t1 + 1
	// -----------------------------------------------------------
	// ASIGNACION: CINEPOLIS
	stack[int(0)] = t0
	// FIN ASIGNACION DE CINEPOLIS
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO FOR

	// -----------------------------------------------------------
	// ARREGLO
	t7 = H            // Apuntador donde se guarda el array
	t8 = H            // Contador del heap
	H = H + 5         //Almacenar espacio en heap longitud + array
	heap[int(t8)] = 4 // Guadar longtud arreglo
	t8 = t8 + 1
	// GUARDAR VALORES
	// PRIMITIVO: PELICULA 1

	t9 = H
	heap[int(H)] = 80 // P
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 105 // i
	H = H + 1
	heap[int(H)] = 99 // c
	H = H + 1
	heap[int(H)] = 117 // u
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 49 // 1
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t8)] = t9
	t8 = t8 + 1
	// PRIMITIVO: PELICULA 2

	t10 = H
	heap[int(H)] = 80 // P
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 105 // i
	H = H + 1
	heap[int(H)] = 99 // c
	H = H + 1
	heap[int(H)] = 117 // u
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 50 // 2
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t8)] = t10
	t8 = t8 + 1
	// PRIMITIVO: PELICULA 3

	t11 = H
	heap[int(H)] = 80 // P
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 105 // i
	H = H + 1
	heap[int(H)] = 99 // c
	H = H + 1
	heap[int(H)] = 117 // u
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 51 // 3
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t8)] = t11
	t8 = t8 + 1
	// PRIMITIVO: PELICULA 3555

	t12 = H
	heap[int(H)] = 80 // P
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 105 // i
	H = H + 1
	heap[int(H)] = 99 // c
	H = H + 1
	heap[int(H)] = 117 // u
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 51 // 3
	H = H + 1
	heap[int(H)] = 53 // 5
	H = H + 1
	heap[int(H)] = 53 // 5
	H = H + 1
	heap[int(H)] = 53 // 5
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t8)] = t12
	t8 = t8 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	t13 = P + 1
	t14 = t7
	stack[int(t13)] = t7
	t15 = stack[int(t13)]
	t16 = heap[int(t15)]
	t19 = 0
	t21 = t14
	t21 = t21 + 1
L0:
	t18 = P + t13
	t17 = stack[int(t18)]
	if t19 < t16 {
		goto L1
	}
	goto L2
L1:
	t20 = heap[int(t21)]
	stack[int(1)] = t20

	// -----------------------------------------------------------
	// INICIO PRINT
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: PELICULA
	t23 = P + 1 //Posicionar el puntero
	t22 = stack[int(t23)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t27 = P + 2
	t27 = t27 + 1
	stack[int(t27)] = t22
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t28 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------
	t19 = t19 + 1
	t21 = t21 + 1
	goto L0
L2:
	// FIN FOR
	// -----------------------------------------------------------

}
