from abc import ABC, abstractmethod

from src.interprete.compilador.simbolos.Entorno import Entorno


class Instruccion(ABC):
    def __init__(self, line, column):
        self.line = line
        self.column = column

    @abstractmethod
    def compilar(self, entorno: Entorno):
        pass
