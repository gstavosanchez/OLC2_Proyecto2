from src.interprete.compilador.expresiones.Aritmetica import Aritmetica
from src.interprete.compilador.expresiones.Logica import Logica
from src.interprete.compilador.expresiones.Primitivo import Primitivo
from src.interprete.compilador.expresiones.Relacional import Relacional
from src.interprete.compilador.instrucciones.Print import Print
from src.interprete.compilador.tipos.Tipo import (
    TipoArtimetico,
    TipoLogico,
    TipoRelational,
    TipoVar,
)

# ==============================================================================
# PALABRAS RESERVADAS
# ==============================================================================
reservadas = {
    # INSTRUCCIONES
    'print': 'RPRINT',
    'println': 'RPRINTLN',
    # RESERVADAS
    'true': 'RTRUE',
    'false': 'RFALSE',
}
# ==============================================================================
# TOKENS
# ==============================================================================
tokens = [
    # SIMBOLOS
    'PCOMA',
    'PARA',
    'PARC',
    'COMA',
    # 'IGUAL',
    # ARITMETICOS
    'MAS',
    'MENOS',
    'POR',
    'DIV',
    'POT',
    'MODULO',
    # RELACIONAL
    'MAYORQ',
    'MENORQ',
    'MAYORI',
    'MENORI',
    'COMPARACION',
    'DIFERENTE',
    # LOGICO
    'OR',
    'AND',
    'NOT',
    # TIPO
    'DECIMAL',
    'ENTERO',
    'CADENA',
    'ID',
] + list(reservadas.values())
# -------------- -> TOKEN SIMBOLOS <- --------------
t_PARA = r'\('
t_PARC = r'\)'
t_COMA = r','
t_PCOMA = r';'
# t_IGUAL = r'='
# -------------- -> TOKEN ARITMETICOS <- --------------
t_MAS = r'\+'
t_MENOS = r'\-'
t_POR = r'\*'
t_DIV = r'\/'
t_MODULO = r'\%'
t_POT = r'\^'
# -------------- -> TOKEN RELACIONAL <- --------------
t_MAYORQ = r'>'
t_MENORQ = r'<'
t_MAYORI = r'>='
t_MENORI = r'<='
t_COMPARACION = r'=='
t_DIFERENTE = r'!='
# -------------- -> TOKEN LOGICO <- --------------
t_OR = r'\|\|'
t_AND = r'&&'
t_NOT = r'!'
# -------------- -> TOKEN DECIMAL <- --------------
def t_DECIMAL(t):
    r'\d+\.\d+'
    try:
        t.value = float(t.value)
    except ValueError:
        print('Float value too large %d', t.value)
        t.value = 0
    return t


# -------------- -> TOKEN ENTERO <- --------------
def t_ENTERO(t):
    r'\d+'
    try:
        t.value = int(t.value)
    except ValueError:
        print('Integer value too large %d', t.value)
        t.value = 0
    return t


# -------------- -> TOKEN ID <- --------------
def t_ID(t):
    r'[a-zA-Z_][a-zA-Z_0-9]*'
    t.type = reservadas.get(t.value, 'ID')  # Check for reserved words
    return t


# -------------- -> TOKEN CADENA <- --------------
def t_CADENA(t):
    r'\"(\\"|.)*?\"'
    t.value = t.value[1:-1]
    t.value = t.value.replace('\\n', '\n')
    t.value = t.value.replace('\\r', '\r')
    t.value = t.value.replace('\\\\', '\\')
    t.value = t.value.replace('\\"', '\"')
    t.value = t.value.replace('\\t', '\t')
    t.value = t.value.replace("\\'", '\'')
    return t


# ==============================================================================
# COMENTARIOS
# ==============================================================================
# Múltiples líneas #= ... =#
def t_COMENTARIO_MULTILINEA(t):
    r'\#=(.|\n)*?=\#'
    t.lexer.lineno += t.value.count('\n')


# Simple # ...
def t_COMENTARIO_SIMPLE(t):
    r'(\#.*\n)|(\#.*)'
    t.lexer.lineno += 1
    pass


# ==============================================================================
# OTROS
# ==============================================================================
# Caracteres ignorados
t_ignore = ' \t'


def t_newline(t):
    r'\n+'
    t.lexer.lineno += t.value.count("\n")


def t_error(t):
    if t.value[0] != ' ':
        print(t.value[0])
        print(
            "Illegal character '%s'" % t.value[0],
            t.lexer.lineno,
            (find_column(input_data, t)),
        )
    t.lexer.skip(1)


def find_column(inp, token):
    line_start = inp.rfind('\n', 0, token.lexpos) + 1
    return (token.lexpos - line_start) + 1


import src.interprete.ply.lex as lex

lexer = lex.lex()

# -------------- -> PRECENDENCIA <- --------------
precedence = (
    ('left', 'OR'),
    ('left', 'AND'),
    (
        'left',
        'COMPARACION',
        'DIFERENTE',
        'MAYORQ',
        'MENORQ',
        'MAYORI',
        'MENORI',
    ),
    ('left', 'MAS', 'MENOS'),
    ('left', 'DIV', 'POR', 'MODULO'),
    ('nonassoc', 'POT'),
    ('right', 'UNOT'),
    ('right', 'UMENOS'),
)

# ==============================================================================


# ANALISIS SINTACTICO


# ==============================================================================


def p_init(t):
    'init               : instrucciones'
    t[0] = t[1]


def p_instrucciones_lista(t):
    'instrucciones      : instrucciones instruccion'
    t[1].append(t[2])
    t[0] = t[1]


def p_instrucciones_instruccion(t):
    'instrucciones      : instruccion'
    t[0] = [t[1]]


# ==============================================================================
# INSTRUCCION
# ==============================================================================
def p_inst(t):
    '''
    instruccion         : print_inst fin_inst
    '''
    t[0] = t[1]


def p_fin_inst(t):
    '''
    fin_inst            : PCOMA
                        |
    '''
    if len(t) == 2 or len(t) == 1:
        t[0] = ';'


# ------------------------------------------------------------------------------
# PRINTLN
def p_inst_print(t):
    '''
    print_inst          : RPRINT PARA expresion_list PARC
                        | RPRINT PARA PARC
    '''
    if len(t) == 5:
        t[0] = Print(t[3], t.lineno(1), find_column(input_data, t.slice[1]))
    else:
        t[0] = Print([], t.lineno(1), find_column(input_data, t.slice[1]))


def p_inst_println(t):
    '''
    print_inst          : RPRINTLN PARA expresion_list PARC
                        | RPRINTLN PARA PARC
    '''
    if len(t) == 5:
        t[0] = Print(
            t[3], t.lineno(1), find_column(input_data, t.slice[1]), True
        )
    else:
        t[0] = Print([], t.lineno(1), find_column(input_data, t.slice[1]), True)


# ==============================================================================
# EXPRESIONES
# ==============================================================================
def p_exp_unaria(t):
    '''
    expresion           : MENOS expresion %prec UMENOS
                        | NOT expresion %prec UNOT
    '''
    if t[1] == '-':
        t[0] = Aritmetica(
            TipoArtimetico.UMENOS,
            None,
            t[2],
            t.lineno(1),
            find_column(input_data, t.slice[1]),
        )
    elif t[1] == '!':
        t[0] = Logica(
            TipoLogico.NOT,
            None,
            t[2],
            t.lineno(1),
            find_column(input_data, t.slice[1]),
        )


def p_exp_agrupacion(t):
    'expresion          : PARA expresion PARC'
    t[0] = t[2]


def p_exp_list(t):
    '''
    expresion_list      : expresion_list COMA expresion
                        | expresion
    '''
    if len(t) == 2:
        t[0] = [t[1]]
    else:
        t[1].append(t[3])
        t[0] = t[1]


def p_exp_binaria(t):
    '''
    expresion           : expresion MAS expresion
                        | expresion MENOS expresion
                        | expresion POR expresion
                        | expresion DIV expresion
                        | expresion POT expresion
                        | expresion MODULO expresion
                        | expresion MAYORQ expresion
                        | expresion MENORQ expresion
                        | expresion MAYORI expresion
                        | expresion MENORI expresion
                        | expresion COMPARACION expresion
                        | expresion DIFERENTE expresion
                        | expresion OR expresion
                        | expresion AND expresion
    '''
    # -------------- -> (+) <- --------------
    if t[2] == '+':
        t[0] = Aritmetica(
            TipoArtimetico.SUMA,
            t[1],
            t[3],
            t.lineno(2),
            find_column(input_data, t.slice[2]),
        )

    # -------------- -> (-) <- --------------
    elif t[2] == '-':
        t[0] = Aritmetica(
            TipoArtimetico.RESTA,
            t[1],
            t[3],
            t.lineno(2),
            find_column(input_data, t.slice[2]),
        )

    # -------------- -> (*) <- --------------
    elif t[2] == '*':
        t[0] = Aritmetica(
            TipoArtimetico.POR,
            t[1],
            t[3],
            t.lineno(2),
            find_column(input_data, t.slice[2]),
        )

    # -------------- -> (/) <- --------------
    elif t[2] == '/':
        t[0] = Aritmetica(
            TipoArtimetico.DIV,
            t[1],
            t[3],
            t.lineno(2),
            find_column(input_data, t.slice[2]),
        )

    # -------------- -> (^) <- --------------
    elif t[2] == '^':
        t[0] = Aritmetica(
            TipoArtimetico.POT,
            t[1],
            t[3],
            t.lineno(2),
            find_column(input_data, t.slice[2]),
        )

    # -------------- -> (%) <- --------------
    elif t[2] == '%':
        t[0] = Aritmetica(
            TipoArtimetico.MODULO,
            t[1],
            t[3],
            t.lineno(2),
            find_column(input_data, t.slice[2]),
        )

    # -------------- -> (>) <- --------------
    elif t[2] == '>':
        t[0] = Relacional(
            TipoRelational.MAYORQ,
            t[1],
            t[3],
            t.lineno(2),
            find_column(input_data, t.slice[2]),
        )

    # -------------- -> (<) <- --------------
    elif t[2] == '<':
        t[0] = Relacional(
            TipoRelational.MENORQ,
            t[1],
            t[3],
            t.lineno(2),
            find_column(input_data, t.slice[2]),
        )

    # -------------- -> (>=) <- --------------
    elif t[2] == '>=':
        t[0] = Relacional(
            TipoRelational.MAYORI,
            t[1],
            t[3],
            t.lineno(2),
            find_column(input_data, t.slice[2]),
        )

    # -------------- -> (<=) <- --------------
    elif t[2] == '<=':
        t[0] = Relacional(
            TipoRelational.MENORI,
            t[1],
            t[3],
            t.lineno(2),
            find_column(input_data, t.slice[2]),
        )

    # -------------- -> (==) <- --------------
    elif t[2] == '==':
        t[0] = Relacional(
            TipoRelational.COMPARACION,
            t[1],
            t[3],
            t.lineno(2),
            find_column(input_data, t.slice[2]),
        )

    # -------------- -> (!=) <- --------------
    elif t[2] == '!=':
        t[0] = Relacional(
            TipoRelational.DIFERENTE,
            t[1],
            t[3],
            t.lineno(2),
            find_column(input_data, t.slice[2]),
        )

    # -------------- -> (||) <- --------------
    elif t[2] == '||':
        t[0] = Logica(
            TipoLogico.OR,
            t[1],
            t[3],
            t.lineno(2),
            find_column(input_data, t.slice[2]),
        )

    # -------------- -> (&&) <- --------------
    elif t[2] == '&&':
        t[0] = Logica(
            TipoLogico.AND,
            t[1],
            t[3],
            t.lineno(2),
            find_column(input_data, t.slice[2]),
        )


def p_exp_fin(t):
    '''
    expresion           : ENTERO
                        | DECIMAL
                        | RTRUE
                        | RFALSE
                        | CADENA
    '''
    if len(t) == 2:
        if t.slice[1].type == 'ENTERO':
            t[0] = Primitivo(
                int(t[1]),
                TipoVar.INT64,
                t.lineno(1),
                find_column(input_data, t.slice[1]),
            )
        elif t.slice[1].type == 'DECIMAL':
            t[0] = Primitivo(
                float(t[1]),
                TipoVar.FLOAT64,
                t.lineno(1),
                find_column(input_data, t.slice[1]),
            )
        elif isinstance(t[1], str):
            value = t[1]
            if value == 'true':
                t[0] = Primitivo(
                    True,
                    TipoVar.BOOLEAN,
                    t.lineno(1),
                    find_column(input_data, t.slice[1]),
                )
            elif value == 'false':
                t[0] = Primitivo(
                    False,
                    TipoVar.BOOLEAN,
                    t.lineno(1),
                    find_column(input_data, t.slice[1]),
                )
            else:
                t[0] = Primitivo(
                    str(value),
                    TipoVar.STRING,
                    t.lineno(1),
                    find_column(input_data, t.slice[1]),
                )


# ==============================================================================
# PARSE
# ==============================================================================
import src.interprete.ply.yacc as yacc

parser = yacc.yacc()
input_data: str = ''


def parse(input_string: str):
    global lexer
    global parser
    lexer = lex.lex()
    parser = yacc.yacc()
    global input_data
    input_data = input_string
    return parser.parse(input_string)
