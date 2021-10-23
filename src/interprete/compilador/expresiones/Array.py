from src.interprete.compilador.tipos.Tipo import TipoVar
from src.interprete.compilador.abstracto.Valor import Valor
from src.interprete.compilador.simbolos.Entorno import Entorno
from src.interprete.compilador.abstracto.Instruccion import Instruccion


class Arreglo(Instruccion):
    def __init__(self, instrucciones, line, column):
        super().__init__(line, column)
        self.inst_list: list = instrucciones
        self.type_list: list = []
        self.values_list = []
        self.entorno: Entorno = None

    def compilar(self, entorno: Entorno):

        self.entorno = entorno

        self.generador.line_break()
        self.generador.begin_comment('ARREGLO')

        tmp_saved, tmp_h = self.generate_temp(len(self.inst_list))
        self.saved_array(self.inst_list, tmp_h)

        self.generador.begin_comment('FIN ARREGLO')
        self.generador.line_break()
        ret_value = Valor(tmp_saved, TipoVar.ARRAY, True, self.type_list)
        ret_value.set_aux_values(self.values_list)
        return ret_value

    # Vector{Vector{Vector{Int64}}}
    # a = [1, 2, [10, 25, 35]]
    def saved_array(self, lista: list, tmp_h_prev):
        for inst in lista:
            value: Valor = inst.compilar(self.entorno)
            self.generador.set_heap(tmp_h_prev, value.get_value())
            self.generador.new_exp(tmp_h_prev, tmp_h_prev, '1', '+')
            # self.type_list.append(value.get_type())
            if value.get_type() == TipoVar.ARRAY:
                # -------------- -> LISTA DE VALORES <- --------------
                lista_values = value.get_aux_values()
                self.values_list.append(lista_values)
                # -------------- -> LISTA DE TIPO <- --------------
                temp_list = value.get_aux_type()
                self.type_list.append(temp_list)
            else:
                self.values_list.append(value)
                self.type_list.append(value.get_type())

    def set_labels(self):
        pass

    def generate_temp(self, length: int):
        """
        Generar temporales

        Args:
            length (int): longitud del arreglo

        Returns:
            tmp: temporal donde se guarda el array
            tmp: contador del heap
        """
        tmp_save = self.generador.new_temp()  # Apuntador donde se guarda el arr
        tmp_index_h = self.generador.new_temp()  # contador del heap

        # -------------- -> EXPRESION <- --------------
        self.generador.new_exp(
            tmp_save, 'H', '', '', 'Apuntador donde se guarda el array'
        )
        self.generador.new_exp(tmp_index_h, 'H', '', '', 'Contador del heap')
        # Almacenar el espacio del arreglo
        self.generador.new_exp(
            'H',
            'H',
            length + 1,
            '+',
            'Almacenar espacio en heap longitud + array',
        )

        self.generador.set_heap(tmp_index_h, length, 'Guadar longtud arreglo')
        self.generador.new_exp(tmp_index_h, tmp_index_h, '1', '+')
        self.generador.new_commnet('Guardar valores')

        return tmp_save, tmp_index_h
