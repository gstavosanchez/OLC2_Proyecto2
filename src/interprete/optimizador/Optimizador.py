class Optimizador:
    def __init__(self, packages, temps, code):
        self.packages = packages
        self.temps = temps
        self.code = code

        self.blocks = []
