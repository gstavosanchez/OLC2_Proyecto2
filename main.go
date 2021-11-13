package main

import (
	"fmt"
)

var P, H float64
var stack [50000000]float64
var heap [50000000]float64
var t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10 float64
var t11, t12, t13, t14, t15, t16, t17, t18, t19, t20 float64
var t21, t22, t23, t24, t25, t26, t27, t28, t29, t30 float64
var t31, t32, t33, t34, t35, t36, t37, t38, t39, t40 float64
var t41, t42, t43, t44, t45, t46, t47, t48, t49, t50 float64
var t51, t52, t53, t54, t55, t56, t57, t58, t59, t60 float64
var t61, t62, t63, t64, t65, t66, t67, t68, t69, t70 float64
var t71, t72, t73, t74, t75, t76, t77, t78, t79, t80 float64
var t81, t82, t83, t84, t85, t86, t87, t88, t89, t90 float64
var t91, t92, t93, t94, t95, t96, t97, t98, t99, t100 float64
var t101, t102, t103, t104, t105, t106, t107, t108, t109, t110 float64
var t111, t112, t113, t114, t115, t116 float64

// -----------------------------------------------------------
// FUNCIONES NATIVAS
func indexValidate() {
	t12 = P + 1
	t13 = stack[int(t12)]
	t12 = t12 + 1
	t14 = stack[int(t12)]
	if t14 < 1 {
		goto L2
	}
	if t14 > t13 {
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
	t15 = 1
	goto L1
L3:
	t15 = 0
L1:
	stack[int(P)] = t15
	return
}
func printStr() {
	t62 = P + 1 // Puntero del parametro
	t63 = stack[int(t62)]
L11:
	t64 = heap[int(t63)]
	if t64 == -1 {
		goto L10
	}
	fmt.Printf("%c", int(t64))
	t63 = t63 + 1 //aumentar el contador del heap
	goto L11
L10:
	return
}

// -----------------------------------------------------------
// FUNCIONES
func printMatriz() {

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: [

	t61 = H
	heap[int(H)] = 91 // [
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t65 = P + 2
	t65 = t65 + 1
	stack[int(t65)] = t61
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t66 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------
	// -----------------------------------------------------------
	// INICIO FOR
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: MATRIX
	t68 = P + 1 //Posicionar el puntero
	t67 = stack[int(t68)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t69 = P + 2
	t70 = t67
	stack[int(t69)] = t67
	t71 = stack[int(t69)]
	t72 = heap[int(t71)]
	t75 = 0
	t77 = t70
	t77 = t77 + 1
L12:
	t74 = P + t69
	t73 = stack[int(t74)]
	if t75 < t72 {
		goto L13
	}
	goto L14
L13:
	t76 = heap[int(t77)]
	stack[int(2)] = t76

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: [

	t78 = H
	heap[int(H)] = 91 // [
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t79 = P + 3
	t79 = t79 + 1
	stack[int(t79)] = t78
	// -----------------------------------------------------------

	P = P + 3
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t80 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 3

	// FIN PRINT
	// -----------------------------------------------------------
	// -----------------------------------------------------------
	// INICIO FOR
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: I
	t82 = P + 2 //Posicionar el puntero
	t81 = stack[int(t82)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t83 = P + 3
	t84 = t81
	stack[int(t83)] = t81
	t85 = stack[int(t83)]
	t86 = heap[int(t85)]
	t89 = 0
	t91 = t84
	t91 = t91 + 1
L15:
	t88 = P + t83
	t87 = stack[int(t88)]
	if t89 < t86 {
		goto L16
	}
	goto L17
L16:
	t90 = heap[int(t91)]
	stack[int(3)] = t90

	// -----------------------------------------------------------
	// INICIO PRINT
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: J
	t93 = P + 3 //Posicionar el puntero
	t92 = stack[int(t93)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// PRIMITIVO:

	t94 = H
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t95 = P + 4
	t95 = t95 + 1
	stack[int(t95)] = t94
	// -----------------------------------------------------------

	P = P + 4
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t96 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 4

	// FIN PRINT
	// -----------------------------------------------------------
	t89 = t89 + 1
	t91 = t91 + 1
	goto L15
L17:
	// FIN FOR
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: ]

	t97 = H
	heap[int(H)] = 93 // ]
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t98 = P + 3
	t98 = t98 + 1
	stack[int(t98)] = t97
	// -----------------------------------------------------------

	P = P + 3
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t99 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 3

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------
	t75 = t75 + 1
	t77 = t77 + 1
	goto L12
L14:
	// FIN FOR
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: ]

	t100 = H
	heap[int(H)] = 93 // ]
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t101 = P + 2
	t101 = t101 + 1
	stack[int(t101)] = t100
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t102 = stack[int(P)]
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
	// ARREGLO
	t0 = H             // Apuntador donde se guarda el array
	t1 = H             // Contador del heap
	H = H + 12         //Almacenar espacio en heap longitud + array
	heap[int(t1)] = 11 // Guadar longtud arreglo
	t1 = t1 + 1
	// GUARDAR VALORES
	// PRIMITIVO: 1
	heap[int(t1)] = 1
	t1 = t1 + 1
	// PRIMITIVO: 5
	heap[int(t1)] = 5
	t1 = t1 + 1
	// PRIMITIVO: 8
	heap[int(t1)] = 8
	t1 = t1 + 1
	// PRIMITIVO: 1
	t2 = 0 - 1

	heap[int(t1)] = t2
	t1 = t1 + 1
	// PRIMITIVO: 21
	heap[int(t1)] = 21
	t1 = t1 + 1
	// PRIMITIVO: 42
	heap[int(t1)] = 42
	t1 = t1 + 1
	// PRIMITIVO: 55
	t3 = 0 - 55

	heap[int(t1)] = t3
	t1 = t1 + 1
	// PRIMITIVO: 123
	heap[int(t1)] = 123
	t1 = t1 + 1
	// PRIMITIVO: 5
	t4 = 0 - 5

	heap[int(t1)] = t4
	t1 = t1 + 1
	// PRIMITIVO: 5
	heap[int(t1)] = 5
	t1 = t1 + 1
	// PRIMITIVO: 11
	heap[int(t1)] = 11
	t1 = t1 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	// -----------------------------------------------------------
	// ASIGNACION: RANDOM
	stack[int(0)] = t0
	// FIN ASIGNACION DE RANDOM
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ARREGLO
	t5 = H            // Apuntador donde se guarda el array
	t6 = H            // Contador del heap
	H = H + 3         //Almacenar espacio en heap longitud + array
	heap[int(t6)] = 2 // Guadar longtud arreglo
	t6 = t6 + 1
	// GUARDAR VALORES

	// -----------------------------------------------------------
	// ARREGLO
	t7 = H            // Apuntador donde se guarda el array
	t8 = H            // Contador del heap
	H = H + 5         //Almacenar espacio en heap longitud + array
	heap[int(t8)] = 4 // Guadar longtud arreglo
	t8 = t8 + 1
	// GUARDAR VALORES
	// PRIMITIVO: 1

	// -----------------------------------------------------------
	// ACCESO ARRAY "RANDOM"
	t9 = stack[int(0)]
	t11 = heap[int(t9)]
	// PRIMITIVO: 1

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T11
	// 1
	t16 = P + 1
	t16 = t16 + 1
	stack[int(t16)] = t11
	t16 = t16 + 1
	stack[int(t16)] = 1
	P = P + 1
	indexValidate()
	t17 = stack[int(P)]
	P = P - 1
	// -----------------------------------------------------------

	if t17 == 1 {
		goto L0
	}
	t9 = t9 + 1 //saltar len
	// PRIMITIVO: 1
	t9 = t9 + 0
	t10 = heap[int(t9)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L0:
	heap[int(t8)] = t10
	t8 = t8 + 1
	// PRIMITIVO: 51
	heap[int(t8)] = 51
	t8 = t8 + 1
	// PRIMITIVO: 4

	// -----------------------------------------------------------
	// ACCESO ARRAY "RANDOM"
	t18 = stack[int(0)]
	t20 = heap[int(t18)]
	// PRIMITIVO: 4

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T20
	// 4
	t21 = P + 1
	t21 = t21 + 1
	stack[int(t21)] = t20
	t21 = t21 + 1
	stack[int(t21)] = 4
	P = P + 1
	indexValidate()
	t22 = stack[int(P)]
	P = P - 1
	// -----------------------------------------------------------

	if t22 == 1 {
		goto L4
	}
	t18 = t18 + 1 //saltar len
	// PRIMITIVO: 4
	t18 = t18 + 3
	t19 = heap[int(t18)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L4:
	// PRIMITIVO: 1
	t23 = t19 * 1

	heap[int(t8)] = t23
	t8 = t8 + 1
	// PRIMITIVO: 3

	// -----------------------------------------------------------
	// ACCESO ARRAY "RANDOM"
	t24 = stack[int(0)]
	t26 = heap[int(t24)]
	// PRIMITIVO: 3

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T26
	// 3
	t27 = P + 1
	t27 = t27 + 1
	stack[int(t27)] = t26
	t27 = t27 + 1
	stack[int(t27)] = 3
	P = P + 1
	indexValidate()
	t28 = stack[int(P)]
	P = P - 1
	// -----------------------------------------------------------

	if t28 == 1 {
		goto L5
	}
	t24 = t24 + 1 //saltar len
	// PRIMITIVO: 3
	t24 = t24 + 2
	t25 = heap[int(t24)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L5:
	// PRIMITIVO: 10
	t29 = t25 * 10

	heap[int(t8)] = t29
	t8 = t8 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	heap[int(t6)] = t7
	t6 = t6 + 1

	// -----------------------------------------------------------
	// ARREGLO
	t30 = H            // Apuntador donde se guarda el array
	t31 = H            // Contador del heap
	H = H + 5          //Almacenar espacio en heap longitud + array
	heap[int(t31)] = 4 // Guadar longtud arreglo
	t31 = t31 + 1
	// GUARDAR VALORES
	// PRIMITIVO: 1
	heap[int(t31)] = 1
	t31 = t31 + 1
	// PRIMITIVO: 2
	heap[int(t31)] = 2
	t31 = t31 + 1
	// PRIMITIVO: 3
	heap[int(t31)] = 3
	t31 = t31 + 1
	// PRIMITIVO: 4
	heap[int(t31)] = 4
	t31 = t31 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	heap[int(t6)] = t30
	t6 = t6 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	// -----------------------------------------------------------
	// ASIGNACION: A
	stack[int(1)] = t5
	// FIN ASIGNACION DE A
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ARREGLO
	t32 = H            // Apuntador donde se guarda el array
	t33 = H            // Contador del heap
	H = H + 3          //Almacenar espacio en heap longitud + array
	heap[int(t33)] = 2 // Guadar longtud arreglo
	t33 = t33 + 1
	// GUARDAR VALORES

	// -----------------------------------------------------------
	// ARREGLO
	t34 = H            // Apuntador donde se guarda el array
	t35 = H            // Contador del heap
	H = H + 5          //Almacenar espacio en heap longitud + array
	heap[int(t35)] = 4 // Guadar longtud arreglo
	t35 = t35 + 1
	// GUARDAR VALORES
	// PRIMITIVO: 1
	heap[int(t35)] = 1
	t35 = t35 + 1
	// PRIMITIVO: 2
	heap[int(t35)] = 2
	t35 = t35 + 1
	// PRIMITIVO: 3
	heap[int(t35)] = 3
	t35 = t35 + 1
	// PRIMITIVO: 4
	heap[int(t35)] = 4
	t35 = t35 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	heap[int(t33)] = t34
	t33 = t33 + 1

	// -----------------------------------------------------------
	// ARREGLO
	t36 = H            // Apuntador donde se guarda el array
	t37 = H            // Contador del heap
	H = H + 5          //Almacenar espacio en heap longitud + array
	heap[int(t37)] = 4 // Guadar longtud arreglo
	t37 = t37 + 1
	// GUARDAR VALORES
	// PRIMITIVO: 1

	// -----------------------------------------------------------
	// ACCESO ARRAY "RANDOM"
	t38 = stack[int(0)]
	t40 = heap[int(t38)]
	// PRIMITIVO: 1

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T40
	// 1
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
		goto L6
	}
	t38 = t38 + 1 //saltar len
	// PRIMITIVO: 1
	t38 = t38 + 0
	t39 = heap[int(t38)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L6:
	heap[int(t37)] = t39
	t37 = t37 + 1
	// PRIMITIVO: 51
	heap[int(t37)] = 51
	t37 = t37 + 1
	// PRIMITIVO: 4

	// -----------------------------------------------------------
	// ACCESO ARRAY "RANDOM"
	t43 = stack[int(0)]
	t45 = heap[int(t43)]
	// PRIMITIVO: 4

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T45
	// 4
	t46 = P + 2
	t46 = t46 + 1
	stack[int(t46)] = t45
	t46 = t46 + 1
	stack[int(t46)] = 4
	P = P + 2
	indexValidate()
	t47 = stack[int(P)]
	P = P - 2
	// -----------------------------------------------------------

	if t47 == 1 {
		goto L7
	}
	t43 = t43 + 1 //saltar len
	// PRIMITIVO: 4
	t43 = t43 + 3
	t44 = heap[int(t43)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L7:
	// PRIMITIVO: 1
	t48 = t44 * 1

	heap[int(t37)] = t48
	t37 = t37 + 1
	// PRIMITIVO: 3

	// -----------------------------------------------------------
	// ACCESO ARRAY "RANDOM"
	t49 = stack[int(0)]
	t51 = heap[int(t49)]
	// PRIMITIVO: 3

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T51
	// 3
	t52 = P + 2
	t52 = t52 + 1
	stack[int(t52)] = t51
	t52 = t52 + 1
	stack[int(t52)] = 3
	P = P + 2
	indexValidate()
	t53 = stack[int(P)]
	P = P - 2
	// -----------------------------------------------------------

	if t53 == 1 {
		goto L8
	}
	t49 = t49 + 1 //saltar len
	// PRIMITIVO: 3
	t49 = t49 + 2
	t50 = heap[int(t49)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L8:
	// PRIMITIVO: 10
	t54 = t50 * 10

	heap[int(t37)] = t54
	t37 = t37 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	heap[int(t33)] = t36
	t33 = t33 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	// -----------------------------------------------------------
	// ASIGNACION: B
	stack[int(2)] = t32
	// FIN ASIGNACION DE B
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ARREGLO
	t55 = H            // Apuntador donde se guarda el array
	t56 = H            // Contador del heap
	H = H + 3          //Almacenar espacio en heap longitud + array
	heap[int(t56)] = 2 // Guadar longtud arreglo
	t56 = t56 + 1
	// GUARDAR VALORES

	// -----------------------------------------------------------
	// ARREGLO
	t57 = H            // Apuntador donde se guarda el array
	t58 = H            // Contador del heap
	H = H + 5          //Almacenar espacio en heap longitud + array
	heap[int(t58)] = 4 // Guadar longtud arreglo
	t58 = t58 + 1
	// GUARDAR VALORES
	// PRIMITIVO: 0.0
	heap[int(t58)] = 0.0
	t58 = t58 + 1
	// PRIMITIVO: 0.0
	heap[int(t58)] = 0.0
	t58 = t58 + 1
	// PRIMITIVO: 0.0
	heap[int(t58)] = 0.0
	t58 = t58 + 1
	// PRIMITIVO: 0.0
	heap[int(t58)] = 0.0
	t58 = t58 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	heap[int(t56)] = t57
	t56 = t56 + 1

	// -----------------------------------------------------------
	// ARREGLO
	t59 = H            // Apuntador donde se guarda el array
	t60 = H            // Contador del heap
	H = H + 5          //Almacenar espacio en heap longitud + array
	heap[int(t60)] = 4 // Guadar longtud arreglo
	t60 = t60 + 1
	// GUARDAR VALORES
	// PRIMITIVO: 0.0
	heap[int(t60)] = 0.0
	t60 = t60 + 1
	// PRIMITIVO: 0.0
	heap[int(t60)] = 0.0
	t60 = t60 + 1
	// PRIMITIVO: 0.0
	heap[int(t60)] = 0.0
	t60 = t60 + 1
	// PRIMITIVO: 0.0
	heap[int(t60)] = 0.0
	t60 = t60 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	heap[int(t56)] = t59
	t56 = t56 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	// -----------------------------------------------------------
	// ASIGNACION: AUXILIAR
	stack[int(3)] = t55
	// FIN ASIGNACION DE AUXILIAR
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: MATRIZ A

	t103 = H
	heap[int(H)] = 77 // M
	H = H + 1
	heap[int(H)] = 65 // A
	H = H + 1
	heap[int(H)] = 84 // T
	H = H + 1
	heap[int(H)] = 82 // R
	H = H + 1
	heap[int(H)] = 73 // I
	H = H + 1
	heap[int(H)] = 90 // Z
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t104 = P + 4
	t104 = t104 + 1
	stack[int(t104)] = t103
	// -----------------------------------------------------------

	P = P + 4
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t105 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 4

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// LLAMADO A FUNCION: PRINTMATRIZ
	// GUARDAR TEMPORALES

	t106 = P + 4
	stack[int(t106)] = t105

	// FIN DE GUARDAR TEMPORALES
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: A
	t107 = stack[int(1)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t108 = P + 6
	stack[int(t108)] = t107
	P = P + 5
	printMatriz()
	t108 = stack[int(P)]
	P = P - 5
	// RECUPERAR TEMPORALES

	t109 = P + 4
	t105 = stack[int(t109)]

	// FIN DE RECUPERACION
	// FIN LLAMADO A FUNCION
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: MATRIZ B

	t110 = H
	heap[int(H)] = 77 // M
	H = H + 1
	heap[int(H)] = 65 // A
	H = H + 1
	heap[int(H)] = 84 // T
	H = H + 1
	heap[int(H)] = 82 // R
	H = H + 1
	heap[int(H)] = 73 // I
	H = H + 1
	heap[int(H)] = 90 // Z
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 98 // b
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t111 = P + 4
	t111 = t111 + 1
	stack[int(t111)] = t110
	// -----------------------------------------------------------

	P = P + 4
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t112 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 4

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// LLAMADO A FUNCION: PRINTMATRIZ
	// GUARDAR TEMPORALES

	t113 = P + 4
	stack[int(t113)] = t105
	t113 = t113 + 1
	stack[int(t113)] = t112

	// FIN DE GUARDAR TEMPORALES
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: B
	t114 = stack[int(2)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t115 = P + 7
	stack[int(t115)] = t114
	P = P + 6
	printMatriz()
	t115 = stack[int(P)]
	P = P - 6
	// RECUPERAR TEMPORALES

	t116 = P + 4
	t105 = stack[int(t116)]
	t116 = t116 + 1
	t112 = stack[int(t116)]

	// FIN DE RECUPERACION
	// FIN LLAMADO A FUNCION
	// -----------------------------------------------------------

}
