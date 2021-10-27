from src.interprete.compilador.tipos.Tipo import TipoVar


class Parametro:
    def __init__(self, id: str, type: TipoVar):
        """Parametro

        Args:
            id (str): identificador del parametro
            type (TipoVar): tipo de parametro
        """
        self.id = id
        self.type = type
        self.tyep_aux = None
        self.verify_type(type)

    def set_id(self, id: str):
        self.id = id

    def get_id(self):
        return self.id

    def set_type(self, type: TipoVar):
        self.type = type

    def get_type(self):
        return self.type

    def verify_type(self, type):
        if isinstance(type, TipoVar):
            self.type = type
        else:
            self.tyep_aux = type
            self.type = TipoVar.STRUCT

    def get_type_aux(self):
        return self.tyep_aux
