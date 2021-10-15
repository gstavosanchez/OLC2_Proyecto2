from src.interprete.compilador.simbolos.Simbolo import Simbolo
from src.interprete.compilador.tipos.Tipo import TipoVar


class Entorno:
    def __init__(self, previous_enviroment=None):
        self.previous: Entorno = previous_enviroment
        self.size = 0
        self.variables: dict = {}
        self.functions: dict = {}
        self.structs: dict = {}
        self.set_size()
        self.break_lb = None
        self.continue_lb = None

    def set_size(self):
        self.size = self.previous.size if self.previous else 0

    def get_size(self):
        return self.size

    # --------------------------------------------------------------------------
    # VARIABLES
    # --------------------------------------------------------------------------

    def save_variable(self, id: str, type: TipoVar, in_heap: bool):
        """Guardar Variable

        Args:
            id (str): identificador del la variable
            type (TipoVar): tipo de variable
            in_heap (bool): esta en el heap

        Returns:
            Simbolo | None: Nuevo simbolo agregado
        """
        # FIXME:
        # Si existe no dejar declarar  o sobreescribirla

        if id in self.variables.keys():
            var: Simbolo = self.variables[id]
            var.set_type(type)
            var.set_in_heap(in_heap)
            var.set_is_global(self.previous == None)
            return var

        new_var = Simbolo(id, type, self.size, self.previous == None, in_heap)
        self.size += 1
        self.variables[id] = new_var
        return new_var

    def get_variable(self, id: str):
        entorno = self
        while entorno != None:
            if id in entorno.variables.keys():
                return entorno.variables[id]
            entorno = entorno.previous
        return None

    # --------------------------------------------------------------------------
    # GLOBAL
    # --------------------------------------------------------------------------
    def get_global(self):
        entorno = self
        while entorno.previous != None:
            entorno = entorno.previous
        return entorno

    # --------------------------------------------------------------------------
    # TRANFERENCIA
    # --------------------------------------------------------------------------
    def set_break(self, break_lb: str):
        self.break_lb = break_lb

    def set_continue(self, continue_lb: str):
        self.continue_lb = continue_lb

    def get_break(self):
        return self.break_lb

    def get_continue(self):
        return self.continue_lb
