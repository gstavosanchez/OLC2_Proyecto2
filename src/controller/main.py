from src.interprete.compilador.tipos.Tipo import get_tipo_var
from src.interprete.compilador.simbolos.Simbolo import Simbolo
from src.interprete.compilador.simbolos.Generador import Generador
from src.interprete.ast.TreeAST import TreeAST
from src.interprete.compilador.simbolos.Entorno import Entorno
from src.interprete.gramatica.compilador.gramatica import *


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
    str_error = str(generator.get_erro_str())
    print('-------------- -> ERRORES <- -------------')
    print(str_error)
    print('')
    print('-------------- -> SALIDA <- --------------')
    # print(f'Entrada: {input_str}')
    print('')
    print(str_output)
    print('------------------------------------------')
    return {
        'compilador': str_output,
        'error_str': str_error,
        'erro_list': generator.get_erro_json(),
        'tabla_simbolo': get_table_simbol(global_enviroment),
    }


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


def get_table_simbol(enviroment: Entorno):
    table_list: list = []
    for key in enviroment.variables:
        simbolo: Simbolo = enviroment.variables[key]
        objeto = {
            'nombre': key,
            'tipo': get_tipo_var(simbolo.get_type()),
            'ambito': 'Global',
        }
        table_list.append(objeto)

    for key in enviroment.functions:
        objeto = {'nombre': key, 'tipo': 'Funcion', 'ambito': 'Global'}
        table_list.append(objeto)
    for key in enviroment.structs:
        objeto = {'nombre': key, 'tipo': 'Struct', 'ambito': 'Global'}
        table_list.append(objeto)

    return table_list
