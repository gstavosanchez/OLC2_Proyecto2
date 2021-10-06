from enum import Enum


class TipoVar(Enum):
    NOTHING = 0         # Null (None)
    INT64 = 1           # Int (Integer)
    FLOAT64 = 2         # Float (Decimal)
    BOOLEAN = 3         # Boolean (True || False)
    CHAR = 4            # Char ('d')
    STRING = 5          # String ("Cadena")
    ARRAY = 6           # Array ([])
    STRUCT = 7          # Struct
    ANY = 8             # Cualquier Valor


class TipoArtimetico(Enum):
    SUMA = 1            # Suma (+)
    RESTA = 2           # Resta (-)
    POR = 3             # Multiplicacion (*)
    DIV = 4             # Division (/)
    POT = 5             # Potencia (^)
    UMENOS = 6          # NEGATIVO (-)
    MODULO = 7          # Modulo (%)

class TipoRelational(Enum):
    MAYORQ = 1           # Mayor (>)
    MENORQ = 2           # Menor (<)
    MAYORI = 3           # Mayor Igual (>=)
    MENORI = 4           # Menor Igual (<=)
    COMPARACION = 5      # Comparacion (==)
    DIFERENTE = 6        # Diferente (!=)

class TipoLogico(Enum):
    OR = 1               # Or (||)
    AND = 2              # And (&&)
    NOT = 3              # Not (!)


class TipoPrint(Enum):
    CARACTER = 1        # "%c"
    FLOAT = 2           # "%g"

def get_type_print(type_print: TipoPrint):
    if type_print == TipoPrint.CARACTER:
        return 'c', 'int'
    return 'g', 'float64'