from src.interprete.compilador.abstracto.Instruccion import Instruccion
from src.interprete.compilador.abstracto.Valor import Valor
from src.interprete.compilador.simbolos.Entorno import Entorno
from src.interprete.compilador.tipos.Tipo import TipoArtimetico, TipoVar


class Aritmetica(Instruccion):
    def __init__(self, type, left, right, line, column):
        """

        Args:
            type (TipoAritmetico): Tipo de operacion
            left (Instruccion): operador izquierdo
            right (Instruccion): operador derecho
            line (number): linea
            column (number): columna
        """
        super().__init__(line, column)
        self.type = type
        self.left: Instruccion = left
        self.right: Instruccion = right
        self.line = line
        self.column = column

    def compilar(self, entorno: Entorno):
        self.generador.new_comment_line()
        self.generador.new_commnet('inicio expresion aritmetica')
        left: Valor = self.left.compilar(entorno) if self.left else None
        right: Valor = self.right.compilar(entorno)

        operation = self.get_type(self.type)
        if self.type == TipoArtimetico.POT:
            self.generador.new_comment_line('fin expresion aritmetica')
            self.generador.new_comment_line()
            return self.get_potencia(left, right)

        temp = self.generador.new_temp()  # Nuevo Temporal -> t1...tn

        left_value = left.get_value() if left else '0'
        self.generador.new_exp(temp, left_value, right.get_value(), operation)

        self.generador.new_comment_line('fin expresion aritmetica')
        self.generador.new_comment_line()
        return Valor(temp, TipoVar.FLOAT64, True)

    def get_type(self, operator: TipoArtimetico):
        if operator == TipoArtimetico.SUMA:
            return '+'
        elif (
            operator == TipoArtimetico.RESTA
            or operator == TipoArtimetico.UMENOS
        ):
            return '-'
        elif operator == TipoArtimetico.POR:
            return '*'
        elif operator == TipoArtimetico.DIV:
            return '/'
        elif operator == TipoArtimetico.POT:
            return '^'
        elif operator == TipoArtimetico.MODULO:
            return '%'
        else:
            return ''

    def get_potencia(self, left: Valor, right: Valor):
        # -------------- -> L0...Ln <- --------------
        while_lb = self.generador.new_label()  # L0: -> While Label
        exit_lb = self.generador.new_label()  # L1: -> Exit Label
        # -------------- -> t0...tn <- --------------
        tmp_index = self.generador.new_temp()
        tmp_result = self.generador.new_temp()

        self.generador.new_exp(tmp_index, '0', '', '')  # t0 = 0
        self.generador.new_exp(tmp_result, '1', '', '')  # t1 = 1
        # -------------- -> L1 -> WHILE LABEL <- --------------
        self.generador.set_label(while_lb)  # L1:
        if right.get_type() == TipoVar.FLOAT64:
            print('Error exponente es decimial')
            return

        # if t0 == 2 { goto exit_label }
        self.generador.new_if(tmp_index, right.get_value(), '==', exit_lb)
        # t1 = 5 * t1
        self.generador.new_exp(tmp_result, left.get_value(), tmp_result, '*')
        # t0 = t0 + 1
        self.generador.new_exp(tmp_index, tmp_index, '1', '+')
        # goto while_label
        self.generador.new_goto(while_lb)
        self.generador.set_label(exit_lb)  # L1: -> exit label

        return Valor(tmp_result, TipoVar.FLOAT64, True)

    def set_labels(self):
        pass
