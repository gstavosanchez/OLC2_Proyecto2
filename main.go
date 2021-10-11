package main

import "fmt"

var P, H float64
var stack [30101999]float64
var heap [30101999]float64
var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20 float64

// -----------------------------------------------------------
// FUNCIONES NATIVAS
func toUpperCase() {
	t1 = H
	t2 = P + 1
	t3 = stack[int(t2)]
L1:
	t4 = heap[int(t3)]
	if t4 == -1 {
		goto L0
	}
	if t4 < 97 {
		goto L2
	}
	if t4 > 122 {
		goto L2
	}
	t4 = t4 - 32 //Convertir en mayuscula
L2:
	heap[int(H)] = t4
	H = H + 1
	t3 = t3 + 1
	goto L1
L0:
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1
	stack[int(P)] = t1 // Guardar donde inica la nueva cadena
	return
}
func printStr() {
	t7 = P + 1 // Puntero del parametro
	t8 = stack[int(t7)]
L4:
	t9 = heap[int(t8)]
	if t9 == -1 {
		goto L3
	}
	fmt.Printf("%c", int(t9))
	t8 = t8 + 1 //aumentar el contador del heap
	goto L4
L3:
	return
}
func toLowerCase() {
	t13 = H
	t14 = P + 1
	t15 = stack[int(t14)]
L6:
	t16 = heap[int(t15)]
	if t16 == -1 {
		goto L5
	}
	if t16 < 65 {
		goto L7
	}
	if t16 > 90 {
		goto L7
	}
	t16 = t16 + 32 //Convertir en mayuscula
L7:
	heap[int(H)] = t16
	H = H + 1
	t15 = t15 + 1
	goto L6
L5:
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1
	stack[int(P)] = t13 // Guardar donde inica la nueva cadena
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
	// -----------------------------------------------------------
	// INICIO PRIMITIVO
	// HOLA

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

	// FIN PRIMITIVO
	// -----------------------------------------------------------

	t5 = P + 0
	t5 = t5 + 1
	stack[int(t5)] = t0
	P = P + 0
	toUpperCase()
	t6 = stack[int(P)]
	P = P - 0

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t10 = P + 0
	t10 = t10 + 1
	stack[int(t10)] = t6
	// -----------------------------------------------------------

	P = P + 0
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t11 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 0

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// -----------------------------------------------------------
	// INICIO PRIMITIVO
	// MUNDO

	t12 = H
	heap[int(H)] = 77 // M
	H = H + 1
	heap[int(H)] = 85 // U
	H = H + 1
	heap[int(H)] = 78 // N
	H = H + 1
	heap[int(H)] = 100 // d
	H = H + 1
	heap[int(H)] = 111 // o
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// FIN PRIMITIVO
	// -----------------------------------------------------------

	t17 = P + 0
	t17 = t17 + 1
	stack[int(t17)] = t12
	P = P + 0
	toLowerCase()
	t18 = stack[int(P)]
	P = P - 0

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t19 = P + 0
	t19 = t19 + 1
	stack[int(t19)] = t18
	// -----------------------------------------------------------

	P = P + 0
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t20 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 0

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

}
