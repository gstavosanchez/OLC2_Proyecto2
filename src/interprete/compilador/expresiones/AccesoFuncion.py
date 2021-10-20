from src.interprete.compilador.tipos.Tipo import TipoVar, get_tipo_var
from src.interprete.compilador.abstracto.Valor import Valor
from src.interprete.compilador.simbolos.SimboloFuncion import SimboloFuncion
from src.interprete.compilador.abstracto.Instruccion import Instruccion
from src.interprete.compilador.simbolos.Entorno import Entorno


class AccesoFuncion(Instruccion):
    def __init__(self, id, param_values, line, column):
        super().__init__(line, column)
        self.id = id
        self.param_exp: list = param_values
        self.line = line
        self.column = column

    def compilar(self, entorno: Entorno):
        self.generador.line_break()
        self.generador.begin_comment(f'Llamado a funcion: {self.id}')
        funct_simbol: SimboloFuncion = entorno.search_function(self.id)
        if funct_simbol is None:
            self.generador.new_error(
                f'No existe la funcion "{self.id}"', self.line, self.column
            )
            return
        params_values: list = []
        size = self.generador.save_temps(entorno)

        paramas_lenght_funct = funct_simbol.get_size_params()
        send_param_lenght = len(self.param_exp)

        if paramas_lenght_funct != send_param_lenght:
            self.generador.new_error(
                f'Se esperaba {paramas_lenght_funct} argumentos, enviados {send_param_lenght}',
                self.line,
                self.column,
            )
            return

        i = 0
        for exp in self.param_exp:
            param_send: Valor = exp.compilar(entorno)
            param_registered = funct_simbol.get_value_param(i)
            if param_send.get_type() != param_registered.get_type():
                erro_str = (
                    'Argumento de tipo '
                    + get_tipo_var(param_send.get_type())
                    + ' no se puede asiganar al parametro de tipo '
                    + get_tipo_var(param_registered.get_type())
                )
                self.generador.new_error(erro_str, self.line, self.column)
                return

            if param_send.get_type() == TipoVar.BOOLEAN:
                tmp = self.generador.new_temp()
                tmp_lb = self.generador.new_label()
                self.generador.free_temp(tmp)

                self.generador.set_label(param_send.get_true_label())
                self.generador.new_exp(
                    tmp, 'P', entorno.get_size() + i + 1, '+'
                )
                self.generador.set_stack(tmp, '1')
                self.generador.new_goto(tmp_lb)

                self.generador.set_label(param_send.get_false_label())
                self.generador.new_exp(
                    tmp, 'P', entorno.get_size() + i + 1, '+'
                )
                self.generador.set_stack(tmp, '0')
                self.generador.set_label(tmp_lb)

            params_values.append(param_send)
            i += 1

        tmp = self.generador.new_temp()
        self.generador.free_temp(tmp)
        if len(params_values) > 0:
            x = 0
            self.generador.new_exp(tmp, 'P', entorno.get_size() + 1, '+')
            for param in params_values:
                if param.get_type() != TipoVar.BOOLEAN:
                    self.generador.set_stack(tmp, param.get_value())

                if x != len(params_values) - 1:
                    self.generador.new_exp(tmp, tmp, '1', '+')

                x += 1

        self.generador.new_entorno(entorno.get_size())
        self.generador.call_function(funct_simbol.get_id())
        self.generador.get_stack(tmp, 'P')
        self.generador.ret_entorno(entorno.get_size())
        self.generador.recover_tmp(entorno, size)
        self.generador.add_temp(tmp)

        if funct_simbol.get_type() != TipoVar.BOOLEAN:
            self.generador.end_comment('fin llamado a funcion')
            return Valor(tmp, funct_simbol.get_type(), True)

        aux_return = Valor("", funct_simbol.get_type(), False)
        self.set_labels()

        self.generador.new_if(tmp, '1', '==', self.true_label)
        self.generador.new_goto(self.false_label)

        aux_return.set_true_label(self.true_label)
        aux_return.set_false_label(self.false_label)
        self.generador.end_comment('fin llamado a funcion')
        return aux_return

    def set_labels(self):
        if self.true_label == '':
            self.true_label = self.generador.new_label()
        if self.false_label == '':
            self.false_label = self.generador.new_label()
