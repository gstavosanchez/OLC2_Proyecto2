package main

import "fmt"

var P, H float64
var stack [30101999]float64
var heap [30101999]float64
var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13 float64

// -----------------------------------------------------------
// FUNCIONES NATIVAS
func divValidate() {
	t1 = P + 1
	t2 = stack[int(t1)]
	t1 = t1 + 1
	t3 = stack[int(t1)]
	if t3 != 0 {
		goto L1
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
	t4 = 0 // resultado incorrecto
	goto L0
L1:
	t4 = t2 / t3
L0:
	stack[int(P)] = t4 // gurdar resultado
	return
}
func printStr() {
	t9 = P + 1 // Puntero del parametro
	t10 = stack[int(t9)]
L3:
	t11 = heap[int(t10)]
	if t11 == -1 {
		goto L2
	}
	fmt.Printf("%c", int(t11))
	t10 = t10 + 1 //aumentar el contador del heap
	goto L3
L2:
	return
}

// ===========================================================
// MAIN
// ===========================================================
func main() {
	P = 0 // Puntero Stack
	H = 0 //Puntero Heap
	// -----------------------------------------------------------
	// INICIO EXPRESION ARITMETICA
	// -----------------------------------------------------------
	// INICIO PRIMITIVO
	// 56
	// FIN PRIMITIVO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO EXPRESION ARITMETICA
	// -----------------------------------------------------------
	// INICIO PRIMITIVO
	// 5
	// FIN PRIMITIVO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRIMITIVO
	// 5
	// FIN PRIMITIVO
	// -----------------------------------------------------------

	t0 = 5 - 5
	// FIN EXPRESION ARITMETICA
	// -----------------------------------------------------------
	// -----------------------------------------------------------
	// PASO DE PARAMETROS
	t5 = P + 0
	// 1ER PARAMETRO
	t5 = t5 + 1
	stack[int(t5)] = 56
	// 2DO PARAMETRO
	t5 = t5 + 1
	stack[int(t5)] = t0
	// FIN PASO DE PARAMETROS
	// -----------------------------------------------------------
	// CAMBIO DE ENTORNO
	P = P + 0
	divValidate()
	// GUARDAR EL RETURN DE LA FUNCION
	t6 = stack[int(P)]
	P = P - 0
	// FIN EXPRESION ARITMETICA
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ASIGNACION: A
	stack[int(0)] = t6
	// FIN ASIGNACION
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: A
	t7 = stack[int(0)]
	// FIN ACCESO
	// -----------------------------------------------------------

	fmt.Printf("%g", float64(t7))
	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// -----------------------------------------------------------
	// INICIO PRIMITIVO
	// ADIOS MUNDO

	t8 = H
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 100 // d
	H = H + 1
	heap[int(H)] = 105 // i
	H = H + 1
	heap[int(H)] = 111 // o
	H = H + 1
	heap[int(H)] = 115 // s
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
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

	// FIN PRIMITIVO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t12 = P + 1
	t12 = t12 + 1
	stack[int(t12)] = t8
	// -----------------------------------------------------------

	P = P + 1
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t13 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 1

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

}
