from src.interprete.compilador.tipos.Tipo import TipoPrint, get_type_print


class Generador:
    __instance = None

    @staticmethod
    def get_instance():
        """
        STATIC ACCESS METHOD

        Forma de correcta de utilizar:
            generator = Generador.get_instance()
        """
        if Generador.__instance is None:
            Generador.__instance = Generador()
        return Generador.__instance

    def __init__(self):
        # contadores
        self.index_temp = 0
        self.index_label = 0
        # codigo
        self.code = ''
        self.func = ''
        self.natives = ''
        self.in_func = False
        self.in_native = False
        # lista temporales y nativas
        self.temp_list: list = []
        # print
        self.is_print = False

    def clean(self):
        # contadores
        self.index_temp = 0
        self.index_label = 0
        # codigo
        self.code = ''
        self.func = ''
        self.natives = ''
        self.in_func = False
        self.in_native = False
        # lista temporales y nativas
        self.temp_list: list = []
        # print
        self.is_print = False

    # --------------------------------------------------------------------------
    # NUEVO TEMPORAL, LABEL, GOTO && IF
    # --------------------------------------------------------------------------
    # -------------- -> TEMPORAL <- --------------
    def new_temp(self):
        """Crea nueva variable temporal

        Returns:
            str: Nuevo temporal de tipo str
        """
        temp = f't{self.index_temp}'
        self.index_temp += 1
        self.temp_list.append(temp)
        return temp

    # -------------- -> LABEL O ETIQUTA <- --------------
    def new_label(self):
        """Crea un nuevo label o etiqueta

        Returns:
            str: Debuevelve el label o etiqueta de tipo str
        """
        label = f'L{self.index_label}'
        self.index_label += 1
        return label

    def set_label(self, label: str):
        """Agregae el label o etiqueta al codigo

        Args:
            label (str): etiqueta
        """
        self.set_code(f'{label}:\n')

    # -------------- -> GOTO <- --------------
    def new_goto(self, label: str):
        """Agrega el label o etiqueta a un goto y setea al codigo

        Args:
            label (str): etiqueta
        """
        self.set_code(f'goto {label};\n')

    # -------------- -> IF <- --------------
    def new_if(self, left: str, right: str, op: str, label: str):
        """Agrega un nuevo if con un goto y setea al codigo

        Args:
            left (str): operador izquierdo
            right (str): operador derecho
            op (str): operacion a realizar
            label (str): etiqueta a saltar
        """
        if_str = f'if {left} {op} {right} '
        if_str += '{ goto ' + str(label) + '; }\n'
        self.set_code(if_str)

    # --------------------------------------------------------------------------
    # CODE MAIN GOLANG
    # --------------------------------------------------------------------------
    def set_code(self, code: str):
        self.code += '\t' + code

    def get_code(self):
        text = self.get_header()
        text += 'func main() {\n'
        text += self.code + '\n'
        text += '}\n'
        return str(text)

    # -------------- -> HEADER <- --------------
    def get_header(self):
        text = 'package main\n'
        text += 'import "fmt"\n'
        # Declaracion de variables
        if len(self.temp_list) > 0:
            text += 'var '
            for i in range(len(self.temp_list)):
                text += self.temp_list[i]
                if i != len(self.temp_list) - 1:
                    text += ', '
            text += ' float64;\n'
        # TODO: Agregar el heap y el stack
        return text

    # -------------- -> FUNCTION <- --------------

    def new_commnet(self, text: str):
        comment = (
            '// -----------------------------------------------------------\n'
        )
        comment += f'\t// {text}\n'
        comment += (
            '\t// -----------------------------------------------------------\n'
        )
        self.set_code(comment)

    def line_break(self):
        """Salto de linea"""
        self.set_code('\n')

    # --------------------------------------------------------------------------
    # EXPRESIONES
    # --------------------------------------------------------------------------
    def new_exp(self, result: str, left: str, right: str, op: str):
        self.set_code(f'{result} = {left} {op} {right}; \n')

    # --------------------------------------------------------------------------
    # INSTRUCCIONES: PRINT
    # --------------------------------------------------------------------------
    # fmt.Printf("%c", 116)     // Imprimir un caracter
    # fmt.Printf("%d", 116)     // Imprimir un numero
    # fmt.Printf("%g", 2.05)    // Imprimir un float
    # fmt.Printf("%c", int(10)) // Imprimir un salto de linea

    def print_new(self, type: TipoPrint, value):
        """Agregar un Print

        Args:
            type (TipoPrint): tipo del print
            value (any): valor a imprimir
        """
        tipo, casteo = get_type_print(type)
        data = f'fmt.Printf("%{tipo}", {casteo}({value}));\n'
        self.set_code(data)

    def print_line_break(self):
        """Salto de Linea"""
        self.print_new(TipoPrint.CARACTER, 10)

    def print_space(self):
        self.print_new(TipoPrint.CARACTER, 32)

    def print_true(self):
        self.print_new(TipoPrint.CARACTER, 116)  # T
        self.print_new(TipoPrint.CARACTER, 114)  # R
        self.print_new(TipoPrint.CARACTER, 117)  # U
        self.print_new(TipoPrint.CARACTER, 101)  # E

    def print_false(self):
        self.print_new(TipoPrint.CARACTER, 102)  # F
        self.print_new(TipoPrint.CARACTER, 97)  # A
        self.print_new(TipoPrint.CARACTER, 108)  # L
        self.print_new(TipoPrint.CARACTER, 115)  # S
        self.print_new(TipoPrint.CARACTER, 101)  # E
