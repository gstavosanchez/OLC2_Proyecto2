from src.interprete.compilador.abstracto.Instruccion import Instruccion
from src.interprete.compilador.abstracto.Valor import Valor
from src.interprete.compilador.simbolos.Entorno import Entorno
from src.interprete.compilador.simbolos.Generador import Generador
from src.interprete.compilador.tipos.Tipo import TipoVar


class Primitivo(Instruccion):
    def __init__(self, value, type, line, column):
        """
        Args:
            value (any): valor del primitivo
            type (Tipo): tipo del primitvo
            line (number):
            column (number):
        """
        Instruccion.__init__(self, line, column)
        self.value = value
        self.type = type

    def compilar(self, entorno: Entorno):
        generador = Generador.get_instance()
        if self.type == TipoVar.INT64 or self.type == TipoVar.FLOAT64:
            return Valor(str(self.value), self.type, False)
        elif self.type == TipoVar.BOOLEAN:
            true_label = generador.new_label()  # L0...Ln
            false_label = generador.new_label()  # L1...Ln

            if self.value:
                generador.new_goto(true_label)  # goto L0
                generador.new_goto(false_label)  # goto L1, evita error en go
            else:
                generador.new_goto(false_label)  # goto L1
                generador.new_goto(true_label)  # goto L0, evita error en go

            value = Valor(self.value, self.type, False)
            value.set_true_label(true_label)
            value.set_false_label(false_label)
            return value
        else:
            # TODO:TERMINAR EL STRING
            pass
