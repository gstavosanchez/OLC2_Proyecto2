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
var t111, t112 float64

// -----------------------------------------------------------
// FUNCIONES NATIVAS
func indexValidate() {
	t14 = P + 1
	t15 = stack[int(t14)]
	t14 = t14 + 1
	t16 = stack[int(t14)]
	if t16 < 1 {
		goto L6
	}
	if t16 > t15 {
		goto L6
	}
	goto L7
L6:
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
	t17 = 1
	goto L5
L7:
	t17 = 0
L5:
	stack[int(P)] = t17
	return
}
func printStr() {
	t106 = P + 1 // Puntero del parametro
	t107 = stack[int(t106)]
L21:
	t108 = heap[int(t107)]
	if t108 == -1 {
		goto L20
	}
	fmt.Printf("%c", int(t108))
	t107 = t107 + 1 //aumentar el contador del heap
	goto L21
L20:
	return
}

// -----------------------------------------------------------
// FUNCIONES
func insertionSort() {
	// -----------------------------------------------------------
	// INICIO FOR
	// PRIMITIVO: 2
	// PRIMITIVO: 6
	t0 = P + 2
	stack[int(t0)] = 2

L1:
	t2 = P + 2 //Posicionar el puntero
	t1 = stack[int(t2)]
	if t1 > 6 {
		goto L2
	}
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: I
	t4 = P + 2 //Posicionar el puntero
	t3 = stack[int(t4)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ASIGNACION: J
	t5 = P + 3
	stack[int(t5)] = t3
	// FIN ASIGNACION DE J
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: I
	t7 = P + 2 //Posicionar el puntero
	t6 = stack[int(t7)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARR"
	t11 = P + 1
	t8 = stack[int(t11)]
	t10 = heap[int(t8)]
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: I
	t13 = P + 2 //Posicionar el puntero
	t12 = stack[int(t13)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T10
	// T12
	t18 = P + 4
	t18 = t18 + 1
	stack[int(t18)] = t10
	t18 = t18 + 1
	stack[int(t18)] = t12
	P = P + 4
	indexValidate()
	t19 = stack[int(P)]
	P = P - 4
	// -----------------------------------------------------------

	if t19 == 1 {
		goto L4
	}
	t8 = t8 + 1 //saltar len
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: I
	t21 = P + 2 //Posicionar el puntero
	t20 = stack[int(t21)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t20 = t20 - 1
	t8 = t8 + t20
	t9 = heap[int(t8)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L4:
	// -----------------------------------------------------------
	// ASIGNACION: TEMP
	t22 = P + 4
	stack[int(t22)] = t9
	// FIN ASIGNACION DE TEMP
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// WHILE
L8:
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: J
	t24 = P + 3 //Posicionar el puntero
	t23 = stack[int(t24)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// PRIMITIVO: 1
	if t23 > 1 {
		goto L11
	}
	goto L10
L11:
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: J
	t26 = P + 3 //Posicionar el puntero
	t25 = stack[int(t26)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// PRIMITIVO: 1
	t27 = t25 - 1

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARR"
	t31 = P + 1
	t28 = stack[int(t31)]
	t30 = heap[int(t28)]
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: J
	t33 = P + 3 //Posicionar el puntero
	t32 = stack[int(t33)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// PRIMITIVO: 1
	t34 = t32 - 1

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T30
	// T34
	t35 = P + 5
	t35 = t35 + 1
	stack[int(t35)] = t30
	t35 = t35 + 1
	stack[int(t35)] = t34
	P = P + 5
	indexValidate()
	t36 = stack[int(P)]
	P = P - 5
	// -----------------------------------------------------------

	if t36 == 1 {
		goto L12
	}
	t28 = t28 + 1 //saltar len
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: J
	t38 = P + 3 //Posicionar el puntero
	t37 = stack[int(t38)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// PRIMITIVO: 1
	t39 = t37 - 1

	t39 = t39 - 1
	t28 = t28 + t39
	t29 = heap[int(t28)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L12:
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: TEMP
	t41 = P + 4 //Posicionar el puntero
	t40 = stack[int(t41)]
	// FIN ACCESO
	// -----------------------------------------------------------

	if t29 > t40 {
		goto L9
	}
	goto L10
L9:
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: J
	t43 = P + 3 //Posicionar el puntero
	t42 = stack[int(t43)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: J
	t45 = P + 3 //Posicionar el puntero
	t44 = stack[int(t45)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// PRIMITIVO: 1
	t46 = t44 - 1

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARR"
	t50 = P + 1
	t47 = stack[int(t50)]
	t49 = heap[int(t47)]
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: J
	t52 = P + 3 //Posicionar el puntero
	t51 = stack[int(t52)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// PRIMITIVO: 1
	t53 = t51 - 1

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T49
	// T53
	t54 = P + 5
	t54 = t54 + 1
	stack[int(t54)] = t49
	t54 = t54 + 1
	stack[int(t54)] = t53
	P = P + 5
	indexValidate()
	t55 = stack[int(P)]
	P = P - 5
	// -----------------------------------------------------------

	if t55 == 1 {
		goto L13
	}
	t47 = t47 + 1 //saltar len
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: J
	t57 = P + 3 //Posicionar el puntero
	t56 = stack[int(t57)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// PRIMITIVO: 1
	t58 = t56 - 1

	t58 = t58 - 1
	t47 = t47 + t58
	t48 = heap[int(t47)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L13:

	// -----------------------------------------------------------
	// ASIGNAR ACCESO ARRAY "ARR"
	t61 = P + 1
	t59 = stack[int(t61)]
	t60 = heap[int(t59)]
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: J
	t63 = P + 3 //Posicionar el puntero
	t62 = stack[int(t63)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	t64 = P + 5
	t64 = t64 + 1
	stack[int(t64)] = t60
	t64 = t64 + 1
	stack[int(t64)] = t62
	P = P + 5
	indexValidate()
	t65 = stack[int(P)]
	P = P - 5
	// -----------------------------------------------------------

	if t65 == 1 {
		goto L14
	}
	t59 = t59 + 1 //saltar len
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: J
	t67 = P + 3 //Posicionar el puntero
	t66 = stack[int(t67)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t66 = t66 - 1
	t59 = t59 + t66
	heap[int(t59)] = t48
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L14:
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: J
	t69 = P + 3 //Posicionar el puntero
	t68 = stack[int(t69)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// PRIMITIVO: 1
	t70 = t68 - 1

	// -----------------------------------------------------------
	// ASIGNACION: J
	t71 = P + 3
	stack[int(t71)] = t70
	// FIN ASIGNACION DE J
	// -----------------------------------------------------------

	goto L8
L10:
	// FIN WHILE
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: J
	t73 = P + 3 //Posicionar el puntero
	t72 = stack[int(t73)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: TEMP
	t75 = P + 4 //Posicionar el puntero
	t74 = stack[int(t75)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ASIGNAR ACCESO ARRAY "ARR"
	t78 = P + 1
	t76 = stack[int(t78)]
	t77 = heap[int(t76)]
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: J
	t80 = P + 3 //Posicionar el puntero
	t79 = stack[int(t80)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	t81 = P + 5
	t81 = t81 + 1
	stack[int(t81)] = t77
	t81 = t81 + 1
	stack[int(t81)] = t79
	P = P + 5
	indexValidate()
	t82 = stack[int(P)]
	P = P - 5
	// -----------------------------------------------------------

	if t82 == 1 {
		goto L15
	}
	t76 = t76 + 1 //saltar len
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: J
	t84 = P + 3 //Posicionar el puntero
	t83 = stack[int(t84)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t83 = t83 - 1
	t76 = t76 + t83
	heap[int(t76)] = t74
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L15:
	goto L3

L3:
	t86 = P + 2 //Posicionar el puntero
	t85 = stack[int(t86)]
	t85 = t85 + 1
	stack[int(t86)] = t85
	goto L1
L2:
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
	t87 = H            // Apuntador donde se guarda el array
	t88 = H            // Contador del heap
	H = H + 7          //Almacenar espacio en heap longitud + array
	heap[int(t88)] = 6 // Guadar longtud arreglo
	t88 = t88 + 1
	// GUARDAR VALORES
	// PRIMITIVO: 100
	heap[int(t88)] = 100
	t88 = t88 + 1
	// PRIMITIVO: 50
	heap[int(t88)] = 50
	t88 = t88 + 1
	// PRIMITIVO: 25
	heap[int(t88)] = 25
	t88 = t88 + 1
	// PRIMITIVO: 150
	heap[int(t88)] = 150
	t88 = t88 + 1
	// PRIMITIVO: 1
	heap[int(t88)] = 1
	t88 = t88 + 1
	// PRIMITIVO: 5
	heap[int(t88)] = 5
	t88 = t88 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	// -----------------------------------------------------------
	// ASIGNACION: ARREGLO
	stack[int(0)] = t87
	// FIN ASIGNACION DE ARREGLO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// LLAMADO A FUNCION: INSERTIONSORT
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: ARREGLO
	t89 = stack[int(0)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t90 = P + 2
	stack[int(t90)] = t89
	P = P + 1
	insertionSort()
	t90 = stack[int(P)]
	P = P - 1
	// FIN LLAMADO A FUNCION
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO FOR
	// PRIMITIVO: 1
	// PRIMITIVO: 6
	t91 = P + 1
	stack[int(t91)] = 1

L16:
	t93 = P + 1 //Posicionar el puntero
	t92 = stack[int(t93)]
	if t92 > 6 {
		goto L17
	}

	// -----------------------------------------------------------
	// INICIO PRINT
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: X
	t95 = P + 1 //Posicionar el puntero
	t94 = stack[int(t95)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARREGLO"
	t96 = stack[int(0)]
	t98 = heap[int(t96)]
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: X
	t100 = P + 1 //Posicionar el puntero
	t99 = stack[int(t100)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T98
	// T99
	t101 = P + 2
	t101 = t101 + 1
	stack[int(t101)] = t98
	t101 = t101 + 1
	stack[int(t101)] = t99
	P = P + 2
	indexValidate()
	t102 = stack[int(P)]
	P = P - 2
	// -----------------------------------------------------------

	if t102 == 1 {
		goto L19
	}
	t96 = t96 + 1 //saltar len
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: X
	t104 = P + 1 //Posicionar el puntero
	t103 = stack[int(t104)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t103 = t103 - 1
	t96 = t96 + t103
	t97 = heap[int(t96)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L19:
	fmt.Printf("%d", int(t97))
	// PRIMITIVO:

	t105 = H
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t109 = P + 2
	t109 = t109 + 1
	stack[int(t109)] = t105
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t110 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	// FIN PRINT
	// -----------------------------------------------------------
	goto L18

L18:
	t112 = P + 1 //Posicionar el puntero
	t111 = stack[int(t112)]
	t111 = t111 + 1
	stack[int(t112)] = t111
	goto L16
L17:
	// FIN FOR
	// -----------------------------------------------------------

}
