from src.interprete.compilador.abstracto.Valor import Valor
from src.interprete.compilador.simbolos.Parametro import Parametro
from src.interprete.compilador.simbolos.SimboloStruct import SimboloStruct
from src.interprete.compilador.tipos.Tipo import TipoVar
from src.interprete.compilador.simbolos.Simbolo import Simbolo
from src.interprete.compilador.simbolos.Entorno import Entorno
from src.interprete.compilador.abstracto.Instruccion import Instruccion


class AccesoStruct(Instruccion):
    def __init__(self, id, att_name, line, column):
        super().__init__(line, column)
        self.id = id
        self.att_name = att_name
        self.line = line
        self.column = column

    def compilar(self, entorno: Entorno):
        variable: Simbolo = entorno.get_variable(self.id)
        if variable is None:
            self.generador.new_error(
                f'No existe la variable "{self.id}"', self.line, self.column
            )
            return

        if (
            variable.get_type() != TipoVar.STRUCT
            or variable.get_tyep_struct() is None
        ):
            self.generador.new_error(
                f'La variable "{self.id}" no es de tipo Struct',
                self.line,
                self.column,
            )
            return

        # -------------- -> T0...Tn <- --------------
        tmp_i = self.generador.new_temp()
        tmp_saved = self.generador.new_temp()
        tmp_pos = variable.get_position()
        # -------------- -> Aux <- --------------
        i = 0
        att_type = None
        struct: SimboloStruct = entorno.get_struct(variable.get_tyep_struct())
        if not variable.get_is_global():
            tmp_pos = self.generador.new_temp()
            self.generador.new_exp(tmp_pos, 'P', variable.get_position(), '+')

        self.generador.get_stack(tmp_i, tmp_pos)

        for att in struct.get_att_list():
            att: Parametro
            if att.get_id() == self.att_name:
                att_type = att.get_type()
                break
            i += 1

        if att_type is None:
            self.generador.new_error(
                f'El atributo "{self.att_name}" no es existe',
                self.line,
                self.column,
            )
            return

        self.generador.new_exp(tmp_i, tmp_i, i, '+')
        self.generador.get_heap(tmp_saved, tmp_i)

        return Valor(tmp_saved, att_type, True)

    def set_labels(self):
        pass
