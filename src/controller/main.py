from src.interprete.compilador.simbolos.Generador import Generador
from src.interprete.ast.TreeAST import TreeAST
from src.interprete.compilador.simbolos.Entorno import Entorno
from src.interprete.gramatica.gramatica import *


def execute(input_str: str):

    generator = Generador.get_instance()
    generator.clean()
    # -------------- -> EXECUTE ANALISIS <- --------------
    global_enviroment: Entorno = Entorno()
    inst_list = parse(input_str)
    # -------------- -> EXECUTE COMPILADOR <- --------------
    ast = TreeAST(inst_list, 0, 0)
    ast.compilar(global_enviroment)

    str_output = str(generator.get_code())
    print('-------------- -> SALIDA <- --------------')
    print(f'Entrada: {input_str}')
    print('')
    print(str_output)
    print('------------------------------------------')
    return {'compilador': str_output}


def get_proyect_data():
    return {
        'Curso': 'Organizacion de lenguajes y compiladores 2',
        'Proyecto': 'Proyecto No.2',
        'Nombre': 'Elmer Gustavo Sanchez Garcia',
        'Carnet': '201801351',
    }


def dev_compilier():
    f = open(
        'C:\\Users\\elmer\\Downloads\\2S2021\\Compiladores 2\\Proyecto_2\\src\\controller\\input.jl',
        'r',
    )
    input_str = f.read()
    return execute(input_str)
    # execute(input_str)
