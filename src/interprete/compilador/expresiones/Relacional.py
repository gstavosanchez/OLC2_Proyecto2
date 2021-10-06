from src.interprete.compilador.abstracto.Instruccion import Instruccion
from src.interprete.compilador.simbolos.Entorno import Entorno


class Relacional(Instruccion):
    def __init__(self, type, left, right, line, column):
        """

        Args:
            type (TipoRelacional): Tipo de operacion
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
        pass
