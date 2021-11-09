function hanoi(discos::Int64, origen::Int64, auxiliar::Int64, destino::Int64)
    if discos == 1
        println("Mover de ", origen, " a ", destino);
    else
        hanoi(discos - 1, origen, destino, auxiliar);
        println("Mover de ", origen, " a ", destino);
        hanoi(discos - 1, auxiliar, origen, destino);
    end;
end;
hanoi(3, 1, 2, 3);