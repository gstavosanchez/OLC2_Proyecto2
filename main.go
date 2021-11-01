package main

import (
	"fmt"
)

var P, H float64
var stack [30101999]float64
var heap [30101999]float64
var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10 float64
var t11, t12, t13, t14, t15, t16, t17, t18, t19, t20 float64
var t21, t22, t23, t24, t25, t26, t27, t28, t29, t30 float64
var t31, t32, t33, t34, t35, t36, t37, t38, t39, t40 float64
var t41, t42, t43, t44 float64

// -----------------------------------------------------------
// FUNCIONES NATIVAS
func indexValidate() {
	t13 = P + 1
	t14 = stack[int(t13)]
	t13 = t13 + 1
	t15 = stack[int(t13)]
	if t15 < 1 {
		goto L2
	}
	if t15 > t14 {
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
	t16 = 1
	goto L1
L3:
	t16 = 0
L1:
	stack[int(P)] = t16
	return
}
func printStr() {
	t22 = P + 1 // Puntero del parametro
	t23 = stack[int(t22)]
L5:
	t24 = heap[int(t23)]
	if t24 == -1 {
		goto L4
	}
	fmt.Printf("%c", int(t24))
	t23 = t23 + 1 //aumentar el contador del heap
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
	H = H + 6         //Almacenar espacio en heap longitud + array
	heap[int(t1)] = 5 // Guadar longtud arreglo
	t1 = t1 + 1
	// GUARDAR VALORES
	// PRIMITIVO: H

	t2 = H
	heap[int(H)] = 104 // h
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t1)] = t2
	t1 = t1 + 1
	// PRIMITIVO: O

	t3 = H
	heap[int(H)] = 111 // o
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t1)] = t3
	t1 = t1 + 1
	// PRIMITIVO: L

	t4 = H
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t1)] = t4
	t1 = t1 + 1
	// PRIMITIVO: A

	t5 = H
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t1)] = t5
	t1 = t1 + 1

	// -----------------------------------------------------------
	// ARREGLO
	t6 = H            // Apuntador donde se guarda el array
	t7 = H            // Contador del heap
	H = H + 3         //Almacenar espacio en heap longitud + array
	heap[int(t7)] = 2 // Guadar longtud arreglo
	t7 = t7 + 1
	// GUARDAR VALORES
	// PRIMITIVO: JA

	t8 = H
	heap[int(H)] = 106 // j
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t7)] = t8
	t7 = t7 + 1
	// PRIMITIVO: JA

	t9 = H
	heap[int(H)] = 106 // j
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t7)] = t9
	t7 = t7 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	heap[int(t1)] = t6
	t1 = t1 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	// -----------------------------------------------------------
	// ASIGNACION: HO
	stack[int(0)] = t0
	// FIN ASIGNACION DE HO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: 5
	// PRIMITIVO: 1

	// -----------------------------------------------------------
	// ACCESO ARRAY "HO"
	t10 = stack[int(0)]
	t12 = heap[int(t10)]
	// PRIMITIVO: 5

	// -----------------------------------------------------------
	t17 = P + 1
	t17 = t17 + 1
	stack[int(t17)] = t12
	t17 = t17 + 1
	stack[int(t17)] = 5
	P = P + 1
	indexValidate()
	t18 = stack[int(P)]
	P = P - 1
	// -----------------------------------------------------------

	if t18 == 1 {
		goto L0
	}
	t10 = t10 + 1 //saltar len
	// PRIMITIVO: 5
	t10 = t10 + 4
	// PRIMITIVO: 1
	t19 = heap[int(t10)]

	// -----------------------------------------------------------
	t20 = P + 1
	t20 = t20 + 1
	stack[int(t20)] = t19
	t20 = t20 + 1
	stack[int(t20)] = 1
	P = P + 1
	indexValidate()
	t21 = stack[int(P)]
	P = P - 1
	// -----------------------------------------------------------

	if t21 == 1 {
		goto L0
	}
	t19 = t19 + 1 //skip len
	t19 = t19 + 0
	t11 = heap[int(t19)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L0:
	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t25 = P + 1
	t25 = t25 + 1
	stack[int(t25)] = t11
	// -----------------------------------------------------------

	P = P + 1
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t26 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 1

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------
	// PRIMITIVO: 5
	// PRIMITIVO: 1
	// PRIMITIVO: KKQQ

	t27 = H
	heap[int(H)] = 75 // K
	H = H + 1
	heap[int(H)] = 75 // K
	H = H + 1
	heap[int(H)] = 113 // q
	H = H + 1
	heap[int(H)] = 113 // q
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// ASIGNAR ACCESO ARRAY "HO"
	t28 = stack[int(0)]
	t29 = heap[int(t28)]
	// PRIMITIVO: 5

	// -----------------------------------------------------------
	t30 = P + 1
	t30 = t30 + 1
	stack[int(t30)] = t29
	t30 = t30 + 1
	stack[int(t30)] = 5
	P = P + 1
	indexValidate()
	t31 = stack[int(P)]
	P = P - 1
	// -----------------------------------------------------------

	if t31 == 1 {
		goto L6
	}
	t28 = t28 + 1 //saltar len
	// PRIMITIVO: 5
	t28 = t28 + 4
	// PRIMITIVO: 1
	t32 = heap[int(t28)]

	// -----------------------------------------------------------
	t33 = P + 1
	t33 = t33 + 1
	stack[int(t33)] = t32
	t33 = t33 + 1
	stack[int(t33)] = 1
	P = P + 1
	indexValidate()
	t34 = stack[int(P)]
	P = P - 1
	// -----------------------------------------------------------

	if t34 == 1 {
		goto L6
	}
	t32 = t32 + 1 //skip len
	t32 = t32 + 0
	heap[int(t32)] = t27
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L6:

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: 5
	// PRIMITIVO: 1

	// -----------------------------------------------------------
	// ACCESO ARRAY "HO"
	t35 = stack[int(0)]
	t37 = heap[int(t35)]
	// PRIMITIVO: 5

	// -----------------------------------------------------------
	t38 = P + 1
	t38 = t38 + 1
	stack[int(t38)] = t37
	t38 = t38 + 1
	stack[int(t38)] = 5
	P = P + 1
	indexValidate()
	t39 = stack[int(P)]
	P = P - 1
	// -----------------------------------------------------------

	if t39 == 1 {
		goto L7
	}
	t35 = t35 + 1 //saltar len
	// PRIMITIVO: 5
	t35 = t35 + 4
	// PRIMITIVO: 1
	t40 = heap[int(t35)]

	// -----------------------------------------------------------
	t41 = P + 1
	t41 = t41 + 1
	stack[int(t41)] = t40
	t41 = t41 + 1
	stack[int(t41)] = 1
	P = P + 1
	indexValidate()
	t42 = stack[int(P)]
	P = P - 1
	// -----------------------------------------------------------

	if t42 == 1 {
		goto L7
	}
	t40 = t40 + 1 //skip len
	t40 = t40 + 0
	t36 = heap[int(t40)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L7:
	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t43 = P + 1
	t43 = t43 + 1
	stack[int(t43)] = t36
	// -----------------------------------------------------------

	P = P + 1
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t44 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 1

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

}
