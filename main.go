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
func printStr() {
	t3 = P + 1 // Puntero del parametro
	t4 = stack[int(t3)]
L2:
	t5 = heap[int(t4)]
	if t5 == -1 {
		goto L1
	}
	fmt.Printf("%c", int(t5))
	t4 = t4 + 1 //aumentar el contador del heap
	goto L2
L1:
	return
}
func indexValidate() {
	t34 = P + 1
	t35 = stack[int(t34)]
	t34 = t34 + 1
	t36 = stack[int(t34)]
	if t36 < 1 {
		goto L5
	}
	if t36 > t35 {
		goto L5
	}
	goto L6
L5:
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
	t37 = 1
	goto L4
L6:
	t37 = 0
L4:
	stack[int(P)] = t37
	return
}

// -----------------------------------------------------------
// FUNCIONES
func saludar() {

	// -----------------------------------------------------------
	// INICIO PRINT
	t2 = P + 1
	t0 = stack[int(t2)]
	t0 = t0 + 0
	t1 = heap[int(t0)]
	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t6 = P + 2
	t6 = t6 + 1
	stack[int(t6)] = t1
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t7 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------
	return
}

// ===========================================================
// MAIN
// ===========================================================
func main() {
	P = 0 // Puntero Stack
	H = 0 //Puntero Heap

	// -----------------------------------------------------------
	// DECLARAR STRUCT ALUMNO()
	t8 = H
	t9 = t8
	H = H + 2 //Almacenar espacio en heap del struct Alumno
	// PRIMITIVO: SISTEMAS

	t10 = H
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

	heap[int(t9)] = t10
	t9 = t9 + 1

	// -----------------------------------------------------------
	// DECLARAR STRUCT PERSONA()
	t11 = H
	t12 = t11
	H = H + 2 //Almacenar espacio en heap del struct Persona
	// PRIMITIVO: 23
	heap[int(t12)] = 23
	t12 = t12 + 1
	// PRIMITIVO: GUS

	t13 = H
	heap[int(H)] = 71 // G
	H = H + 1
	heap[int(H)] = 117 // u
	H = H + 1
	heap[int(H)] = 115 // s
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t12)] = t13
	t12 = t12 + 1
	heap[int(t9)] = t11
	t9 = t9 + 1
	// -----------------------------------------------------------
	// ASIGNACION: ESTUDIANTE
	stack[int(0)] = t8
	// FIN ASIGNACION DE ESTUDIANTE
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// LLAMADO A FUNCION: SALUDAR
	// GUARDAR TEMPORALES

	t14 = P + 1
	stack[int(t14)] = t8
	t14 = t14 + 1
	stack[int(t14)] = t9
	t14 = t14 + 1
	stack[int(t14)] = t10
	t14 = t14 + 1
	stack[int(t14)] = t11
	t14 = t14 + 1
	stack[int(t14)] = t12
	t14 = t14 + 1
	stack[int(t14)] = t13

	// FIN DE GUARDAR TEMPORALES
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: ESTUDIANTE
	t15 = stack[int(0)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t16 = P + 8
	stack[int(t16)] = t15
	P = P + 7
	saludar()
	t16 = stack[int(P)]
	P = P - 7
	// RECUPERAR TEMPORALES

	t17 = P + 1
	t8 = stack[int(t17)]
	t17 = t17 + 1
	t9 = stack[int(t17)]
	t17 = t17 + 1
	t10 = stack[int(t17)]
	t17 = t17 + 1
	t11 = stack[int(t17)]
	t17 = t17 + 1
	t12 = stack[int(t17)]
	t17 = t17 + 1
	t13 = stack[int(t17)]
	t17 = t17 + 1
	t15 = stack[int(t17)]

	// FIN DE RECUPERACION
	// FIN LLAMADO A FUNCION
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: HOLA MUNDO

	t18 = H
	heap[int(H)] = 104 // h
	H = H + 1
	heap[int(H)] = 111 // o
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 97 // a
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

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t19 = P + 1
	t19 = t19 + 1
	stack[int(t19)] = t18
	// -----------------------------------------------------------

	P = P + 1
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t20 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 1

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ARREGLO
	t21 = H            // Apuntador donde se guarda el array
	t22 = H            // Contador del heap
	H = H + 6          //Almacenar espacio en heap longitud + array
	heap[int(t22)] = 5 // Guadar longtud arreglo
	t22 = t22 + 1
	// GUARDAR VALORES
	// PRIMITIVO: H

	t23 = H
	heap[int(H)] = 104 // h
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t22)] = t23
	t22 = t22 + 1
	// PRIMITIVO: O

	t24 = H
	heap[int(H)] = 111 // o
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t22)] = t24
	t22 = t22 + 1
	// PRIMITIVO: L

	t25 = H
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t22)] = t25
	t22 = t22 + 1
	// PRIMITIVO: A

	t26 = H
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t22)] = t26
	t22 = t22 + 1

	// -----------------------------------------------------------
	// ARREGLO
	t27 = H            // Apuntador donde se guarda el array
	t28 = H            // Contador del heap
	H = H + 3          //Almacenar espacio en heap longitud + array
	heap[int(t28)] = 2 // Guadar longtud arreglo
	t28 = t28 + 1
	// GUARDAR VALORES
	// PRIMITIVO: JA

	t29 = H
	heap[int(H)] = 106 // j
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t28)] = t29
	t28 = t28 + 1
	// PRIMITIVO: JA

	t30 = H
	heap[int(H)] = 106 // j
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t28)] = t30
	t28 = t28 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	heap[int(t22)] = t27
	t22 = t22 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	// -----------------------------------------------------------
	// ASIGNACION: HO
	stack[int(1)] = t21
	// FIN ASIGNACION DE HO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: 5
	// PRIMITIVO: 1

	// -----------------------------------------------------------
	// ACCESO ARRAY "HO"
	t31 = stack[int(1)]
	t33 = heap[int(t31)]
	// PRIMITIVO: 5

	// -----------------------------------------------------------
	t38 = P + 2
	t38 = t38 + 1
	stack[int(t38)] = t33
	t38 = t38 + 1
	stack[int(t38)] = 5
	P = P + 2
	indexValidate()
	t39 = stack[int(P)]
	P = P - 2
	// -----------------------------------------------------------

	if t39 == 1 {
		goto L3
	}
	t31 = t31 + 1 //saltar len
	// PRIMITIVO: 5
	t31 = t31 + 4
	// PRIMITIVO: 1
	t40 = heap[int(t31)]

	// -----------------------------------------------------------
	t41 = P + 2
	t41 = t41 + 1
	stack[int(t41)] = t40
	t41 = t41 + 1
	stack[int(t41)] = 1
	P = P + 2
	indexValidate()
	t42 = stack[int(P)]
	P = P - 2
	// -----------------------------------------------------------

	if t42 == 1 {
		goto L3
	}
	t40 = t40 + 1 //skip len
	t40 = t40 + 0
	t32 = heap[int(t40)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L3:
	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t43 = P + 2
	t43 = t43 + 1
	stack[int(t43)] = t32
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t44 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

}
