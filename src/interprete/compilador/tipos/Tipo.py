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
    VOID = 9            # Tipo vacio


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

class TipoUpLowCase(Enum):
    UPPER = 1           # Upper Case
    LOWER = 2           # Lower Case

class TipoScoope(Enum):
    GLOBAL = 1          # global
    LOCAL = 2           # local
    UNDEFINED = 3       # NO definido

    
def get_type_print(type_print: TipoPrint):
    if type_print == TipoPrint.CARACTER:
        return 'c', 'int'
    return 'g', None



def get_tipo_var(type: TipoVar):
    if type == TipoVar.INT64:
        return 'Int64'
    elif type == TipoVar.FLOAT64:
        return 'Float64'
    elif type == TipoVar.BOOLEAN:
        return 'Bool'
    elif type == TipoVar.STRING:
        return 'String'
    elif type == TipoVar.CHAR:
        return 'Char'
    elif type == TipoVar.ARRAY:
        return 'Array'
    elif type == TipoVar.STRUCT:
        return 'Struct'
    elif type == TipoVar.ANY:
        return 'Any'
    elif type == TipoVar.VOID:
        return 'Void'
    else:
        return 'Nothing'