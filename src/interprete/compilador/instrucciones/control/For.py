from src.interprete.compilador.tipos.Tipo import TipoVar
from src.interprete.compilador.abstracto.Valor import Valor
from src.interprete.compilador.simbolos.Entorno import Entorno
from src.interprete.compilador.abstracto.Instruccion import Instruccion


class For(Instruccion):
    def __init__(self, id, exp_1, instructions, line, column, exp_2):
        super().__init__(line, column)
        self.id_var = id
        self.exp_1: Instruccion = exp_1
        self.exp_2: Instruccion = exp_2
        self.inst_list: Instruccion = instructions

    def compilar(self, entorno: Entorno):
        value_1: Valor = self.exp_1.compilar(entorno)
        value_2: Valor = self.exp_2.compilar(entorno) if self.exp_2 else None
        self.generador.begin_comment('Inicio For')
        new_env = Entorno(entorno)
        if value_2 is None:
            pass
        else:
            self.for_normal(value_1, value_2, new_env)
        self.generador.end_comment('fin for')

    def set_labels(self):
        pass

    def for_normal(self, value_1: Valor, value_2: Valor, entorno: Entorno):

        if self.is_valid(value_1, value_2) is False:
            self.generador.new_error('Expresion no es de tipo Int64 or Float64')
            return

        new_var = entorno.save_variable(self.id_var, TipoVar.INT64, False)
        # -------------- -> ASIGNACION VARIABLE <- --------------
        tmp_pos = new_var.get_position()

        if not new_var.get_is_global():
            tmp_pos = self.generador.new_temp()
            self.generador.new_exp(tmp_pos, 'P', new_var.get_position(), '+')

        self.generador.set_stack(tmp_pos, value_1.get_value())
        self.generador.line_break()

        # -------------- -> FIN ASIGNACION VARIABLE <- --------------
        # -------------- -> L0...Ln <- --------------
        lb_while = self.generador.new_label()
        true_lb = self.generador.new_label()
        continue_lb = self.generador.new_label()  # continue label

        # -------------- -> WHILE <- --------------
        self.generador.set_label(lb_while)
        # Acceso a variable
        tmp_var = self.generador.new_temp()  # tmp para guardar var
        tmp_pos = new_var.get_position()  # Obtener posicion de la variable
        if not new_var.get_is_global():
            tmp_pos = self.generador.new_temp()
            self.generador.new_exp(
                tmp_pos,
                'P',
                new_var.get_position(),
                '+',
                'Posicionar el puntero',
            )
        self.generador.get_stack(tmp_var, tmp_pos)
        # if (t1 == 10) goto L1 -> Salir
        self.generador.new_if(tmp_var, value_2.get_value(), '>', true_lb)
        entorno.set_break(true_lb)
        entorno.set_continue(lb_while)
        self.inst_list.compilar(entorno)
        self.generador.new_goto(continue_lb)
        self.generador.line_break()
        self.generador.set_label(continue_lb)
        # Acceso a variable

        # Acceso a variable
        tmp_var = self.generador.new_temp()  # tmp para guardar var
        tmp_pos = new_var.get_position()  # Obtener posicion de la variable
        if not new_var.get_is_global():
            tmp_pos = self.generador.new_temp()
            self.generador.new_exp(
                tmp_pos,
                'P',
                new_var.get_position(),
                '+',
                'Posicionar el puntero',
            )
        self.generador.get_stack(tmp_var, tmp_pos)
        self.generador.new_exp(tmp_var, tmp_var, '1', '+')
        self.generador.set_stack(tmp_pos, tmp_var)

        self.generador.new_goto(lb_while)
        self.generador.set_label(true_lb)

    def for_each(self):
        pass

    def is_valid(self, left: Valor, right: Valor):
        if (
            left.get_type() == TipoVar.INT64
            or left.get_type() == TipoVar.FLOAT64
        ):
            if (
                right.get_type() == TipoVar.INT64
                or right.get_type() == TipoVar.FLOAT64
            ):
                return True

        return False
