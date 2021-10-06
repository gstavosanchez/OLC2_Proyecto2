from src.interprete.compilador.abstracto.Instruccion import Instruccion
from src.interprete.compilador.abstracto.Valor import Valor
from src.interprete.compilador.simbolos.Entorno import Entorno
from src.interprete.compilador.simbolos.Generador import Generador
from src.interprete.compilador.tipos.Tipo import TipoPrint, TipoVar


class Print(Instruccion):
    def __init__(self, expresion_list, line, column, new_line: bool = False):
        super().__init__(line, column)
        self.exp_list = expresion_list
        self.new_line = new_line
        self.line = line
        self.column = column
        self.generador = Generador.get_instance()

    def compilar(self, entorno: Entorno):
        for exp in self.exp_list:
            exp: Instruccion
            if exp is not None:
                value: Valor = exp.compilar(entorno)
                if value is None:
                    return

                # --------------------------------------------------------------
                # INT || FLOAT
                # --------------------------------------------------------------
                if (
                    value.get_type() == TipoVar.INT64
                    or value.get_type() == TipoVar.FLOAT64
                ):
                    self.generador.print_new(TipoPrint.FLOAT, value.get_value())
                # --------------------------------------------------------------
                # BOOLEAN
                # --------------------------------------------------------------
                elif value.get_type() == TipoVar.BOOLEAN:
                    exit_label = self.generador.new_label()  # etiqueta salida
                    # -------------- -> L0...Ln: <- --------------
                    # L0:
                    #   PRINT TRUE
                    self.generador.set_label(value.get_true_label())  # L0:
                    self.generador.print_true()  # TRUE
                    # -------------- -> GOTO SALIDA <- --------------
                    # GOTO de salida o escape
                    # goto L2
                    self.generador.new_goto(exit_label)  # GOTO SALIDA
                    # -------------- -> L1...Ln: <- --------------
                    # L1:
                    #   PRINT FALSE
                    self.generador.set_label(value.get_false_label())  # L1:
                    self.generador.print_false()  # FALSE
                    # -------------- -> L2...Ln: <- --------------
                    # ETIQUETA SALIDA
                    # L2:
                    #  '\n'
                    self.generador.set_label(exit_label)  # L2:
                    # self.generador.print_line_break()
                    self.generador.print_space()

                else:
                    pass

        if self.new_line:
            self.generador.print_line_break()
