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
var t81, t82, t83, t84, t85, t86, t87, t88, t89, t90 float64
var t91, t92, t93, t94, t95, t96, t97, t98, t99, t100 float64
var t101, t102, t103, t104, t105, t106, t107, t108, t109, t110 float64
var t111, t112, t113, t114, t115, t116, t117, t118, t119, t120 float64
var t121, t122, t123, t124, t125, t126, t127, t128, t129, t130 float64
var t131, t132, t133, t134, t135 float64

// -----------------------------------------------------------
// FUNCIONES NATIVAS
func printStr() {
	t31 = P + 1 // Puntero del parametro
	t32 = stack[int(t31)]
L5:
	t33 = heap[int(t32)]
	if t33 == -1 {
		goto L4
	}
	fmt.Printf("%c", int(t33))
	t32 = t32 + 1 //aumentar el contador del heap
	goto L5
L4:
	return
}
func indexValidate() {
	t88 = P + 1
	t89 = stack[int(t88)]
	t88 = t88 + 1
	t90 = stack[int(t88)]
	if t90 < 1 {
		goto L14
	}
	if t90 > t89 {
		goto L14
	}
	goto L15
L14:
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
	t91 = 1
	goto L13
L15:
	t91 = 0
L13:
	stack[int(P)] = t91
	return
}

// -----------------------------------------------------------
// FUNCIONES
func contratar() {

	// -----------------------------------------------------------
	// DECLARAR STRUCT CONTRATO()
	t12 = H
	t13 = t12
	H = H + 2 //Almacenar espacio en heap del struct Contrato
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: ACTOR
	t15 = P + 1 //Posicionar el puntero
	t14 = stack[int(t15)]
	// FIN ACCESO
	// -----------------------------------------------------------

	heap[int(t13)] = t14
	t13 = t13 + 1
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: PELICULA
	t17 = P + 2 //Posicionar el puntero
	t16 = stack[int(t17)]
	// FIN ACCESO
	// -----------------------------------------------------------

	heap[int(t13)] = t16
	t13 = t13 + 1
	stack[int(P)] = t12
	goto L0
L0:
	return
}
func crearActor() {

	// -----------------------------------------------------------
	// DECLARAR STRUCT ACTOR()
	t18 = H
	t19 = t18
	H = H + 2 //Almacenar espacio en heap del struct Actor
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: NOMBRE
	t21 = P + 1 //Posicionar el puntero
	t20 = stack[int(t21)]
	// FIN ACCESO
	// -----------------------------------------------------------

	heap[int(t19)] = t20
	t19 = t19 + 1
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: EDAD
	t23 = P + 2 //Posicionar el puntero
	t22 = stack[int(t23)]
	// FIN ACCESO
	// -----------------------------------------------------------

	heap[int(t19)] = t22
	t19 = t19 + 1
	stack[int(P)] = t18
	goto L1
L1:
	return
}
func crearPelicula() {

	// -----------------------------------------------------------
	// DECLARAR STRUCT PELICULA()
	t24 = H
	t25 = t24
	H = H + 2 //Almacenar espacio en heap del struct Pelicula
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: NOMBRE
	t27 = P + 1 //Posicionar el puntero
	t26 = stack[int(t27)]
	// FIN ACCESO
	// -----------------------------------------------------------

	heap[int(t25)] = t26
	t25 = t25 + 1
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: POSICION
	t29 = P + 2 //Posicionar el puntero
	t28 = stack[int(t29)]
	// FIN ACCESO
	// -----------------------------------------------------------

	heap[int(t25)] = t28
	t25 = t25 + 1
	stack[int(P)] = t24
	goto L2
L2:
	return
}
func imprimir() {

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: ACTOR:

	t30 = H
	heap[int(H)] = 65 // A
	H = H + 1
	heap[int(H)] = 99 // c
	H = H + 1
	heap[int(H)] = 116 // t
	H = H + 1
	heap[int(H)] = 111 // o
	H = H + 1
	heap[int(H)] = 114 // r
	H = H + 1
	heap[int(H)] = 58 // :
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t34 = P + 2
	t34 = t34 + 1
	stack[int(t34)] = t30
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t35 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	t38 = P + 1
	t36 = stack[int(t38)]
	t36 = t36 + 0
	t39 = heap[int(t36)]
	t39 = t39 + 0
	t37 = heap[int(t39)]
	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t40 = P + 2
	t40 = t40 + 1
	stack[int(t40)] = t37
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t41 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	// PRIMITIVO:    EDAD:

	t42 = H
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 69 // E
	H = H + 1
	heap[int(H)] = 100 // d
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 100 // d
	H = H + 1
	heap[int(H)] = 58 // :
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t43 = P + 2
	t43 = t43 + 1
	stack[int(t43)] = t42
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t44 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	t47 = P + 1
	t45 = stack[int(t47)]
	t45 = t45 + 0
	t48 = heap[int(t45)]
	t48 = t48 + 1
	t46 = heap[int(t48)]
	fmt.Printf("%d", int(t46))
	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: PELICULA:

	t49 = H
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
	heap[int(H)] = 58 // :
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t50 = P + 2
	t50 = t50 + 1
	stack[int(t50)] = t49
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t51 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	t54 = P + 1
	t52 = stack[int(t54)]
	t52 = t52 + 1
	t55 = heap[int(t52)]
	t55 = t55 + 0
	t53 = heap[int(t55)]
	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t56 = P + 2
	t56 = t56 + 1
	stack[int(t56)] = t53
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t57 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	// PRIMITIVO:    GENERO:

	t58 = H
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 71 // G
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 110 // n
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 114 // r
	H = H + 1
	heap[int(H)] = 111 // o
	H = H + 1
	heap[int(H)] = 58 // :
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t59 = P + 2
	t59 = t59 + 1
	stack[int(t59)] = t58
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t60 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	t63 = P + 1
	t61 = stack[int(t63)]
	t61 = t61 + 1
	t64 = heap[int(t61)]
	t64 = t64 + 1
	t62 = heap[int(t64)]
	fmt.Printf("%d", int(t62))
	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------
	return
}
func contratos() {
	// -----------------------------------------------------------
	// INICIO FOR
	// PRIMITIVO: 1
	// PRIMITIVO: 1
	// PRIMITIVO: 1
	t65 = 1 * 1

	// PRIMITIVO: 2
	t66 = t65 + 2

	t67 = P + 1
	stack[int(t67)] = 1

L7:
	t69 = P + 1 //Posicionar el puntero
	t68 = stack[int(t69)]
	if t68 > t66 {
		goto L8
	}

	// -----------------------------------------------------------
	// DECLARAR STRUCT CONTRATO()
	t70 = H
	t71 = t70
	H = H + 2 //Almacenar espacio en heap del struct Contrato

	// -----------------------------------------------------------
	// DECLARAR STRUCT ACTOR()
	t72 = H
	t73 = t72
	H = H + 2 //Almacenar espacio en heap del struct Actor
	// PRIMITIVO:

	t74 = H
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t73)] = t74
	t73 = t73 + 1
	// PRIMITIVO: 0
	heap[int(t73)] = 0
	t73 = t73 + 1
	heap[int(t71)] = t72
	t71 = t71 + 1

	// -----------------------------------------------------------
	// DECLARAR STRUCT PELICULA()
	t75 = H
	t76 = t75
	H = H + 2 //Almacenar espacio en heap del struct Pelicula
	// PRIMITIVO:

	t77 = H
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t76)] = t77
	t76 = t76 + 1
	// PRIMITIVO: 0
	heap[int(t76)] = 0
	t76 = t76 + 1
	heap[int(t71)] = t75
	t71 = t71 + 1
	// -----------------------------------------------------------
	// ASIGNACION: CONTRATO
	t78 = P + 2
	stack[int(t78)] = t70
	// FIN ASIGNACION DE CONTRATO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO IF
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: I
	t80 = P + 1 //Posicionar el puntero
	t79 = stack[int(t80)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// PRIMITIVO: 4
	if t79 < 4 {
		goto L10
	}
	goto L11
L10:

	// -----------------------------------------------------------
	// LLAMADO A FUNCION: CREARACTOR
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: I
	t82 = P + 1 //Posicionar el puntero
	t81 = stack[int(t82)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO ARRAY "ACTORES"
	t83 = stack[int(0)]
	t85 = heap[int(t83)]
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: I
	t87 = P + 1 //Posicionar el puntero
	t86 = stack[int(t87)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	t92 = P + 3
	t92 = t92 + 1
	stack[int(t92)] = t85
	t92 = t92 + 1
	stack[int(t92)] = t86
	P = P + 3
	indexValidate()
	t93 = stack[int(P)]
	P = P - 3
	// -----------------------------------------------------------

	if t93 == 1 {
		goto L12
	}
	t83 = t83 + 1 //saltar len
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: I
	t95 = P + 1 //Posicionar el puntero
	t94 = stack[int(t95)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t94 = t94 - 1
	t83 = t83 + t94
	t84 = heap[int(t83)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L12:
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: I
	t97 = P + 1 //Posicionar el puntero
	t96 = stack[int(t97)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// PRIMITIVO: 38
	t98 = t96 + 38

	t99 = P + 4
	stack[int(t99)] = t84
	t99 = t99 + 1
	stack[int(t99)] = t98
	P = P + 3
	crearActor()
	t99 = stack[int(P)]
	P = P - 3
	// RECUPERAR TEMPORALES

	t100 = P + 3
	t81 = stack[int(t100)]

	// FIN DE RECUPERACION
	// FIN LLAMADO A FUNCION
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ASIGNACION: ACTOR
	t101 = P + 3
	stack[int(t101)] = t99
	// FIN ASIGNACION DE ACTOR
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// LLAMADO A FUNCION: CREARPELICULA
	// GUARDAR TEMPORALES

	t102 = P + 4
	stack[int(t102)] = t81

	// FIN DE GUARDAR TEMPORALES
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: I
	t104 = P + 1 //Posicionar el puntero
	t103 = stack[int(t104)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO ARRAY "PELICULAS"
	t105 = stack[int(1)]
	t107 = heap[int(t105)]
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: I
	t109 = P + 1 //Posicionar el puntero
	t108 = stack[int(t109)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	t110 = P + 5
	t110 = t110 + 1
	stack[int(t110)] = t107
	t110 = t110 + 1
	stack[int(t110)] = t108
	P = P + 5
	indexValidate()
	t111 = stack[int(P)]
	P = P - 5
	// -----------------------------------------------------------

	if t111 == 1 {
		goto L16
	}
	t105 = t105 + 1 //saltar len
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: I
	t113 = P + 1 //Posicionar el puntero
	t112 = stack[int(t113)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t112 = t112 - 1
	t105 = t105 + t112
	t106 = heap[int(t105)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L16:
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: I
	t115 = P + 1 //Posicionar el puntero
	t114 = stack[int(t115)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t116 = P + 6
	stack[int(t116)] = t106
	t116 = t116 + 1
	stack[int(t116)] = t114
	P = P + 5
	crearPelicula()
	t116 = stack[int(P)]
	P = P - 5
	// RECUPERAR TEMPORALES

	t117 = P + 4
	t81 = stack[int(t117)]
	t117 = t117 + 1
	t103 = stack[int(t117)]

	// FIN DE RECUPERACION
	// FIN LLAMADO A FUNCION
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ASIGNACION: PELICULA
	t118 = P + 4
	stack[int(t118)] = t116
	// FIN ASIGNACION DE PELICULA
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// LLAMADO A FUNCION: CONTRATAR
	// GUARDAR TEMPORALES

	t119 = P + 5
	stack[int(t119)] = t81
	t119 = t119 + 1
	stack[int(t119)] = t103

	// FIN DE GUARDAR TEMPORALES
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: ACTOR
	t121 = P + 3 //Posicionar el puntero
	t120 = stack[int(t121)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: PELICULA
	t123 = P + 4 //Posicionar el puntero
	t122 = stack[int(t123)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t124 = P + 8
	stack[int(t124)] = t120
	t124 = t124 + 1
	stack[int(t124)] = t122
	P = P + 7
	contratar()
	t124 = stack[int(P)]
	P = P - 7
	// RECUPERAR TEMPORALES

	t125 = P + 5
	t81 = stack[int(t125)]
	t125 = t125 + 1
	t103 = stack[int(t125)]

	// FIN DE RECUPERACION
	// FIN LLAMADO A FUNCION
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ASIGNACION: CONTRATO
	t126 = P + 2
	stack[int(t126)] = t124
	// FIN ASIGNACION DE CONTRATO
	// -----------------------------------------------------------

L11:
	// FIN IF
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// LLAMADO A FUNCION: IMPRIMIR
	// GUARDAR TEMPORALES

	t127 = P + 5
	stack[int(t127)] = t81
	t127 = t127 + 1
	stack[int(t127)] = t103

	// FIN DE GUARDAR TEMPORALES
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: CONTRATO
	t129 = P + 2 //Posicionar el puntero
	t128 = stack[int(t129)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t130 = P + 8
	stack[int(t130)] = t128
	P = P + 7
	imprimir()
	t130 = stack[int(P)]
	P = P - 7
	// RECUPERAR TEMPORALES

	t131 = P + 5
	t81 = stack[int(t131)]
	t131 = t131 + 1
	t103 = stack[int(t131)]

	// FIN DE RECUPERACION
	// FIN LLAMADO A FUNCION
	// -----------------------------------------------------------

	goto L9

L9:
	t133 = P + 1 //Posicionar el puntero
	t132 = stack[int(t133)]
	t132 = t132 + 1
	stack[int(t133)] = t132
	goto L7
L8:
	// FIN FOR
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
	t0 = H            // Apuntador donde se guarda el array
	t1 = H            // Contador del heap
	H = H + 5         //Almacenar espacio en heap longitud + array
	heap[int(t1)] = 4 // Guadar longtud arreglo
	t1 = t1 + 1
	// GUARDAR VALORES
	// PRIMITIVO: ELIZABETH OLSEN

	t2 = H
	heap[int(H)] = 69 // E
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 105 // i
	H = H + 1
	heap[int(H)] = 122 // z
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 98 // b
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 116 // t
	H = H + 1
	heap[int(H)] = 104 // h
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 79 // O
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 115 // s
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 110 // n
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t1)] = t2
	t1 = t1 + 1
	// PRIMITIVO: ADAM SANDLER

	t3 = H
	heap[int(H)] = 65 // A
	H = H + 1
	heap[int(H)] = 100 // d
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 109 // m
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 83 // S
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 110 // n
	H = H + 1
	heap[int(H)] = 100 // d
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 114 // r
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t1)] = t3
	t1 = t1 + 1
	// PRIMITIVO: CHRISTIAN BALE

	t4 = H
	heap[int(H)] = 67 // C
	H = H + 1
	heap[int(H)] = 104 // h
	H = H + 1
	heap[int(H)] = 114 // r
	H = H + 1
	heap[int(H)] = 105 // i
	H = H + 1
	heap[int(H)] = 115 // s
	H = H + 1
	heap[int(H)] = 116 // t
	H = H + 1
	heap[int(H)] = 105 // i
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 110 // n
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 66 // B
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t1)] = t4
	t1 = t1 + 1
	// PRIMITIVO: JENNIFER ANISTON

	t5 = H
	heap[int(H)] = 74 // J
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 110 // n
	H = H + 1
	heap[int(H)] = 110 // n
	H = H + 1
	heap[int(H)] = 105 // i
	H = H + 1
	heap[int(H)] = 102 // f
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 114 // r
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 65 // A
	H = H + 1
	heap[int(H)] = 110 // n
	H = H + 1
	heap[int(H)] = 105 // i
	H = H + 1
	heap[int(H)] = 115 // s
	H = H + 1
	heap[int(H)] = 116 // t
	H = H + 1
	heap[int(H)] = 111 // o
	H = H + 1
	heap[int(H)] = 110 // n
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t1)] = t5
	t1 = t1 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	// -----------------------------------------------------------
	// ASIGNACION: ACTORES
	stack[int(0)] = t0
	// FIN ASIGNACION DE ACTORES
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ARREGLO
	t6 = H            // Apuntador donde se guarda el array
	t7 = H            // Contador del heap
	H = H + 5         //Almacenar espacio en heap longitud + array
	heap[int(t7)] = 4 // Guadar longtud arreglo
	t7 = t7 + 1
	// GUARDAR VALORES
	// PRIMITIVO: AVENGERS: AGE OF ULTRON

	t8 = H
	heap[int(H)] = 65 // A
	H = H + 1
	heap[int(H)] = 118 // v
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 110 // n
	H = H + 1
	heap[int(H)] = 103 // g
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 114 // r
	H = H + 1
	heap[int(H)] = 115 // s
	H = H + 1
	heap[int(H)] = 58 // :
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 65 // A
	H = H + 1
	heap[int(H)] = 103 // g
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 111 // o
	H = H + 1
	heap[int(H)] = 102 // f
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 85 // U
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 116 // t
	H = H + 1
	heap[int(H)] = 114 // r
	H = H + 1
	heap[int(H)] = 111 // o
	H = H + 1
	heap[int(H)] = 110 // n
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t7)] = t8
	t7 = t7 + 1
	// PRIMITIVO: MR. DEEDS

	t9 = H
	heap[int(H)] = 77 // M
	H = H + 1
	heap[int(H)] = 114 // r
	H = H + 1
	heap[int(H)] = 46 // .
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 68 // D
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 100 // d
	H = H + 1
	heap[int(H)] = 115 // s
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t7)] = t9
	t7 = t7 + 1
	// PRIMITIVO: BATMAN: THE DARK KNIGHT

	t10 = H
	heap[int(H)] = 66 // B
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 116 // t
	H = H + 1
	heap[int(H)] = 109 // m
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 110 // n
	H = H + 1
	heap[int(H)] = 58 // :
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 84 // T
	H = H + 1
	heap[int(H)] = 104 // h
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 68 // D
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 114 // r
	H = H + 1
	heap[int(H)] = 107 // k
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 75 // K
	H = H + 1
	heap[int(H)] = 110 // n
	H = H + 1
	heap[int(H)] = 105 // i
	H = H + 1
	heap[int(H)] = 103 // g
	H = H + 1
	heap[int(H)] = 104 // h
	H = H + 1
	heap[int(H)] = 116 // t
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t7)] = t10
	t7 = t7 + 1
	// PRIMITIVO: MARLEY & ME

	t11 = H
	heap[int(H)] = 77 // M
	H = H + 1
	heap[int(H)] = 97 // a
	H = H + 1
	heap[int(H)] = 114 // r
	H = H + 1
	heap[int(H)] = 108 // l
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = 121 // y
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 38 // &
	H = H + 1
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 77 // M
	H = H + 1
	heap[int(H)] = 101 // e
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	heap[int(t7)] = t11
	t7 = t7 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	// -----------------------------------------------------------
	// ASIGNACION: PELICULAS
	stack[int(1)] = t6
	// FIN ASIGNACION DE PELICULAS
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// LLAMADO A FUNCION: CONTRATOS
	P = P + 2
	contratos()
	t134 = stack[int(P)]
	P = P - 2
	// RECUPERAR TEMPORALES

	t135 = P + 2
	t134 = stack[int(t135)]

	// FIN DE RECUPERACION
	// FIN LLAMADO A FUNCION
	// -----------------------------------------------------------

}
