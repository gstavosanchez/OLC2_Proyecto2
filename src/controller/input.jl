# arr = [0, 2, 3, 4, 5];

# i = 1;
# while i <= 5
#     println(arr[i])
#     i = i + 1;
# end;

# arr[1] = 1;
# println(arr[1]);

# arr[4] = [11, 20, 30, 40];
# println(arr[4][1]);
# println(arr[4][3]);

# arr[4][1] = 10;
# println(arr[4][1]);

# #=
#     ---------------------------------------
#             otro arreglo
#     ---------------------------------------
# =#
# array = [
#         [12, 9, 4, 99, 56, 34 , 78, 22, 1, 3, 13, 120],
#         [32, 7*3, 7, 2, 9, 9874^0, 44, 3, 820*10, 11, 8*0+8, 10]
#         ];
# # Cada arreglo interno es de tamaño 12

# i = 1;
# while i <= 2            #tamaño del arreglo padre
#     j = 0;
#     while j <= 12       #tamaño de cada arreglo hijo
#         print(array[i][j]);
#         print(" - ");
#         j = j + 1;
#     end;
#     println("");
#     i = i + 1;
# end;

# #=
#     ---------------------------------------
#             ultimo arreglo
#     ---------------------------------------
# =#

# arrD = [12, 9, 4, 99, 56, 34 , 78, 22, 1, 3, 13, 120, 32, 7*3, 7, 2, 9, 9874^0, 44, 3, 820*10, 11, 8*0+8, 10];
# ## tamaño de 24

# i = 25;
# while i > 0
#     j = 1;
#     while j < i - 1
#         if arrD[j] > arrD[j+1]
#             aux = arrD[j];
#             arrD[j] = arrD[j+1];
#             arrD[j+1] = aux;
#         end;
#         j = j + 1;
#     end;
#     i = i - 1;
# end;
# println("Fin de ordenamiento");

# i = 1;
# while i <= 24
#     print(arrD[i]);
#     print(" - ");
#     i = i + 1;
# end;



# function insertionSort(arr::Vector{Int64}) 

#     for i in 2:6
#         j = i;
#         temp = arr[i];
#         while(j > 1 && arr[j - 1] > temp)
#             arr[j] = arr[j-1];
#             j = j - 1;
#         end;
#         arr[j] = temp;
#     end;
# end;


# arreglo = [100,50,25,150,1,5]::Vector{Int64};
# insertionSort(arreglo);
# for x in 1:6
# 	print(arreglo[x], " ");
# end;

# array = [
#         [12,9,4,99,56,34,78,22,1,3,10,13,120]
        
#     ]::Vector{Vector{Int64}};


# mutable struct Persona
#     nombre::String;
#     edad::Int64;
#     numeroFamiliares::Int64;
# end;

# function RegistrarPersona(nombre::String, edad::Int64, numeroFamiliares::Int64)::Persona
#     return Persona(nombre, edad, numeroFamiliares);
# end;

# function AgregarFamiliar(persona::Persona)
#     persona.numeroFamiliares = persona.numeroFamiliares + 1;
# end;

# function ImprimirDatosPersona(persona::Persona)
#     println("Nombre: ", persona.nombre);
#     println("Edad: ", persona.edad);
#     println("Numero de familiares: ", persona.numeroFamiliares);
# end;

# manuel = RegistrarPersona("Manuel", 22, 4);
# ImprimirDatosPersona(manuel);
# AgregarFamiliar(manuel);
# ImprimirDatosPersona(manuel);

# struct Cartelera
#     peliculas::Vector{String};
# end;

# cinepolis = Cartelera(["El Padrino", "El Padrino 2", "El Padrino 3"]);
# for pelicula in cinepolis.peliculas
#     println(pelicula);
# end;

# function quicksort(array::Vector{Int64},low::Int64,n::Int64)::Int64
# 	mid = array[(trunc(((low + n) / 2)))]::Int64;
# end;

# struct Cartelera
#     peliculas::Vector{String};
# end;

# cinepolis = Cartelera(["Pelicula 1", "Pelicula 2", "Pelicula 3"]);

# for pelicula in ["Pelicula 1", "Pelicula 2", "Pelicula 3", "Pelicula 3555"]
#     println(pelicula);
# end;


# function quicksort(a::Int64)::Int64
# 	mid = a::Int64;
#   	println(a);
# end;
# array = [1, 2, 3]::Vector{Vector{Vector{String}}}




# function swap(i::Int64, j::Int64, arr::Vector{Int64}) 
#     temp = arr[i];
#     arr[i] = arr[j];
#     arr[j] = temp;
# end;

# function bubbleSort(arr::Vector{Int64})
#     for i in 0:(length(arr) - 1)
#         for j in 1:(length(arr) - 1)
#             if(arr[j] > arr[j + 1])
#                 swap(j, j+1, arr);
#             end;
#         end;
#     end;
# end;

# arreglo = [32,7*3,7,89,56,909,109,2,9,9874^0,44,3,820*10,11,8*0+8,10] :: Vector{Int64};

# println("BubbleSort => Antes del ordenamiento");
# for x in 1:length(arreglo)
# 	print(arreglo[x], " ");
# end;
# println(" ");
# # bubbleSort(arreglo);
# println("BubbleSort => Despues del ordenamiento");
# for x in 1:length(arreglo)
# 	print(arreglo[x], " ");
# end;



function quicksort(array::Vector{Int64},low::Int64,n::Int64)::Int64
	lo = low::Int64;
	hi = n::Int64;
	if lo >= n
		return 0;
    end;
	mid = array[(trunc(((lo + hi) / 2)))]::Int64;
	while lo < hi
		while (lo<hi && array[lo] < mid)
			lo = lo + 1;
        end;
		while (lo<hi && array[hi] > mid)
			hi = hi - 1;
        end;
		if lo < hi
			T = array[lo]::Int64;
			array[lo] = array[hi];
			array[hi] = T;
        end;
    end;
	if (hi < lo)
		T = hi::Int64;
		hi = lo;
		lo = T;
    end;
	quicksort(array,low,lo);
	cond = 0::Int64;
	if ( lo == low)
		cond = lo + 1;
	else
		cond = lo;
    end;
	quicksort(array,cond,n);
end;

i = 0::Int64;

array = [
        [12,9,4,99,56,34,78,22,1,3,10,13,120],
        [32,7*3,7,89,56,909,109,2,9,9874^0,44,3,820*10,11,8*0+8,10]
    ]::Vector{Vector{Int64}};

println("Quick Sort");
println("Valores antes de Quicksort");
for x in 1:length(array[1])
	print(array[1][x],"	");
end;
println("");

println("-------------------------");
quicksort(array[1],1,length(array[1]));
println("Valores despues de QuickSort:");
for y in 1:length(array[1])
	print(array[1][y], "	");
end;
println("");

println("Valores antes de Quicksort");
for x in 1:length(array[2])
	print(array[2][x], "	");
end;
println("");

println("-------------------------");
quicksort(array[2],1,length(array[2]));
println("Valores despues de QuickSort:");
for y in 1:length(array[2])
	print(array[2][y],"	");
end;



