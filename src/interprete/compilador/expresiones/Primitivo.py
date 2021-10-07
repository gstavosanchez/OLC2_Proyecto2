from src.interprete.compilador.abstracto.Instruccion import Instruccion
from src.interprete.compilador.abstracto.Valor import Valor
from src.interprete.compilador.simbolos.Entorno import Entorno
from src.interprete.compilador.tipos.Tipo import TipoVar


class Primitivo(Instruccion):
    def __init__(self, valor, type, line, column):
        """
        Args:
            value (any): valor del primitivo
            type (Tipo): tipo del primitvo
            line (number):
            column (number):
        """
        super().__init__(line, column)
        self.value = valor
        self.type = type

    def compilar(self, entorno: Entorno):
        if self.type == TipoVar.INT64 or self.type == TipoVar.FLOAT64:
            return Valor(self.value, self.type, False)
        elif self.type == TipoVar.BOOLEAN:
            self.set_labels()

            if self.value:
                self.generador.new_goto(self.true_label)  # goto L0
                self.generador.new_goto(self.false_label)  # goto L1, evita err
            else:
                self.generador.new_goto(self.false_label)  # goto L1
                self.generador.new_goto(self.true_label)  # goto L0, evita error

            valor_return = Valor(self.value, self.type, False)
            valor_return.set_true_label(self.true_label)
            valor_return.set_false_label(self.false_label)
            return valor_return
        else:
            # TODO:TERMINAR EL STRING
            pass

    def set_labels(self):
        if self.true_label == '':
            self.true_label = self.generador.new_label()  # L0...Ln
        if self.false_label == '':
            self.false_label = self.generador.new_label()  # L1...Ln
