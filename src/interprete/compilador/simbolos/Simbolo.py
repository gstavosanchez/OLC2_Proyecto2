from src.interprete.compilador.tipos.Tipo import TipoVar


class Simbolo:
    def __init__(self, id: str, type: TipoVar, position, global_var, in_heap):
        """Constructor

        Args:
            id (str): identificadro de la variable
            tipo (TipoVar): tipo de variable
            posicion (any):
            global_var (any):
            in_heap (any):
        """
        self.id = id
        self.type = type
        self.globa_var = global_var
        self.in_heap = in_heap
        self.position = position

    # ==========================================================================
    # SET AND GET: ID
    # ==========================================================================
    def set_id(self, id):
        self.id = id

    def get_id(self):
        return self.id

    # ==========================================================================
    # SET AND GET: TYPE
    # ==========================================================================
    def set_type(self, type: TipoVar):
        self.type = type

    def get_type(self):
        return self.type

    # ==========================================================================
    # SET AND GET: GLOBA_VAR
    # ==========================================================================
    def set_global_var(self, global_var):
        self.globa_var = global_var

    def get_global_var(self):
        return self.globa_var

    # ==========================================================================
    # SET AND GET: IN_HEAP
    # ==========================================================================
    def set_in_heap(self, in_heap):
        self.in_heap = in_heap

    def get_in_heap(self):
        return self.in_heap

    # ==========================================================================
    # SET AND GET: POSITION
    # ==========================================================================
    def set_position(self, position):
        self.position = position

    def get_position(self):
        return self.position
