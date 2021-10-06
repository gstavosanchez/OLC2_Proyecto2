class Entorno:
    def __init__(self, previous_enviroment=None):
        self.previous: Entorno = previous_enviroment
        self.size = 0
        self.variables: dict = {}
        self.functions: dict = {}
        self.structs: dict = {}
        self.set_size()

    def set_size(self):
        self.size = self.previous.size if self.previous else 0
