from src.interprete.compilador.simbolos.Simbolo import Simbolo
from src.interprete.compilador.abstracto.Valor import Valor
from src.interprete.compilador.tipos.Tipo import TipoScoope, TipoVar
from src.interprete.compilador.simbolos.Entorno import Entorno
from src.interprete.compilador.abstracto.Instruccion import Instruccion


class Asignacion(Instruccion):
    def __init__(
        self, id, type, value, line, column, scoope=TipoScoope.UNDEFINED
    ):
        super().__init__(line, column)
        self.id = id
        self.value: Instruccion = value
        self.type = type
        self.line = line
        self.column = column
        self.scoope: TipoScoope = scoope

    def compilar(self, entorno: Entorno):
        # Compilar el valor
        value_compiled: Valor = self.value.compilar(entorno)

        if value_compiled is None:
            return

        # Validar tipo
        if self.type:
            if self.type != value_compiled.get_type():
                self.generador.new_error(
                    f'"{self.id}" no se le puede asignar ese tipo',
                    self.line,
                    self.column,
                )
                return
        type_aux = self.type if self.type else value_compiled.get_type()

        new_var = entorno.get_variable(self.id)

        if new_var is None:
            if self.scoope == TipoScoope.GLOBAL:
                env_global = entorno.get_global()
                new_var = env_global.save_variable(
                    self.id,
                    type_aux,
                    self.in_heap(value_compiled.get_type()),
                )
            else:
                new_var = entorno.save_variable(
                    self.id,
                    type_aux,
                    self.in_heap(value_compiled.get_type()),
                )
        else:
            new_var.set_type(type_aux)
            new_var.set_in_heap(self.in_heap(value_compiled.get_type()))

        self.save_arrays(new_var, value_compiled)
        # if new_var is None:
        #     self.generador.new_error(
        #         f'"{self.id}" no puede ser declarada en este scoop',
        #         self.line,
        #         self.column,
        #     )
        #     return
        self.generador.new_comment_line()
        self.generador.new_commnet(f'Asignacion: {self.id}')
        tmp_pos = new_var.get_position()
        if not new_var.get_is_global():
            tmp_pos = self.generador.new_temp()
            self.generador.free_temp(tmp_pos)  # FIXME: REVISAR
            self.generador.new_exp(tmp_pos, 'P', new_var.get_position(), '+')

        if value_compiled.get_type() == TipoVar.BOOLEAN:
            lb_end = self.generador.new_label()

            # -------------- -> TRUE LABEL <- --------------
            self.generador.set_label(value_compiled.get_true_label())
            self.generador.set_stack(tmp_pos, '1')

            self.generador.new_goto(lb_end)

            # -------------- -> FALSE LABEL <- --------------
            self.generador.set_label(value_compiled.get_false_label())
            self.generador.set_stack(tmp_pos, '0')

            self.generador.set_label(lb_end)
        else:
            self.generador.set_stack(tmp_pos, value_compiled.get_value())

        self.generador.new_commnet(f'Fin asignacion de {self.id}')
        self.generador.new_comment_line()
        self.generador.line_break()

    def set_labels(self):
        pass

    def in_heap(self, tipo: TipoVar):
        if (
            tipo == TipoVar.STRUCT
            or tipo == TipoVar.ARRAY
            or tipo == TipoVar.ARRAY
        ):
            return True
        return False

    def save_arrays(self, new_var: Simbolo, value: Valor):
        if (
            new_var.get_type() == TipoVar.ARRAY
            and value.get_type() == TipoVar.ARRAY
        ):
            new_var.set_list_aux_types(value.get_aux_type())
            new_var.set_list_aux_values(value.get_aux_values())
