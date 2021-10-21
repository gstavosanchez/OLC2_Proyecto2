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

    def set_id(self, id: str):
        self.id = id

    def get_id(self):
        return self.id

    def set_type(self, type: TipoVar):
        self.type = type

    def get_type(self):
        return self.type

