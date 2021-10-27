# a = 1;
# for i in 1:4 
#     a = a + i;
#     println(a); 
# end;

# Esto no funciona
# function sumar(a::Float64):: Float64
#     b = 2.0
#     for i in 1:4
#         b = b + a * i
#     end;
#     return b;
# end;

# a = sumar(5.0);
# println(a);


# 5 + 5 = float
# if "a"  == "a"
# println("ha" ^ 3);

# a = [1, 2, 3, 4, 5, [10, 25, [50, 20, 200]]]

# a = [100, 200, 300]
# a = ["h", "o", "l", "a", ["ja", "ja"]]
# a = [1, 2, 3, 4, 5, [10, 25, [50, 25, 200]]];
# println(a[6][100])
# println("gus")

mutable struct Persona
    edad::Int64;
    nombre::String;
end;
mutable struct Alumno
    carrera::String;
    people:: Persona;
end;
a = Persona(23, "Gus");
println(a.edad)

estudiante = Alumno("sistemas", a)

println(estudiante.carrera)
println(estudiante.people.nombre)

estudiante.people.nombre = 5
println(a.nombre)

# acceso a persona.edad
# t3 = stack[0];
# t3 = t3 + 0;
# t4 = heap[t3];

# acceso a alumuno.people.edad
# t3 = stack[0];
# t3 = t3 + 1;

# t4 = heap[t3]
# t5 = t5 + 0

# t6 = heap[t5]

