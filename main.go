package main

import (
	"fmt"
)

var P, H float64
var stack [30101999]float64
var heap [30101999]float64
var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14 float64

// -----------------------------------------------------------
// FUNCIONES NATIVAS
func potenciaStr() {
	t1 = H     // Puntero, iniciara la nueva cadena
	t2 = P + 1 //Puntero, 1er parametro
	t3 = stack[int(t2)]
	t2 = t2 + 1 //Puntero, 2do parametro
	t4 = stack[int(t2)]
	t6 = 1  // Contador, compara con el expo
	t7 = t3 // Copia del incio del 1er param
L1:
	t5 = heap[int(t3)]
	if t5 == -1 {
		goto L2
	}
	heap[int(H)] = t5
	H = H + 1
	t3 = t3 + 1 //Aumentar puntero del heap
	goto L1
L2:
	if t6 == t4 {
		goto L0
	}
	t6 = t6 + 1 //Aumentar contador de comparacion con el exp.
	t3 = t7     // Reiniciar puntero del heap
	goto L1
L0:
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1
	stack[int(P)] = t1 // Guardar donde inica la cadena unida
	return
}
func printStr() {
	t10 = P + 1 // Puntero del parametro
	t11 = stack[int(t10)]
L4:
	t12 = heap[int(t11)]
	if t12 == -1 {
		goto L3
	}
	fmt.Printf("%c", int(t12))
	t11 = t11 + 1 //aumentar el contador del heap
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
	// INICIO PRINT
	// PRIMITIVO: HA

	t0 = H
	heap[int(H)] = 104 // h
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// PRIMITIVO: 3
	// -----------------------------------------------------------
	// PASO DE PARAMETROS
	t8 = P + 0
	// 1ER PARAMETRO
	t8 = t8 + 1
	stack[int(t8)] = t0
	// 2DO PARAMETRO
	t8 = t8 + 1
	stack[int(t8)] = 3
	// FIN PASO DE PARAMETROS
	// -----------------------------------------------------------
	// CAMBIO DE ENTORNO
	P = P + 0
	potenciaStr()
	// GUARDAR EL RETURN DE LA FUNCION
	t9 = stack[int(P)]
	P = P - 0

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t13 = P + 0
	t13 = t13 + 1
	stack[int(t13)] = t9
	// -----------------------------------------------------------

	P = P + 0
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t14 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 0

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

}
