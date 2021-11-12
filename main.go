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
var t131, t132, t133, t134, t135, t136, t137, t138, t139, t140 float64
var t141, t142, t143, t144, t145, t146, t147, t148, t149, t150 float64
var t151, t152, t153, t154, t155, t156, t157, t158, t159, t160 float64
var t161, t162, t163, t164, t165, t166, t167, t168, t169, t170 float64
var t171, t172, t173, t174, t175, t176, t177, t178, t179, t180 float64
var t181, t182, t183, t184, t185, t186, t187, t188, t189, t190 float64
var t191, t192, t193, t194, t195, t196, t197, t198, t199, t200 float64
var t201, t202, t203, t204, t205, t206, t207, t208, t209, t210 float64
var t211, t212, t213, t214, t215, t216, t217, t218, t219, t220 float64
var t221, t222, t223, t224, t225, t226, t227, t228, t229, t230 float64
var t231, t232, t233, t234, t235, t236, t237, t238, t239, t240 float64
var t241, t242, t243, t244, t245, t246, t247, t248, t249, t250 float64
var t251, t252, t253, t254, t255, t256, t257, t258, t259, t260 float64
var t261, t262, t263, t264, t265, t266, t267, t268, t269, t270 float64
var t271, t272, t273, t274, t275, t276, t277, t278, t279, t280 float64
var t281, t282, t283, t284, t285, t286, t287, t288, t289, t290 float64
var t291, t292, t293, t294, t295, t296, t297, t298, t299, t300 float64
var t301, t302, t303, t304, t305, t306, t307, t308, t309, t310 float64
var t311, t312, t313, t314, t315, t316, t317, t318, t319, t320 float64
var t321, t322, t323, t324, t325, t326, t327, t328, t329, t330 float64
var t331, t332, t333, t334, t335, t336, t337, t338, t339, t340 float64
var t341, t342, t343, t344, t345, t346, t347, t348, t349, t350 float64
var t351, t352, t353, t354, t355, t356, t357, t358, t359, t360 float64
var t361, t362, t363, t364, t365, t366, t367, t368, t369, t370 float64
var t371, t372, t373 float64

// -----------------------------------------------------------
// FUNCIONES NATIVAS
func divValidate() {
	t15 = P + 1
	t16 = stack[int(t15)]
	t15 = t15 + 1
	t17 = stack[int(t15)]
	if t17 != 0 {
		goto L4
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
	t18 = 0 // resultado incorrecto
	goto L3
L4:
	t18 = t16 / t17
L3:
	stack[int(P)] = t18 // gurdar resultado
	return
}
func truncNum() {
	t21 = P + 1
	t22 = stack[int(t21)]
	t23 = 0
L6:
	if t23 > t22 {
		goto L5
	}
	t23 = t23 + 1
	goto L6
L5:
	t23 = t23 - 1
	stack[int(P)] = t23 // gurdar resultado
	return
}
func indexValidate() {
	t39 = P + 1
	t40 = stack[int(t39)]
	t39 = t39 + 1
	t41 = stack[int(t39)]
	if t41 < 1 {
		goto L9
	}
	if t41 > t40 {
		goto L9
	}
	goto L10
L9:
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
	t42 = 1
	goto L8
L10:
	t42 = 0
L8:
	stack[int(P)] = t42
	return
}
func printStr() {
	t212 = P + 1 // Puntero del parametro
	t213 = stack[int(t212)]
L39:
	t214 = heap[int(t213)]
	if t214 == -1 {
		goto L38
	}
	fmt.Printf("%c", int(t214))
	t213 = t213 + 1 //aumentar el contador del heap
	goto L39
L38:
	return
}

// -----------------------------------------------------------
// FUNCIONES
func quicksort() {
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LOW
	t1 = P + 2 //Posicionar el puntero
	t0 = stack[int(t1)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ASIGNACION: LO
	t2 = P + 4
	stack[int(t2)] = t0
	// FIN ASIGNACION DE LO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: N
	t4 = P + 3 //Posicionar el puntero
	t3 = stack[int(t4)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ASIGNACION: HI
	t5 = P + 5
	stack[int(t5)] = t3
	// FIN ASIGNACION DE HI
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO IF
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t7 = P + 4 //Posicionar el puntero
	t6 = stack[int(t7)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: N
	t9 = P + 3 //Posicionar el puntero
	t8 = stack[int(t9)]
	// FIN ACCESO
	// -----------------------------------------------------------

	if t6 >= t8 {
		goto L1
	}
	goto L2
L1:
	// PRIMITIVO: 0
	stack[int(P)] = 0
	goto L0
L2:
	// FIN IF
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t11 = P + 4 //Posicionar el puntero
	t10 = stack[int(t11)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: HI
	t13 = P + 5 //Posicionar el puntero
	t12 = stack[int(t13)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t14 = t10 + t12

	// PRIMITIVO: 2
	// -----------------------------------------------------------
	// PASO DE PARAMETROS
	t19 = P + 6
	// 1ER PARAMETRO
	t19 = t19 + 1
	stack[int(t19)] = t14
	// 2DO PARAMETRO
	t19 = t19 + 1
	stack[int(t19)] = 2
	// FIN PASO DE PARAMETROS
	// -----------------------------------------------------------
	// CAMBIO DE ENTORNO
	P = P + 6
	divValidate()
	// GUARDAR EL RETURN DE LA FUNCION
	t20 = stack[int(P)]
	P = P - 6

	// -----------------------------------------------------------
	t24 = P + 6
	t24 = t24 + 1
	stack[int(t24)] = t20
	P = P + 6
	truncNum()
	t25 = stack[int(P)]
	P = P - 6
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARRAY"
	t29 = P + 1
	t26 = stack[int(t29)]
	t28 = heap[int(t26)]
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t31 = P + 4 //Posicionar el puntero
	t30 = stack[int(t31)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: HI
	t33 = P + 5 //Posicionar el puntero
	t32 = stack[int(t33)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t34 = t30 + t32

	// PRIMITIVO: 2
	// -----------------------------------------------------------
	// PASO DE PARAMETROS
	t35 = P + 6
	// 1ER PARAMETRO
	t35 = t35 + 1
	stack[int(t35)] = t34
	// 2DO PARAMETRO
	t35 = t35 + 1
	stack[int(t35)] = 2
	// FIN PASO DE PARAMETROS
	// -----------------------------------------------------------
	// CAMBIO DE ENTORNO
	P = P + 6
	divValidate()
	// GUARDAR EL RETURN DE LA FUNCION
	t36 = stack[int(P)]
	P = P - 6

	// -----------------------------------------------------------
	t37 = P + 6
	t37 = t37 + 1
	stack[int(t37)] = t36
	P = P + 6
	truncNum()
	t38 = stack[int(P)]
	P = P - 6
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T28
	// T38
	t43 = P + 6
	t43 = t43 + 1
	stack[int(t43)] = t28
	t43 = t43 + 1
	stack[int(t43)] = t38
	P = P + 6
	indexValidate()
	t44 = stack[int(P)]
	P = P - 6
	// -----------------------------------------------------------

	if t44 == 1 {
		goto L7
	}
	t26 = t26 + 1 //saltar len
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t46 = P + 4 //Posicionar el puntero
	t45 = stack[int(t46)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: HI
	t48 = P + 5 //Posicionar el puntero
	t47 = stack[int(t48)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t49 = t45 + t47

	// PRIMITIVO: 2
	// -----------------------------------------------------------
	// PASO DE PARAMETROS
	t50 = P + 6
	// 1ER PARAMETRO
	t50 = t50 + 1
	stack[int(t50)] = t49
	// 2DO PARAMETRO
	t50 = t50 + 1
	stack[int(t50)] = 2
	// FIN PASO DE PARAMETROS
	// -----------------------------------------------------------
	// CAMBIO DE ENTORNO
	P = P + 6
	divValidate()
	// GUARDAR EL RETURN DE LA FUNCION
	t51 = stack[int(P)]
	P = P - 6

	// -----------------------------------------------------------
	t52 = P + 6
	t52 = t52 + 1
	stack[int(t52)] = t51
	P = P + 6
	truncNum()
	t53 = stack[int(P)]
	P = P - 6
	// -----------------------------------------------------------

	t53 = t53 - 1
	t26 = t26 + t53
	t27 = heap[int(t26)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L7:
	// -----------------------------------------------------------
	// ASIGNACION: MID
	t54 = P + 6
	stack[int(t54)] = t27
	// FIN ASIGNACION DE MID
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// WHILE
L11:
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t56 = P + 4 //Posicionar el puntero
	t55 = stack[int(t56)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: HI
	t58 = P + 5 //Posicionar el puntero
	t57 = stack[int(t58)]
	// FIN ACCESO
	// -----------------------------------------------------------

	if t55 < t57 {
		goto L12
	}
	goto L13
L12:
	// -----------------------------------------------------------
	// WHILE
L14:
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t60 = P + 4 //Posicionar el puntero
	t59 = stack[int(t60)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: HI
	t62 = P + 5 //Posicionar el puntero
	t61 = stack[int(t62)]
	// FIN ACCESO
	// -----------------------------------------------------------

	if t59 < t61 {
		goto L17
	}
	goto L16
L17:
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t64 = P + 4 //Posicionar el puntero
	t63 = stack[int(t64)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARRAY"
	t68 = P + 1
	t65 = stack[int(t68)]
	t67 = heap[int(t65)]
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t70 = P + 4 //Posicionar el puntero
	t69 = stack[int(t70)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T67
	// T69
	t71 = P + 7
	t71 = t71 + 1
	stack[int(t71)] = t67
	t71 = t71 + 1
	stack[int(t71)] = t69
	P = P + 7
	indexValidate()
	t72 = stack[int(P)]
	P = P - 7
	// -----------------------------------------------------------

	if t72 == 1 {
		goto L18
	}
	t65 = t65 + 1 //saltar len
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t74 = P + 4 //Posicionar el puntero
	t73 = stack[int(t74)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t73 = t73 - 1
	t65 = t65 + t73
	t66 = heap[int(t65)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L18:
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: MID
	t76 = P + 6 //Posicionar el puntero
	t75 = stack[int(t76)]
	// FIN ACCESO
	// -----------------------------------------------------------

	if t66 < t75 {
		goto L15
	}
	goto L16
L15:
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t78 = P + 4 //Posicionar el puntero
	t77 = stack[int(t78)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// PRIMITIVO: 1
	t79 = t77 + 1

	// -----------------------------------------------------------
	// ASIGNACION: LO
	t80 = P + 4
	stack[int(t80)] = t79
	// FIN ASIGNACION DE LO
	// -----------------------------------------------------------

	goto L14
L16:
	// FIN WHILE
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// WHILE
L19:
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t82 = P + 4 //Posicionar el puntero
	t81 = stack[int(t82)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: HI
	t84 = P + 5 //Posicionar el puntero
	t83 = stack[int(t84)]
	// FIN ACCESO
	// -----------------------------------------------------------

	if t81 < t83 {
		goto L22
	}
	goto L21
L22:
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: HI
	t86 = P + 5 //Posicionar el puntero
	t85 = stack[int(t86)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARRAY"
	t90 = P + 1
	t87 = stack[int(t90)]
	t89 = heap[int(t87)]
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: HI
	t92 = P + 5 //Posicionar el puntero
	t91 = stack[int(t92)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T89
	// T91
	t93 = P + 7
	t93 = t93 + 1
	stack[int(t93)] = t89
	t93 = t93 + 1
	stack[int(t93)] = t91
	P = P + 7
	indexValidate()
	t94 = stack[int(P)]
	P = P - 7
	// -----------------------------------------------------------

	if t94 == 1 {
		goto L23
	}
	t87 = t87 + 1 //saltar len
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: HI
	t96 = P + 5 //Posicionar el puntero
	t95 = stack[int(t96)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t95 = t95 - 1
	t87 = t87 + t95
	t88 = heap[int(t87)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L23:
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: MID
	t98 = P + 6 //Posicionar el puntero
	t97 = stack[int(t98)]
	// FIN ACCESO
	// -----------------------------------------------------------

	if t88 > t97 {
		goto L20
	}
	goto L21
L20:
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: HI
	t100 = P + 5 //Posicionar el puntero
	t99 = stack[int(t100)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// PRIMITIVO: 1
	t101 = t99 - 1

	// -----------------------------------------------------------
	// ASIGNACION: HI
	t102 = P + 5
	stack[int(t102)] = t101
	// FIN ASIGNACION DE HI
	// -----------------------------------------------------------

	goto L19
L21:
	// FIN WHILE
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO IF
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t104 = P + 4 //Posicionar el puntero
	t103 = stack[int(t104)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: HI
	t106 = P + 5 //Posicionar el puntero
	t105 = stack[int(t106)]
	// FIN ACCESO
	// -----------------------------------------------------------

	if t103 < t105 {
		goto L24
	}
	goto L25
L24:
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t108 = P + 4 //Posicionar el puntero
	t107 = stack[int(t108)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARRAY"
	t112 = P + 1
	t109 = stack[int(t112)]
	t111 = heap[int(t109)]
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t114 = P + 4 //Posicionar el puntero
	t113 = stack[int(t114)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T111
	// T113
	t115 = P + 7
	t115 = t115 + 1
	stack[int(t115)] = t111
	t115 = t115 + 1
	stack[int(t115)] = t113
	P = P + 7
	indexValidate()
	t116 = stack[int(P)]
	P = P - 7
	// -----------------------------------------------------------

	if t116 == 1 {
		goto L26
	}
	t109 = t109 + 1 //saltar len
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t118 = P + 4 //Posicionar el puntero
	t117 = stack[int(t118)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t117 = t117 - 1
	t109 = t109 + t117
	t110 = heap[int(t109)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L26:
	// -----------------------------------------------------------
	// ASIGNACION: T
	t119 = P + 7
	stack[int(t119)] = t110
	// FIN ASIGNACION DE T
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t121 = P + 4 //Posicionar el puntero
	t120 = stack[int(t121)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: HI
	t123 = P + 5 //Posicionar el puntero
	t122 = stack[int(t123)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARRAY"
	t127 = P + 1
	t124 = stack[int(t127)]
	t126 = heap[int(t124)]
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: HI
	t129 = P + 5 //Posicionar el puntero
	t128 = stack[int(t129)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T126
	// T128
	t130 = P + 8
	t130 = t130 + 1
	stack[int(t130)] = t126
	t130 = t130 + 1
	stack[int(t130)] = t128
	P = P + 8
	indexValidate()
	t131 = stack[int(P)]
	P = P - 8
	// -----------------------------------------------------------

	if t131 == 1 {
		goto L27
	}
	t124 = t124 + 1 //saltar len
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: HI
	t133 = P + 5 //Posicionar el puntero
	t132 = stack[int(t133)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t132 = t132 - 1
	t124 = t124 + t132
	t125 = heap[int(t124)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L27:

	// -----------------------------------------------------------
	// ASIGNAR ACCESO ARRAY "ARRAY"
	t136 = P + 1
	t134 = stack[int(t136)]
	t135 = heap[int(t134)]
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t138 = P + 4 //Posicionar el puntero
	t137 = stack[int(t138)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	t139 = P + 8
	t139 = t139 + 1
	stack[int(t139)] = t135
	t139 = t139 + 1
	stack[int(t139)] = t137
	P = P + 8
	indexValidate()
	t140 = stack[int(P)]
	P = P - 8
	// -----------------------------------------------------------

	if t140 == 1 {
		goto L28
	}
	t134 = t134 + 1 //saltar len
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t142 = P + 4 //Posicionar el puntero
	t141 = stack[int(t142)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t141 = t141 - 1
	t134 = t134 + t141
	heap[int(t134)] = t125
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L28:
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: HI
	t144 = P + 5 //Posicionar el puntero
	t143 = stack[int(t144)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: T
	t146 = P + 7 //Posicionar el puntero
	t145 = stack[int(t146)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ASIGNAR ACCESO ARRAY "ARRAY"
	t149 = P + 1
	t147 = stack[int(t149)]
	t148 = heap[int(t147)]
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: HI
	t151 = P + 5 //Posicionar el puntero
	t150 = stack[int(t151)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	t152 = P + 8
	t152 = t152 + 1
	stack[int(t152)] = t148
	t152 = t152 + 1
	stack[int(t152)] = t150
	P = P + 8
	indexValidate()
	t153 = stack[int(P)]
	P = P - 8
	// -----------------------------------------------------------

	if t153 == 1 {
		goto L29
	}
	t147 = t147 + 1 //saltar len
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: HI
	t155 = P + 5 //Posicionar el puntero
	t154 = stack[int(t155)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t154 = t154 - 1
	t147 = t147 + t154
	heap[int(t147)] = t145
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L29:
L25:
	// FIN IF
	// -----------------------------------------------------------

	goto L11
L13:
	// FIN WHILE
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO IF
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: HI
	t157 = P + 5 //Posicionar el puntero
	t156 = stack[int(t157)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t159 = P + 4 //Posicionar el puntero
	t158 = stack[int(t159)]
	// FIN ACCESO
	// -----------------------------------------------------------

	if t156 < t158 {
		goto L30
	}
	goto L31
L30:
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: HI
	t161 = P + 5 //Posicionar el puntero
	t160 = stack[int(t161)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ASIGNACION: T
	t162 = P + 7
	stack[int(t162)] = t160
	// FIN ASIGNACION DE T
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t164 = P + 4 //Posicionar el puntero
	t163 = stack[int(t164)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ASIGNACION: HI
	t165 = P + 5
	stack[int(t165)] = t163
	// FIN ASIGNACION DE HI
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: T
	t167 = P + 7 //Posicionar el puntero
	t166 = stack[int(t167)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ASIGNACION: LO
	t168 = P + 4
	stack[int(t168)] = t166
	// FIN ASIGNACION DE LO
	// -----------------------------------------------------------

L31:
	// FIN IF
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// LLAMADO A FUNCION: QUICKSORT
	// GUARDAR TEMPORALES

	t169 = P + 8
	stack[int(t169)] = t25
	t169 = t169 + 1
	stack[int(t169)] = t63
	t169 = t169 + 1
	stack[int(t169)] = t85
	t169 = t169 + 1
	stack[int(t169)] = t107
	t169 = t169 + 1
	stack[int(t169)] = t120
	t169 = t169 + 1
	stack[int(t169)] = t122
	t169 = t169 + 1
	stack[int(t169)] = t143

	// FIN DE GUARDAR TEMPORALES
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: ARRAY
	t171 = P + 1 //Posicionar el puntero
	t170 = stack[int(t171)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LOW
	t173 = P + 2 //Posicionar el puntero
	t172 = stack[int(t173)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t175 = P + 4 //Posicionar el puntero
	t174 = stack[int(t175)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t176 = P + 16
	stack[int(t176)] = t170
	t176 = t176 + 1
	stack[int(t176)] = t172
	t176 = t176 + 1
	stack[int(t176)] = t174
	P = P + 15
	quicksort()
	t176 = stack[int(P)]
	P = P - 15
	// RECUPERAR TEMPORALES

	t177 = P + 8
	t25 = stack[int(t177)]
	t177 = t177 + 1
	t63 = stack[int(t177)]
	t177 = t177 + 1
	t85 = stack[int(t177)]
	t177 = t177 + 1
	t107 = stack[int(t177)]
	t177 = t177 + 1
	t120 = stack[int(t177)]
	t177 = t177 + 1
	t122 = stack[int(t177)]
	t177 = t177 + 1
	t143 = stack[int(t177)]

	// FIN DE RECUPERACION
	// FIN LLAMADO A FUNCION
	// -----------------------------------------------------------

	// PRIMITIVO: 0
	// -----------------------------------------------------------
	// ASIGNACION: COND
	t178 = P + 8
	stack[int(t178)] = 0
	// FIN ASIGNACION DE COND
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO IF
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t180 = P + 4 //Posicionar el puntero
	t179 = stack[int(t180)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LOW
	t182 = P + 2 //Posicionar el puntero
	t181 = stack[int(t182)]
	// FIN ACCESO
	// -----------------------------------------------------------

	if t179 == t181 {
		goto L32
	}
	goto L33
L32:
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t184 = P + 4 //Posicionar el puntero
	t183 = stack[int(t184)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// PRIMITIVO: 1
	t185 = t183 + 1

	// -----------------------------------------------------------
	// ASIGNACION: COND
	t186 = P + 8
	stack[int(t186)] = t185
	// FIN ASIGNACION DE COND
	// -----------------------------------------------------------

	goto L34
L33:
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: LO
	t188 = P + 4 //Posicionar el puntero
	t187 = stack[int(t188)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ASIGNACION: COND
	t189 = P + 8
	stack[int(t189)] = t187
	// FIN ASIGNACION DE COND
	// -----------------------------------------------------------

L34:
	// FIN IF
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// LLAMADO A FUNCION: QUICKSORT
	// GUARDAR TEMPORALES

	t190 = P + 9
	stack[int(t190)] = t25
	t190 = t190 + 1
	stack[int(t190)] = t63
	t190 = t190 + 1
	stack[int(t190)] = t85
	t190 = t190 + 1
	stack[int(t190)] = t107
	t190 = t190 + 1
	stack[int(t190)] = t120
	t190 = t190 + 1
	stack[int(t190)] = t122
	t190 = t190 + 1
	stack[int(t190)] = t143

	// FIN DE GUARDAR TEMPORALES
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: ARRAY
	t192 = P + 1 //Posicionar el puntero
	t191 = stack[int(t192)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: COND
	t194 = P + 8 //Posicionar el puntero
	t193 = stack[int(t194)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO A VARIABLE: N
	t196 = P + 3 //Posicionar el puntero
	t195 = stack[int(t196)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t197 = P + 17
	stack[int(t197)] = t191
	t197 = t197 + 1
	stack[int(t197)] = t193
	t197 = t197 + 1
	stack[int(t197)] = t195
	P = P + 16
	quicksort()
	t197 = stack[int(P)]
	P = P - 16
	// RECUPERAR TEMPORALES

	t198 = P + 9
	t25 = stack[int(t198)]
	t198 = t198 + 1
	t63 = stack[int(t198)]
	t198 = t198 + 1
	t85 = stack[int(t198)]
	t198 = t198 + 1
	t107 = stack[int(t198)]
	t198 = t198 + 1
	t120 = stack[int(t198)]
	t198 = t198 + 1
	t122 = stack[int(t198)]
	t198 = t198 + 1
	t143 = stack[int(t198)]

	// FIN DE RECUPERACION
	// FIN LLAMADO A FUNCION
	// -----------------------------------------------------------

L0:
	return
}

// ===========================================================
// MAIN
// ===========================================================
func main() {
	P = 0 // Puntero Stack
	H = 0 //Puntero Heap
	// PRIMITIVO: 0
	// -----------------------------------------------------------
	// ASIGNACION: I
	stack[int(0)] = 0
	// FIN ASIGNACION DE I
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ARREGLO
	t199 = H            // Apuntador donde se guarda el array
	t200 = H            // Contador del heap
	H = H + 3           //Almacenar espacio en heap longitud + array
	heap[int(t200)] = 2 // Guadar longtud arreglo
	t200 = t200 + 1
	// GUARDAR VALORES

	// -----------------------------------------------------------
	// ARREGLO
	t201 = H             // Apuntador donde se guarda el array
	t202 = H             // Contador del heap
	H = H + 14           //Almacenar espacio en heap longitud + array
	heap[int(t202)] = 13 // Guadar longtud arreglo
	t202 = t202 + 1
	// GUARDAR VALORES
	// PRIMITIVO: 12
	heap[int(t202)] = 12
	t202 = t202 + 1
	// PRIMITIVO: 9
	heap[int(t202)] = 9
	t202 = t202 + 1
	// PRIMITIVO: 4
	heap[int(t202)] = 4
	t202 = t202 + 1
	// PRIMITIVO: 99
	heap[int(t202)] = 99
	t202 = t202 + 1
	// PRIMITIVO: 56
	heap[int(t202)] = 56
	t202 = t202 + 1
	// PRIMITIVO: 34
	heap[int(t202)] = 34
	t202 = t202 + 1
	// PRIMITIVO: 78
	heap[int(t202)] = 78
	t202 = t202 + 1
	// PRIMITIVO: 22
	heap[int(t202)] = 22
	t202 = t202 + 1
	// PRIMITIVO: 1
	heap[int(t202)] = 1
	t202 = t202 + 1
	// PRIMITIVO: 3
	heap[int(t202)] = 3
	t202 = t202 + 1
	// PRIMITIVO: 10
	heap[int(t202)] = 10
	t202 = t202 + 1
	// PRIMITIVO: 13
	heap[int(t202)] = 13
	t202 = t202 + 1
	// PRIMITIVO: 120
	heap[int(t202)] = 120
	t202 = t202 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	heap[int(t200)] = t201
	t200 = t200 + 1

	// -----------------------------------------------------------
	// ARREGLO
	t203 = H             // Apuntador donde se guarda el array
	t204 = H             // Contador del heap
	H = H + 17           //Almacenar espacio en heap longitud + array
	heap[int(t204)] = 16 // Guadar longtud arreglo
	t204 = t204 + 1
	// GUARDAR VALORES
	// PRIMITIVO: 32
	heap[int(t204)] = 32
	t204 = t204 + 1
	// PRIMITIVO: 7
	// PRIMITIVO: 3
	t205 = 7 * 3

	heap[int(t204)] = t205
	t204 = t204 + 1
	// PRIMITIVO: 7
	heap[int(t204)] = 7
	t204 = t204 + 1
	// PRIMITIVO: 89
	heap[int(t204)] = 89
	t204 = t204 + 1
	// PRIMITIVO: 56
	heap[int(t204)] = 56
	t204 = t204 + 1
	// PRIMITIVO: 909
	heap[int(t204)] = 909
	t204 = t204 + 1
	// PRIMITIVO: 109
	heap[int(t204)] = 109
	t204 = t204 + 1
	// PRIMITIVO: 2
	heap[int(t204)] = 2
	t204 = t204 + 1
	// PRIMITIVO: 9
	heap[int(t204)] = 9
	t204 = t204 + 1
	// PRIMITIVO: 9874
	// PRIMITIVO: 0
	t206 = 0
	t207 = 1
	if 0 == 0 {
		goto L37
	}
L35:
	if t206 == 0 {
		goto L36
	}
	t207 = 9874 * t207
	t206 = t206 + 1
	goto L35
L37:
	t207 = 1
L36:
	heap[int(t204)] = t207
	t204 = t204 + 1
	// PRIMITIVO: 44
	heap[int(t204)] = 44
	t204 = t204 + 1
	// PRIMITIVO: 3
	heap[int(t204)] = 3
	t204 = t204 + 1
	// PRIMITIVO: 820
	// PRIMITIVO: 10
	t208 = 820 * 10

	heap[int(t204)] = t208
	t204 = t204 + 1
	// PRIMITIVO: 11
	heap[int(t204)] = 11
	t204 = t204 + 1
	// PRIMITIVO: 8
	// PRIMITIVO: 0
	t209 = 8 * 0

	// PRIMITIVO: 8
	t210 = t209 + 8

	heap[int(t204)] = t210
	t204 = t204 + 1
	// PRIMITIVO: 10
	heap[int(t204)] = 10
	t204 = t204 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	heap[int(t200)] = t203
	t200 = t200 + 1
	// -----------------------------------------------------------
	// FIN ARREGLO

	// -----------------------------------------------------------
	// ASIGNACION: ARRAY
	stack[int(1)] = t199
	// FIN ASIGNACION DE ARRAY
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: QUICK SORT

	t211 = H
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
	heap[int(H)] = 32 //
	H = H + 1
	heap[int(H)] = 83 // S
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
	t215 = P + 2
	t215 = t215 + 1
	stack[int(t215)] = t211
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t216 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: VALORES ANTES DE QUICKSORT

	t217 = H
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
	t218 = P + 2
	t218 = t218 + 1
	stack[int(t218)] = t217
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t219 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------
	// -----------------------------------------------------------
	// INICIO FOR
	// PRIMITIVO: 1
	// PRIMITIVO: 1

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARRAY"
	t220 = stack[int(1)]
	t222 = heap[int(t220)]
	// PRIMITIVO: 1

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T222
	// 1
	t223 = P + 2
	t223 = t223 + 1
	stack[int(t223)] = t222
	t223 = t223 + 1
	stack[int(t223)] = 1
	P = P + 2
	indexValidate()
	t224 = stack[int(P)]
	P = P - 2
	// -----------------------------------------------------------

	if t224 == 1 {
		goto L40
	}
	t220 = t220 + 1 //saltar len
	// PRIMITIVO: 1
	t220 = t220 + 0
	t221 = heap[int(t220)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L40:
	t221 = heap[int(t221)]
	t225 = P + 2
	stack[int(t225)] = 1

L41:
	t227 = P + 2 //Posicionar el puntero
	t226 = stack[int(t227)]
	if t226 > t221 {
		goto L42
	}

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: 1
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: X
	t229 = P + 2 //Posicionar el puntero
	t228 = stack[int(t229)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARRAY"
	t230 = stack[int(1)]
	t232 = heap[int(t230)]
	// PRIMITIVO: 1

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T232
	// 1
	t233 = P + 3
	t233 = t233 + 1
	stack[int(t233)] = t232
	t233 = t233 + 1
	stack[int(t233)] = 1
	P = P + 3
	indexValidate()
	t234 = stack[int(P)]
	P = P - 3
	// -----------------------------------------------------------

	if t234 == 1 {
		goto L44
	}
	t230 = t230 + 1 //saltar len
	// PRIMITIVO: 1
	t230 = t230 + 0
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: X
	t236 = P + 2 //Posicionar el puntero
	t235 = stack[int(t236)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t237 = heap[int(t230)]
	t238 = heap[int(t237)]

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T238
	// T235
	t239 = P + 3
	t239 = t239 + 1
	stack[int(t239)] = t238
	t239 = t239 + 1
	stack[int(t239)] = t235
	P = P + 3
	indexValidate()
	t240 = stack[int(P)]
	P = P - 3
	// -----------------------------------------------------------

	if t240 == 1 {
		goto L44
	}
	t237 = t237 + 1 //skip len
	t235 = t235 - 1
	t237 = t237 + t235
	t231 = heap[int(t237)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L44:
	fmt.Printf("%d", int(t231))
	// PRIMITIVO:

	t241 = H
	heap[int(H)] = 9 //
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t242 = P + 3
	t242 = t242 + 1
	stack[int(t242)] = t241
	// -----------------------------------------------------------

	P = P + 3
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t243 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 3

	// FIN PRINT
	// -----------------------------------------------------------
	goto L43

L43:
	t245 = P + 2 //Posicionar el puntero
	t244 = stack[int(t245)]
	t244 = t244 + 1
	stack[int(t245)] = t244
	goto L41
L42:
	// FIN FOR
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO:

	t246 = H
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t247 = P + 2
	t247 = t247 + 1
	stack[int(t247)] = t246
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t248 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: -------------------------

	t249 = H
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
	t250 = P + 2
	t250 = t250 + 1
	stack[int(t250)] = t249
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t251 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// LLAMADO A FUNCION: QUICKSORT
	// GUARDAR TEMPORALES

	t252 = P + 2
	stack[int(t252)] = t216
	t252 = t252 + 1
	stack[int(t252)] = t219
	t252 = t252 + 1
	stack[int(t252)] = t228
	t252 = t252 + 1
	stack[int(t252)] = t243
	t252 = t252 + 1
	stack[int(t252)] = t248
	t252 = t252 + 1
	stack[int(t252)] = t251

	// FIN DE GUARDAR TEMPORALES
	// PRIMITIVO: 1

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARRAY"
	t253 = stack[int(1)]
	t255 = heap[int(t253)]
	// PRIMITIVO: 1

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T255
	// 1
	t256 = P + 8
	t256 = t256 + 1
	stack[int(t256)] = t255
	t256 = t256 + 1
	stack[int(t256)] = 1
	P = P + 8
	indexValidate()
	t257 = stack[int(P)]
	P = P - 8
	// -----------------------------------------------------------

	if t257 == 1 {
		goto L45
	}
	t253 = t253 + 1 //saltar len
	// PRIMITIVO: 1
	t253 = t253 + 0
	t254 = heap[int(t253)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L45:
	// PRIMITIVO: 1
	// PRIMITIVO: 1

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARRAY"
	t258 = stack[int(1)]
	t260 = heap[int(t258)]
	// PRIMITIVO: 1

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T260
	// 1
	t261 = P + 8
	t261 = t261 + 1
	stack[int(t261)] = t260
	t261 = t261 + 1
	stack[int(t261)] = 1
	P = P + 8
	indexValidate()
	t262 = stack[int(P)]
	P = P - 8
	// -----------------------------------------------------------

	if t262 == 1 {
		goto L46
	}
	t258 = t258 + 1 //saltar len
	// PRIMITIVO: 1
	t258 = t258 + 0
	t259 = heap[int(t258)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L46:
	t259 = heap[int(t259)]
	t263 = P + 9
	stack[int(t263)] = t254
	t263 = t263 + 1
	stack[int(t263)] = 1
	t263 = t263 + 1
	stack[int(t263)] = t259
	P = P + 8
	quicksort()
	t263 = stack[int(P)]
	P = P - 8
	// RECUPERAR TEMPORALES

	t264 = P + 2
	t216 = stack[int(t264)]
	t264 = t264 + 1
	t219 = stack[int(t264)]
	t264 = t264 + 1
	t228 = stack[int(t264)]
	t264 = t264 + 1
	t243 = stack[int(t264)]
	t264 = t264 + 1
	t248 = stack[int(t264)]
	t264 = t264 + 1
	t251 = stack[int(t264)]

	// FIN DE RECUPERACION
	// FIN LLAMADO A FUNCION
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: VALORES DESPUES DE QUICKSORT:

	t265 = H
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
	t266 = P + 2
	t266 = t266 + 1
	stack[int(t266)] = t265
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t267 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------
	// -----------------------------------------------------------
	// INICIO FOR
	// PRIMITIVO: 1
	// PRIMITIVO: 1

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARRAY"
	t268 = stack[int(1)]
	t270 = heap[int(t268)]
	// PRIMITIVO: 1

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T270
	// 1
	t271 = P + 2
	t271 = t271 + 1
	stack[int(t271)] = t270
	t271 = t271 + 1
	stack[int(t271)] = 1
	P = P + 2
	indexValidate()
	t272 = stack[int(P)]
	P = P - 2
	// -----------------------------------------------------------

	if t272 == 1 {
		goto L47
	}
	t268 = t268 + 1 //saltar len
	// PRIMITIVO: 1
	t268 = t268 + 0
	t269 = heap[int(t268)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L47:
	t269 = heap[int(t269)]
	t273 = P + 2
	stack[int(t273)] = 1

L48:
	t275 = P + 2 //Posicionar el puntero
	t274 = stack[int(t275)]
	if t274 > t269 {
		goto L49
	}

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: 1
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: Y
	t277 = P + 2 //Posicionar el puntero
	t276 = stack[int(t277)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARRAY"
	t278 = stack[int(1)]
	t280 = heap[int(t278)]
	// PRIMITIVO: 1

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T280
	// 1
	t281 = P + 3
	t281 = t281 + 1
	stack[int(t281)] = t280
	t281 = t281 + 1
	stack[int(t281)] = 1
	P = P + 3
	indexValidate()
	t282 = stack[int(P)]
	P = P - 3
	// -----------------------------------------------------------

	if t282 == 1 {
		goto L51
	}
	t278 = t278 + 1 //saltar len
	// PRIMITIVO: 1
	t278 = t278 + 0
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: Y
	t284 = P + 2 //Posicionar el puntero
	t283 = stack[int(t284)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t285 = heap[int(t278)]
	t286 = heap[int(t285)]

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T286
	// T283
	t287 = P + 3
	t287 = t287 + 1
	stack[int(t287)] = t286
	t287 = t287 + 1
	stack[int(t287)] = t283
	P = P + 3
	indexValidate()
	t288 = stack[int(P)]
	P = P - 3
	// -----------------------------------------------------------

	if t288 == 1 {
		goto L51
	}
	t285 = t285 + 1 //skip len
	t283 = t283 - 1
	t285 = t285 + t283
	t279 = heap[int(t285)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L51:
	fmt.Printf("%d", int(t279))
	// PRIMITIVO:

	t289 = H
	heap[int(H)] = 9 //
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t290 = P + 3
	t290 = t290 + 1
	stack[int(t290)] = t289
	// -----------------------------------------------------------

	P = P + 3
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t291 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 3

	// FIN PRINT
	// -----------------------------------------------------------
	goto L50

L50:
	t293 = P + 2 //Posicionar el puntero
	t292 = stack[int(t293)]
	t292 = t292 + 1
	stack[int(t293)] = t292
	goto L48
L49:
	// FIN FOR
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO:

	t294 = H
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t295 = P + 2
	t295 = t295 + 1
	stack[int(t295)] = t294
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t296 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: VALORES ANTES DE QUICKSORT

	t297 = H
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
	t298 = P + 2
	t298 = t298 + 1
	stack[int(t298)] = t297
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t299 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------
	// -----------------------------------------------------------
	// INICIO FOR
	// PRIMITIVO: 1
	// PRIMITIVO: 2

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARRAY"
	t300 = stack[int(1)]
	t302 = heap[int(t300)]
	// PRIMITIVO: 2

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T302
	// 2
	t303 = P + 2
	t303 = t303 + 1
	stack[int(t303)] = t302
	t303 = t303 + 1
	stack[int(t303)] = 2
	P = P + 2
	indexValidate()
	t304 = stack[int(P)]
	P = P - 2
	// -----------------------------------------------------------

	if t304 == 1 {
		goto L52
	}
	t300 = t300 + 1 //saltar len
	// PRIMITIVO: 2
	t300 = t300 + 1
	t301 = heap[int(t300)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L52:
	t301 = heap[int(t301)]
	t305 = P + 2
	stack[int(t305)] = 1

L53:
	t307 = P + 2 //Posicionar el puntero
	t306 = stack[int(t307)]
	if t306 > t301 {
		goto L54
	}

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: 2
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: X
	t309 = P + 2 //Posicionar el puntero
	t308 = stack[int(t309)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARRAY"
	t310 = stack[int(1)]
	t312 = heap[int(t310)]
	// PRIMITIVO: 2

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T312
	// 2
	t313 = P + 3
	t313 = t313 + 1
	stack[int(t313)] = t312
	t313 = t313 + 1
	stack[int(t313)] = 2
	P = P + 3
	indexValidate()
	t314 = stack[int(P)]
	P = P - 3
	// -----------------------------------------------------------

	if t314 == 1 {
		goto L56
	}
	t310 = t310 + 1 //saltar len
	// PRIMITIVO: 2
	t310 = t310 + 1
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: X
	t316 = P + 2 //Posicionar el puntero
	t315 = stack[int(t316)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t317 = heap[int(t310)]
	t318 = heap[int(t317)]

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T318
	// T315
	t319 = P + 3
	t319 = t319 + 1
	stack[int(t319)] = t318
	t319 = t319 + 1
	stack[int(t319)] = t315
	P = P + 3
	indexValidate()
	t320 = stack[int(P)]
	P = P - 3
	// -----------------------------------------------------------

	if t320 == 1 {
		goto L56
	}
	t317 = t317 + 1 //skip len
	t315 = t315 - 1
	t317 = t317 + t315
	t311 = heap[int(t317)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L56:
	fmt.Printf("%d", int(t311))
	// PRIMITIVO:

	t321 = H
	heap[int(H)] = 9 //
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t322 = P + 3
	t322 = t322 + 1
	stack[int(t322)] = t321
	// -----------------------------------------------------------

	P = P + 3
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t323 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 3

	// FIN PRINT
	// -----------------------------------------------------------
	goto L55

L55:
	t325 = P + 2 //Posicionar el puntero
	t324 = stack[int(t325)]
	t324 = t324 + 1
	stack[int(t325)] = t324
	goto L53
L54:
	// FIN FOR
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO:

	t326 = H
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t327 = P + 2
	t327 = t327 + 1
	stack[int(t327)] = t326
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t328 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: -------------------------

	t329 = H
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
	t330 = P + 2
	t330 = t330 + 1
	stack[int(t330)] = t329
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t331 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// LLAMADO A FUNCION: QUICKSORT
	// GUARDAR TEMPORALES

	t332 = P + 2
	stack[int(t332)] = t216
	t332 = t332 + 1
	stack[int(t332)] = t219
	t332 = t332 + 1
	stack[int(t332)] = t228
	t332 = t332 + 1
	stack[int(t332)] = t243
	t332 = t332 + 1
	stack[int(t332)] = t248
	t332 = t332 + 1
	stack[int(t332)] = t251
	t332 = t332 + 1
	stack[int(t332)] = t267
	t332 = t332 + 1
	stack[int(t332)] = t276
	t332 = t332 + 1
	stack[int(t332)] = t291
	t332 = t332 + 1
	stack[int(t332)] = t296
	t332 = t332 + 1
	stack[int(t332)] = t299
	t332 = t332 + 1
	stack[int(t332)] = t308
	t332 = t332 + 1
	stack[int(t332)] = t323
	t332 = t332 + 1
	stack[int(t332)] = t328
	t332 = t332 + 1
	stack[int(t332)] = t331

	// FIN DE GUARDAR TEMPORALES
	// PRIMITIVO: 2

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARRAY"
	t333 = stack[int(1)]
	t335 = heap[int(t333)]
	// PRIMITIVO: 2

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T335
	// 2
	t336 = P + 17
	t336 = t336 + 1
	stack[int(t336)] = t335
	t336 = t336 + 1
	stack[int(t336)] = 2
	P = P + 17
	indexValidate()
	t337 = stack[int(P)]
	P = P - 17
	// -----------------------------------------------------------

	if t337 == 1 {
		goto L57
	}
	t333 = t333 + 1 //saltar len
	// PRIMITIVO: 2
	t333 = t333 + 1
	t334 = heap[int(t333)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L57:
	// PRIMITIVO: 1
	// PRIMITIVO: 2

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARRAY"
	t338 = stack[int(1)]
	t340 = heap[int(t338)]
	// PRIMITIVO: 2

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T340
	// 2
	t341 = P + 17
	t341 = t341 + 1
	stack[int(t341)] = t340
	t341 = t341 + 1
	stack[int(t341)] = 2
	P = P + 17
	indexValidate()
	t342 = stack[int(P)]
	P = P - 17
	// -----------------------------------------------------------

	if t342 == 1 {
		goto L58
	}
	t338 = t338 + 1 //saltar len
	// PRIMITIVO: 2
	t338 = t338 + 1
	t339 = heap[int(t338)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L58:
	t339 = heap[int(t339)]
	t343 = P + 18
	stack[int(t343)] = t334
	t343 = t343 + 1
	stack[int(t343)] = 1
	t343 = t343 + 1
	stack[int(t343)] = t339
	P = P + 17
	quicksort()
	t343 = stack[int(P)]
	P = P - 17
	// RECUPERAR TEMPORALES

	t344 = P + 2
	t216 = stack[int(t344)]
	t344 = t344 + 1
	t219 = stack[int(t344)]
	t344 = t344 + 1
	t228 = stack[int(t344)]
	t344 = t344 + 1
	t243 = stack[int(t344)]
	t344 = t344 + 1
	t248 = stack[int(t344)]
	t344 = t344 + 1
	t251 = stack[int(t344)]
	t344 = t344 + 1
	t267 = stack[int(t344)]
	t344 = t344 + 1
	t276 = stack[int(t344)]
	t344 = t344 + 1
	t291 = stack[int(t344)]
	t344 = t344 + 1
	t296 = stack[int(t344)]
	t344 = t344 + 1
	t299 = stack[int(t344)]
	t344 = t344 + 1
	t308 = stack[int(t344)]
	t344 = t344 + 1
	t323 = stack[int(t344)]
	t344 = t344 + 1
	t328 = stack[int(t344)]
	t344 = t344 + 1
	t331 = stack[int(t344)]

	// FIN DE RECUPERACION
	// FIN LLAMADO A FUNCION
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: VALORES DESPUES DE QUICKSORT:

	t345 = H
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
	t346 = P + 2
	t346 = t346 + 1
	stack[int(t346)] = t345
	// -----------------------------------------------------------

	P = P + 2
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t347 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 2

	fmt.Printf("%c", int(10))
	// FIN PRINT
	// -----------------------------------------------------------
	// -----------------------------------------------------------
	// INICIO FOR
	// PRIMITIVO: 1
	// PRIMITIVO: 2

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARRAY"
	t348 = stack[int(1)]
	t350 = heap[int(t348)]
	// PRIMITIVO: 2

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T350
	// 2
	t351 = P + 2
	t351 = t351 + 1
	stack[int(t351)] = t350
	t351 = t351 + 1
	stack[int(t351)] = 2
	P = P + 2
	indexValidate()
	t352 = stack[int(P)]
	P = P - 2
	// -----------------------------------------------------------

	if t352 == 1 {
		goto L59
	}
	t348 = t348 + 1 //saltar len
	// PRIMITIVO: 2
	t348 = t348 + 1
	t349 = heap[int(t348)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L59:
	t349 = heap[int(t349)]
	t353 = P + 2
	stack[int(t353)] = 1

L60:
	t355 = P + 2 //Posicionar el puntero
	t354 = stack[int(t355)]
	if t354 > t349 {
		goto L61
	}

	// -----------------------------------------------------------
	// INICIO PRINT
	// PRIMITIVO: 2
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: Y
	t357 = P + 2 //Posicionar el puntero
	t356 = stack[int(t357)]
	// FIN ACCESO
	// -----------------------------------------------------------

	// -----------------------------------------------------------
	// ACCESO ARRAY "ARRAY"
	t358 = stack[int(1)]
	t360 = heap[int(t358)]
	// PRIMITIVO: 2

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T360
	// 2
	t361 = P + 3
	t361 = t361 + 1
	stack[int(t361)] = t360
	t361 = t361 + 1
	stack[int(t361)] = 2
	P = P + 3
	indexValidate()
	t362 = stack[int(P)]
	P = P - 3
	// -----------------------------------------------------------

	if t362 == 1 {
		goto L63
	}
	t358 = t358 + 1 //saltar len
	// PRIMITIVO: 2
	t358 = t358 + 1
	// -----------------------------------------------------------
	// ACCESO A VARIABLE: Y
	t364 = P + 2 //Posicionar el puntero
	t363 = stack[int(t364)]
	// FIN ACCESO
	// -----------------------------------------------------------

	t365 = heap[int(t358)]
	t366 = heap[int(t365)]

	// -----------------------------------------------------------
	// INDEX VALIDATE
	// T366
	// T363
	t367 = P + 3
	t367 = t367 + 1
	stack[int(t367)] = t366
	t367 = t367 + 1
	stack[int(t367)] = t363
	P = P + 3
	indexValidate()
	t368 = stack[int(P)]
	P = P - 3
	// -----------------------------------------------------------

	if t368 == 1 {
		goto L63
	}
	t365 = t365 + 1 //skip len
	t363 = t363 - 1
	t365 = t365 + t363
	t359 = heap[int(t365)]
	// FIN ACCESO ARRAY
	// -----------------------------------------------------------

L63:
	fmt.Printf("%d", int(t359))
	// PRIMITIVO:

	t369 = H
	heap[int(H)] = 9 //
	H = H + 1
	heap[int(H)] = -1 // FIN CADENA
	H = H + 1

	// -----------------------------------------------------------
	// GUARDAR VARIABLE EN STACK
	t370 = P + 3
	t370 = t370 + 1
	stack[int(t370)] = t369
	// -----------------------------------------------------------

	P = P + 3
	printStr()

	// GUARDAR RETURN DE LA FUNCION
	t371 = stack[int(P)]
	// REGRESO DE ENTORNO
	P = P - 3

	// FIN PRINT
	// -----------------------------------------------------------
	goto L62

L62:
	t373 = P + 2 //Posicionar el puntero
	t372 = stack[int(t373)]
	t372 = t372 + 1
	stack[int(t373)] = t372
	goto L60
L61:
	// FIN FOR
	// -----------------------------------------------------------

}
