from src.interprete.compilador.abstracto.Instruccion import Instruccion
from src.interprete.compilador.abstracto.Valor import Valor
from src.interprete.compilador.simbolos.Entorno import Entorno
from src.interprete.compilador.simbolos.Generador import Generador
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
        generador = Generador.get_instance()
        left: Valor = self.left.compilar(entorno) if self.left else None
        right: Valor = self.right.compilar(entorno)

        operation = self.get_type(self.type)
        temp = generador.new_temp()  # Nuevo Temporal -> t1...tn

        left_value = left.get_value() if left else '0'
        generador.new_exp(temp, left_value, right.get_value(), operation)

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
