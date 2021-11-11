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
var t41, t42, t43, t44, t45, t46, t47, t48, t49, t50 float64
var t51, t52, t53, t54, t55, t56, t57, t58, t59, t60 float64
var t61, t62, t63, t64, t65, t66, t67, t68, t69, t70 float64
var t71, t72, t73, t74, t75, t76, t77, t78, t79, t80 float64
var t81, t82, t83, t84, t85 float64

// -----------------------------------------------------------
// FUNCIONES NATIVAS
func printStr() {
	t13 = P + 1 // Puntero del parametro
	t14 = stack[int(t13)]
L4:
	t15 = heap[int(t14)]
	if t15 == -1 {
		goto L3
	}
	fmt.Printf("%c", int(t15))
	t14 = t14 + 1 //aumentar el contador del heap
	goto L4
L3:
	return
}
func indexValidate() {
	t24 = P + 1
	t25 = stack[int(t24)]
	t24 = t24 + 1
	t26 = stack[int(t24)]
	if t26 < 1 {
		goto L7
	}
	if t26 > t25 {
		goto L7
	}
	goto L8
L7:
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
	t27 = 1
	goto L6
L8:
	t27 = 0
L6:
	stack[int(P)] = t27
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
	H = H + 3         //Almacenar espacio en heap longitud + array
	heap[int(t1)] = 2 // Guadar longtud arreglo
	t1 = t1 + 1
	// GUARDAR VALORES

	// -----------------------------------------------------------
	// ARREGLO
	t2 = H             // Apuntador donde se guarda el array
	t3 = H             // Contador del heap
	H = H + 14         //Almacenar espacio en heap longitud + array
	heap[int(t3)] = 13 // Guadar longtud arreglo
	t3 = t3 + 1
	// GUARDAR VALORES
	// PRIMITIVO: 12
	heap[int(t3)] = 12
	t3 = t3 + 1
	// PRIMITIVO: 9
	heap[int(t3)] = 9
	t3 = t3 + 1
	// PRIMITIVO: 4
	heap[int(t3)] = 4
	t3 = t3 + 1
	// PRIMITIVO: 99
	heap[int(t3)] = 99
	t3 = t3 + 1
	// PRIMITIVO: 56
	heap[int(t3)] = 56
	t3 = t3 + 1
	// PRIMITIVO: 34
	heap[int(t3)] = 34
	t3 = t3 + 1
	// PRIMITIVO: 78
	heap[int(t3)] = 78
	t3 = t3 + 1
	// PRIMITIVO: 22
	heap[int(t3)] = 22
	t3 = t3 + 1
	// PRIMITIVO: 1
	heap[int(t3)] = 1
	t3 = t3 + 1
	// PRIMITIVO: 3
	heap[int(t3)] = 3
	t3 = t3 + 1
	// PRIMITIVO: 10
	heap[int(t3)] = 10
	t3 = t3 + 1
	// PRIMITIVO: 13
	heap[int(t3)] = 13
	t3 = t3 + 1
	// PRIMITIVO: 120
	heap[int(t3)] = 120
	t3 = t3 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	heap[int(t1)] = t2
	t1 = t1 + 1

	// -----------------------------------------------------------
	// ARREGLO
	t4 = H             // Apuntador donde se guarda el array
	t5 = H             // Contador del heap
	H = H + 17         //Almacenar espacio en heap longitud + array
	heap[int(t5)] = 16 // Guadar longtud arreglo
	t5 = t5 + 1
	// GUARDAR VALORES
	// PRIMITIVO: 32
	heap[int(t5)] = 32
	t5 = t5 + 1
	// PRIMITIVO: 7
	// PRIMITIVO: 3
	t6 = 7 * 3

	heap[int(t5)] = t6
	t5 = t5 + 1
	// PRIMITIVO: 7
	heap[int(t5)] = 7
	t5 = t5 + 1
	// PRIMITIVO: 89
	heap[int(t5)] = 89
	t5 = t5 + 1
	// PRIMITIVO: 56
	heap[int(t5)] = 56
	t5 = t5 + 1
	// PRIMITIVO: 909
	heap[int(t5)] = 909
	t5 = t5 + 1
	// PRIMITIVO: 109
	heap[int(t5)] = 109
	t5 = t5 + 1
	// PRIMITIVO: 2
	heap[int(t5)] = 2
	t5 = t5 + 1
	// PRIMITIVO: 9
	heap[int(t5)] = 9
	t5 = t5 + 1
	// PRIMITIVO: 9874
	// PRIMITIVO: 0
	t7 = 0
	t8 = 1
	if 0 == 0 {
		goto L2
	}
L0:
	if t7 == 0 {
		goto L1
	}
	t8 = 9874 * t8
	t7 = t7 + 1
	goto L0
L2:
	t8 = 1
L1:
	heap[int(t5)] = t8
	t5 = t5 + 1
	// PRIMITIVO: 44
	heap[int(t5)] = 44
	t5 = t5 + 1
	// PRIMITIVO: 3
	heap[int(t5)] = 3
	t5 = t5 + 1
	// PRIMITIVO: 820
	// PRIMITIVO: 10
	t9 = 820 * 10

	heap[int(t5)] = t9
	t5 = t5 + 1
	// PRIMITIVO: 11
	heap[int(t5)] = 11
	t5 = t5 + 1
	// PRIMITIVO: 8
	// PRIMITIVO: 0
	t10 = 8 * 0

	// PRIMITIVO: 8
	t11 = t10 + 8

	heap[int(t5)] = t11
	t5 = t5 + 1
	// PRIMITIVO: 10
	heap[int(t5)] = 10
	t5 = t5 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	heap[int(t1)] = t4
	t1 = t1 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	// -----------------------------------------------------------
	// ASIGNACION: ARRAY
	stack[int(0)] = t0
	// FIN ASIGNACION DE ARRAY
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: -------------------------

	t12 = H
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = 45 // -
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t16 = P + 1
	t16 = t16 + 1
	stack[int(t16)] = t12
	// -----------------------------------------------------------

	P = P + 1
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t17 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 1

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: VALORES DESPUES DE QUICKSORT:

	t18 = H
	heap[int(H)] = 86 // V
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 111 // o
	H = H + 1
	heap[int(H)] = 114 // r
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 115 // s
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 100 // d
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 115 // s
	H = H + 1
	heap[int(H)] = 112 // p
	H = H + 1
	heap[int(H)] = 117 // u
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 115 // s
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 100 // d
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 81 // Q
	H = H + 1
	heap[int(H)] = 117 // u
	H = H + 1
	heap[int(H)] = 105 // i
	H = H + 1
	heap[int(H)] = 99 // c
	H = H + 1
	heap[int(H)] = 107 // k
	H = H + 1
	heap[int(H)] = 83 // S
	H = H + 1
	heap[int(H)] = 111 // o
	H = H + 1
	heap[int(H)] = 114 // r
	H = H + 1
	heap[int(H)] = 116 // t
	H = H + 1
	heap[int(H)] = 58 // :
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
	// INICIO FOR
	// PRIMITIVO: 1
	// PRIMITIVO: 1

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARRAY"
	t21 = stack[int(0)]
	t23 = heap[int(t21)]
	// PRIMITIVO: 1

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T23
	// 1
	t28 = P + 1
	t28 = t28 + 1
	stack[int(t28)] = t23
	t28 = t28 + 1
	stack[int(t28)] = 1
	P = P + 1
	indexValidate()
	t29 = stack[int(P)]
	P = P - 1
	// -----------------------------------------------------------

	if t29 == 1 {
		goto L5
	}
	t21 = t21 + 1 //saltar len
	// PRIMITIVO: 1
	t21 = t21 + 0
	t22 = heap[int(t21)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L5:
	t22 = heap[int(t22)]
	t30 = P + 1
	stack[int(t30)] = 1

L9:
	t32 = P + 1 //Posicionar el puntero
	t31 = stack[int(t32)]
	if t31 > t22 {
		goto L10
	}

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: 1
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: Y
	t34 = P + 1 //Posicionar el puntero
	t33 = stack[int(t34)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARRAY"
	t35 = stack[int(0)]
	t37 = heap[int(t35)]
	// PRIMITIVO: 1

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T37
	// 1
	t38 = P + 2
	t38 = t38 + 1
	stack[int(t38)] = t37
	t38 = t38 + 1
	stack[int(t38)] = 1
	P = P + 2
	indexValidate()
	t39 = stack[int(P)]
	P = P - 2
	// -----------------------------------------------------------

	if t39 == 1 {
		goto L12
	}
	t35 = t35 + 1 //saltar len
	// PRIMITIVO: 1
	t35 = t35 + 0
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: Y
	t41 = P + 1 //Posicionar el puntero
	t40 = stack[int(t41)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t42 = heap[int(t35)]
	t43 = heap[int(t42)]

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T43
	// T40
	t44 = P + 2
	t44 = t44 + 1
	stack[int(t44)] = t43
	t44 = t44 + 1
	stack[int(t44)] = t40
	P = P + 2
	indexValidate()
	t45 = stack[int(P)]
	P = P - 2
	// -----------------------------------------------------------

	if t45 == 1 {
		goto L12
	}
	t42 = t42 + 1 //skip len
	t40 = t40 - 1
	t42 = t42 + t40
	t36 = heap[int(t42)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L12:
	fmt.Printf("%d", int(t36))
	// PRIMITIVO:

	t46 = H
	heap[int(H)] = 9 //
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t47 = P + 2
	t47 = t47 + 1
	stack[int(t47)] = t46
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t48 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	// FIN PRINT
	// -----------------------------------------------------------
	goto L11

L11:
	t50 = P + 1 //Posicionar el puntero
	t49 = stack[int(t50)]
	t49 = t49 + 1
	stack[int(t50)] = t49
	goto L9
L10:
	// FIN FOR
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO:

	t51 = H
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t52 = P + 1
	t52 = t52 + 1
	stack[int(t52)] = t51
	// -----------------------------------------------------------

	P = P + 1
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t53 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 1

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: VALORES ANTES DE QUICKSORT

	t54 = H
	heap[int(H)] = 86 // V
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 111 // o
	H = H + 1
	heap[int(H)] = 114 // r
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 115 // s
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 110 // n
	H = H + 1
	heap[int(H)] = 116 // t
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 115 // s
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 100 // d
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 81 // Q
	H = H + 1
	heap[int(H)] = 117 // u
	H = H + 1
	heap[int(H)] = 105 // i
	H = H + 1
	heap[int(H)] = 99 // c
	H = H + 1
	heap[int(H)] = 107 // k
	H = H + 1
	heap[int(H)] = 115 // s
	H = H + 1
	heap[int(H)] = 111 // o
	H = H + 1
	heap[int(H)] = 114 // r
	H = H + 1
	heap[int(H)] = 116 // t
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t55 = P + 1
	t55 = t55 + 1
	stack[int(t55)] = t54
	// -----------------------------------------------------------

	P = P + 1
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t56 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 1

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------
	// -----------------------------------------------------------
	// INICIO FOR
	// PRIMITIVO: 1
	// PRIMITIVO: 2

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARRAY"
	t57 = stack[int(0)]
	t59 = heap[int(t57)]
	// PRIMITIVO: 2

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T59
	// 2
	t60 = P + 1
	t60 = t60 + 1
	stack[int(t60)] = t59
	t60 = t60 + 1
	stack[int(t60)] = 2
	P = P + 1
	indexValidate()
	t61 = stack[int(P)]
	P = P - 1
	// -----------------------------------------------------------

	if t61 == 1 {
		goto L13
	}
	t57 = t57 + 1 //saltar len
	// PRIMITIVO: 2
	t57 = t57 + 1
	t58 = heap[int(t57)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L13:
	t58 = heap[int(t58)]
	t62 = P + 1
	stack[int(t62)] = 1

L14:
	t64 = P + 1 //Posicionar el puntero
	t63 = stack[int(t64)]
	if t63 > t58 {
		goto L15
	}

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: 2
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: X
	t66 = P + 1 //Posicionar el puntero
	t65 = stack[int(t66)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARRAY"
	t67 = stack[int(0)]
	t69 = heap[int(t67)]
	// PRIMITIVO: 2

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T69
	// 2
	t70 = P + 2
	t70 = t70 + 1
	stack[int(t70)] = t69
	t70 = t70 + 1
	stack[int(t70)] = 2
	P = P + 2
	indexValidate()
	t71 = stack[int(P)]
	P = P - 2
	// -----------------------------------------------------------

	if t71 == 1 {
		goto L17
	}
	t67 = t67 + 1 //saltar len
	// PRIMITIVO: 2
	t67 = t67 + 1
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: X
	t73 = P + 1 //Posicionar el puntero
	t72 = stack[int(t73)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t74 = heap[int(t67)]
	t75 = heap[int(t74)]

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T75
	// T72
	t76 = P + 2
	t76 = t76 + 1
	stack[int(t76)] = t75
	t76 = t76 + 1
	stack[int(t76)] = t72
	P = P + 2
	indexValidate()
	t77 = stack[int(P)]
	P = P - 2
	// -----------------------------------------------------------

	if t77 == 1 {
		goto L17
	}
	t74 = t74 + 1 //skip len
	t72 = t72 - 1
	t74 = t74 + t72
	t68 = heap[int(t74)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L17:
	fmt.Printf("%d", int(t68))
	// PRIMITIVO:

	t78 = H
	heap[int(H)] = 9 //
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t79 = P + 2
	t79 = t79 + 1
	stack[int(t79)] = t78
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t80 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	// FIN PRINT
	// -----------------------------------------------------------
	goto L16

L16:
	t82 = P + 1 //Posicionar el puntero
	t81 = stack[int(t82)]
	t81 = t81 + 1
	stack[int(t82)] = t81
	goto L14
L15:
	// FIN FOR
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO:

	t83 = H
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t84 = P + 1
	t84 = t84 + 1
	stack[int(t84)] = t83
	// -----------------------------------------------------------

	P = P + 1
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t85 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 1

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

}
