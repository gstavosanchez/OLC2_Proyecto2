package main

import (
	"fmt"
)

var P, H float64
var stack [30101999]float64
var heap [30101999]float64
var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23 float64

// -----------------------------------------------------------
// FUNCIONES NATIVAS
func indexValidate() {
	t9 = P + 1
	t10 = stack[int(t9)]
	t9 = t9 + 1
	t11 = stack[int(t9)]
	if t11 < 1 {
		goto L2
	}
	if t11 > t10 {
		goto L2
	}
	goto L3
L2:
	fmt.Printf("%c", int(66))
	fmt.Printf("%c", int(111))
	fmt.Printf("%c", int(117))
	fmt.Printf("%c", int(110))
	fmt.Printf("%c", int(100))
	fmt.Printf("%c", int(115))
	fmt.Printf("%c", int(69))
	fmt.Printf("%c", int(114))
	fmt.Printf("%c", int(114))
	fmt.Printf("%c", int(111))
	fmt.Printf("%c", int(114))
	t12 = 1
	goto L1
L3:
	t12 = 0
L1:
	stack[int(P)] = t12
	return
}
func printStr() {
	t19 = P + 1 // Puntero del parametro
	t20 = stack[int(t19)]
L5:
	t21 = heap[int(t20)]
	if t21 == -1 {
		goto L4
	}
	fmt.Printf("%c", int(t21))
	t20 = t20 + 1 //aumentar el contador del heap
	goto L5
L4:
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
	H = H + 7         //Almacenar espacio en heap longitud + array
	heap[int(t1)] = 6 // Guadar longtud arreglo
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
	// ARREGLO
	t2 = H            // Apuntador donde se guarda el array
	t3 = H            // Contador del heap
	H = H + 4         //Almacenar espacio en heap longitud + array
	heap[int(t3)] = 3 // Guadar longtud arreglo
	t3 = t3 + 1
	// GUARDAR VALORES
	// PRIMITIVO: 10
	heap[int(t3)] = 10
	t3 = t3 + 1
	// PRIMITIVO: 25
	heap[int(t3)] = 25
	t3 = t3 + 1

	// -----------------------------------------------------------
	// ARREGLO
	t4 = H            // Apuntador donde se guarda el array
	t5 = H            // Contador del heap
	H = H + 4         //Almacenar espacio en heap longitud + array
	heap[int(t5)] = 3 // Guadar longtud arreglo
	t5 = t5 + 1
	// GUARDAR VALORES
	// PRIMITIVO: 50
	heap[int(t5)] = 50
	t5 = t5 + 1
	// PRIMITIVO: 25
	heap[int(t5)] = 25
	t5 = t5 + 1
	// PRIMITIVO: 200
	heap[int(t5)] = 200
	t5 = t5 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	heap[int(t3)] = t4
	t3 = t3 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	heap[int(t1)] = t2
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
	// PRIMITIVO: 6
	// PRIMITIVO: 100

	// -----------------------------------------------------------
	// ACCESO ARRAY "A"
	t6 = stack[int(0)]
	t8 = heap[int(t6)]
	// PRIMITIVO: 6

	// -----------------------------------------------------------
	t13 = P + 1
	t13 = t13 + 1
	stack[int(t13)] = t8
	t13 = t13 + 1
	stack[int(t13)] = 6
	P = P + 1
	indexValidate()
	t14 = stack[int(P)]
	P = P - 1
	// -----------------------------------------------------------

	if t14 == 1 {
		goto L0
	}
	t6 = t6 + 1 //saltar len
	// PRIMITIVO: 6
	t6 = t6 + 5
	// PRIMITIVO: 100
	t15 = heap[int(t6)]

	// -----------------------------------------------------------
	t16 = P + 1
	t16 = t16 + 1
	stack[int(t16)] = t15
	t16 = t16 + 1
	stack[int(t16)] = 100
	P = P + 1
	indexValidate()
	t17 = stack[int(P)]
	P = P - 1
	// -----------------------------------------------------------

	if t17 == 1 {
		goto L0
	}
	t15 = t15 + 1 //skip len
	t15 = t15 + 99
	t7 = heap[int(t15)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L0:
	fmt.Printf("%g", t7)
	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: GUS

	t18 = H
	heap[int(H)] = 103 // g
	H = H + 1
	heap[int(H)] = 117 // u
	H = H + 1
	heap[int(H)] = 115 // s
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t22 = P + 1
	t22 = t22 + 1
	stack[int(t22)] = t18
	// -----------------------------------------------------------

	P = P + 1
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t23 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 1

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

}
