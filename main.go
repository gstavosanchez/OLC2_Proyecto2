package main

import (
	"fmt"
)

var P, H float64
var stack [30101999]float64
var heap [30101999]float64
var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23, t24, t25, t26, t27 float64

// -----------------------------------------------------------
// FUNCIONES NATIVAS
func printStr() {
	t11 = P + 1 // Puntero del parametro
	t12 = stack[int(t11)]
L1:
	t13 = heap[int(t12)]
	if t13 == -1 {
		goto L0
	}
	fmt.Printf("%c", int(t13))
	t12 = t12 + 1 //aumentar el contador del heap
	goto L1
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

	// -----------------------------------------------------------
	// DECLARAR STRUCT ALUMNO()
	t5 = H
	t6 = t5
	H = H + 2 //Almacenar espacio en heap del struct Alumno
	// PRIMITIVO: SISTEMAS

	t7 = H
	heap[int(H)] = 115 // s
	H = H + 1
	heap[int(H)] = 105 // i
	H = H + 1
	heap[int(H)] = 115 // s
	H = H + 1
	heap[int(H)] = 116 // t
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 109 // m
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 115 // s
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t6)] = t7
	t6 = t6 + 1
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: A
	t8 = stack[int(0)]
	// FIN ACCESO
	// -----------------------------------------------------------

	heap[int(t6)] = t8
	t6 = t6 + 1
	// -----------------------------------------------------------
	// ASIGNACION: ESTUDIANTE
	stack[int(1)] = t5
	// FIN ASIGNACION DE ESTUDIANTE
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	t9 = stack[int(1)]
	t9 = t9 + 0
	t10 = heap[int(t9)]
	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t14 = P + 2
	t14 = t14 + 1
	stack[int(t14)] = t10
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t15 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	t16 = stack[int(1)]
	t16 = t16 + 1
	t18 = heap[int(t16)]
	t18 = t18 + 1
	t17 = heap[int(t18)]
	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t19 = P + 2
	t19 = t19 + 1
	stack[int(t19)] = t17
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t20 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------
	// PRIMITIVO: ELMER

	t21 = H
	heap[int(H)] = 69 // E
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 109 // m
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 114 // r
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	t22 = stack[int(1)]
	t22 = t22 + 1
	t23 = heap[int(t22)]
	t23 = t23 + 1
	heap[int(t23)] = t21

	// -----------------------------------------------------------
	// INICIO PRINT
	t24 = stack[int(0)]
	t24 = t24 + 1
	t25 = heap[int(t24)]
	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t26 = P + 2
	t26 = t26 + 1
	stack[int(t26)] = t25
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t27 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

}
