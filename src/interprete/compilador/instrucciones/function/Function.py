from src.interprete.compilador.tipos.Tipo import TipoVar
from src.interprete.compilador.simbolos.SimboloFuncion import SimboloFuncion
from src.interprete.compilador.simbolos.Entorno import Entorno
from src.interprete.compilador.abstracto.Instruccion import Instruccion


class Funcion(Instruccion):
    def __init__(self, id, params_list, instruccions, line, column, tipo):
        super().__init__(line, column)
        self.id = id
        self.param_list: list = params_list
        self.instruccions: Instruccion = instruccions
        self.tipo = self.set_tipo(tipo)
        self.line = line
        self.column = column
        self.pre_compie = True

    def compilar(self, entorno: Entorno):
        if self.pre_compie:
            self.pre_compie = False
            funct_simbol: SimboloFuncion = entorno.save_function(
                self.id, self.param_list, self.instruccions, self.tipo
            )
            if funct_simbol is None:
                self.generador.new_error(
                    f'La funcion "{self.id}" ya existe', self.line, self.column
                )
                return

        new_env = Entorno(entorno)
        funct_simbol = entorno.get_function(self.id)
        return_lb = self.generador.new_label()

        tmp_list_copy = self.generador.get_temp_list()

        new_env.set_enviroment_funct(funct_simbol, return_lb)

        for param in self.param_list:
            new_env.save_variable(
                param.get_id(),
                param.get_type(),
                (
                    param.get_type() == TipoVar.STRING
                    or param.get_type() == TipoVar.STRUCT
                ),
            )
        self.generador.clear_temp_list()

        self.generador.new_begin_func(self.id)
        self.instruccions.compilar(new_env)
        self.generador.set_label(return_lb)
        self.generador.new_end_funct()

        # PRUEBA
        # tmp_list_copy2 = self.generador.get_temp_list()

        self.generador.set_temp_list(tmp_list_copy)
        # self.generador.set_temp_list(
        #     self.set_list_copy(tmp_list_copy, tmp_list_copy2)
        # )

    def set_labels(self):
        pass

    def set_tipo(self, tipo: TipoVar):
        # if tipo == TipoVar.INT64:
        #     tipo = TipoVar.FLOAT64
        return tipo

    def set_list_copy(self, copy1: list, copy2: list):
        for tmp in copy2:
            if tmp not in copy1:
                copy1.append(tmp)

        return copy1
