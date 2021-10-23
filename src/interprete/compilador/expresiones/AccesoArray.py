from src.interprete.compilador.abstracto.Valor import Valor
from src.interprete.compilador.tipos.Tipo import TipoVar
from src.interprete.compilador.simbolos.Simbolo import Simbolo
from src.interprete.compilador.simbolos.Entorno import Entorno
from src.interprete.compilador.abstracto.Instruccion import Instruccion


class AccesoArray(Instruccion):
    def __init__(self, id, list_position, line, column):
        super().__init__(line, column)
        self.id = id
        self.position_list: list = list_position
        self.line = line
        self.column = column

    def compilar(self, entorno: Entorno):
        variable: Simbolo = entorno.get_variable(self.id)
        if variable is None:
            self.generador.new_error(
                f'No existe la variable "{self.id}"', self.line, self.column
            )
            return

        if variable.get_type() != TipoVar.ARRAY:
            self.generador.new_error(
                f'La variable "{self.id}" no es de tipo Array',
                self.line,
                self.column,
            )
            return

        tmp_recov = self.generador.new_temp()
        tmp_saved = self.generador.new_temp()
        self.generador.get_stack(tmp_recov, variable.get_position())
        self.generador.new_exp(tmp_recov, tmp_recov, '1', '+', 'saltar len')
        type_aux = TipoVar.ANY
        if len(self.position_list) == 1:
            # verficar la longitud del arreglo para el error
            index_val: Valor = self.position_list[0].compilar(entorno)
            # verficar el tipo del index_val no puede ser float
            if index_val.get_value() > 0:
                index = index_val.get_value() - 1
                type_aux = variable.get_list_aux_types()[index]
                self.generador.new_exp(tmp_recov, tmp_recov, index, '+')
        else:
            pass
        self.generador.get_heap(tmp_saved, tmp_recov)
        return Valor(tmp_saved, type_aux, True)

    def set_labels(self):
        pass
