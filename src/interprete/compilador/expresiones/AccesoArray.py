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
        if len(self.position_list) > self.verify_dimension(
            variable.get_list_aux_types()
        ):
            self.generador.new_error(
                f'La dimension del arreglo "{self.id}" no son las mismas',
                self.line,
                self.column,
            )
            return

        if not self.verify_index(entorno):
            self.generador.new_error(
                'El indice debe ser tipo Int64 y mayor que 0',
                self.line,
                self.column,
            )
            return
        self.generador.line_break()
        self.generador.begin_comment(f'acceso array "{self.id}"')

        # -------------- -> L0..Ln <- --------------
        exit_lb = self.generador.new_label()
        # -------------- -> T0..Tn <- --------------
        tmp_recov = self.generador.new_temp()
        tmp_saved = self.generador.new_temp()
        tmp_value_h = self.generador.new_temp()
        # FIXME:
        # REVISAR
        tmp_pos = variable.get_position()
        if not variable.get_is_global():
            tmp_pos = self.generador.new_temp()
            self.generador.new_exp(tmp_pos, 'P', variable.get_position(), '+')

        self.generador.get_stack(tmp_recov, tmp_pos)
        # -------------- -> VALIDACION <- --------------
        self.generador.get_heap(tmp_value_h, tmp_recov)
        index_val: Valor = self.position_list[0].compilar(entorno)
        validate: Valor = self.validate_index(
            tmp_value_h, index_val.get_value(), entorno
        )
        self.generador.new_if(validate.get_value(), '1', '==', exit_lb)
        # -------------- -> FIN VALIDACION <- --------------
        self.generador.new_exp(tmp_recov, tmp_recov, '1', '+', 'saltar len')
        type_aux = variable.get_tipo_aux()
        if len(self.position_list) == 1:
            # verficar la longitud del arreglo para el error
            index_val: Valor = self.position_list[0].compilar(entorno)
            index = index_val.get_value() - 1

            # type_aux = variable.get_list_aux_types()[index]
            self.generador.new_exp(tmp_recov, tmp_recov, index, '+')
            self.generador.get_heap(tmp_saved, tmp_recov)

            # self.generador.end_comment(f'fin acceso array')
            # return Valor(tmp_saved, type_aux, True)
        else:
            # a = [1, 2, 3, 4, 5, [10, 25, [50, 20, 200]]]
            # a[6][3][1]
            # t1 = t1 + 5 = pos 6

            # t2 = heap[t1] = puntero al nuevo array
            # t2 = t2 + 1 = skip len
            # t2 = t2 + 2 = pos 3

            # t3 = heap[t2] = puntero al nuevo array
            # t3 = t3 + 1 = skip len
            # t3 = t3 + 0 = pos = 1

            # t4 = heap[t3]
            for x in range(len(self.position_list)):
                index_val: Valor = self.position_list[x].compilar(entorno)
                index = index_val.get_value()
                # -------------- -> TEMPRALES <- --------------
                tmp_aux = tmp_recov  # t1
                if x == 0:
                    self.generador.new_exp(tmp_aux, tmp_aux, index - 1, '+')
                    continue
                tmp_recov = self.generador.new_temp()  # t2

                self.generador.get_heap(tmp_recov, tmp_aux)
                # -------------- -> VALIDACION <- --------------
                validate: Valor = self.validate_index(tmp_recov, index, entorno)
                self.generador.new_if(validate.get_value(), '1', '==', exit_lb)

                # -------------- -> FIN VALIDACION <- --------------
                self.generador.new_exp(
                    tmp_recov, tmp_recov, '1', '+', 'skip len'
                )
                self.generador.new_exp(tmp_recov, tmp_recov, index - 1, '+')

            self.generador.get_heap(tmp_saved, tmp_recov)
        self.generador.end_comment(f'fin acceso array')
        self.generador.set_label(exit_lb)
        return Valor(tmp_saved, type_aux, True)

    def set_labels(self):
        pass

    def get_value(self):
        pass

    def verify_index(self, entorno: Entorno):
        for i in self.position_list:
            index: Valor = i.compilar(entorno)
            if index.get_type() != TipoVar.INT64:
                return False
            if index.get_value() < 1:
                return False

        return True

    def verify_dimension(self, lista: list):
        lenght = 1
        for i in lista:
            if isinstance(i, list):
                lenght += 1
        return lenght

    def validate_index(self, leght: str, index: int, entorno: Entorno):
        self.generador.validate_index_array()
        self.generador.line_break()
        self.generador.new_comment_line()
        # self.generador.new_commnet('Paso de parametros')
        # Temporal para almacenar parametros
        tmp_p = self.generador.new_temp()
        self.generador.new_exp(tmp_p, 'P', entorno.get_size(), '+')
        # Guardar primer parametro
        # self.generador.new_commnet('1er Parametro')
        self.generador.new_exp(tmp_p, tmp_p, '1', '+')
        self.generador.set_stack(tmp_p, leght)
        # Guardar segundo parametro
        # self.generador.new_commnet('2do Parametro')
        self.generador.new_exp(tmp_p, tmp_p, '1', '+')
        self.generador.set_stack(tmp_p, index)
        # self.generador.new_commnet('fin paso de parametros')
        # self.generador.new_comment_line()
        # Cambio de parametro para buscar los parametros
        # self.generador.new_commnet('cambio de entorno')
        self.generador.new_entorno(entorno.get_size())
        self.generador.call_function('indexValidate')
        # Gurdar el return de la funcion
        # self.generador.new_commnet('Guardar el return de la funcion')
        return_p = self.generador.new_temp()
        self.generador.get_stack(return_p, 'P')
        self.generador.ret_entorno(entorno.get_size())

        # self.generador.new_commnet('fin expresion aritmetica')
        self.generador.new_comment_line()
        self.generador.line_break()

        return Valor(return_p, TipoVar.INT64, True)
