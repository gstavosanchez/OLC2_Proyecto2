function insertionSort(arr::Vector{Int64}) 

    for i in 2:6
        j = i;
        temp = arr[i];
        while(j > 1 && arr[j - 1] > temp)
            arr[j] = arr[j-1];
            j = j - 1;
        end;
        arr[j] = temp;
    end;
end;


arreglo = [100,50,25,150,1,5]::Vector{Int64};
insertionSort(arreglo);
for x in 1:6
	print(arreglo[x], " ");
end;

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

# function bubbleSort(arr)
#     for i in 0:(length(arr) - 1)
#         for j in 1:(length(arr) - 1)
#             if(arr[j] > arr[j + 1])
#                 swap(j, j+1, arr);
#             end;
#         end;
#     end;
# end;

# arreglo = [32,7*3,7,89,56,909,109,2,9,9874^0,44,3,820*10,11,8*0+8,10] :: Vector{Int64};
# println("BubbleSort => ", arreglo);
# bubbleSort(arreglo);
# println("BubbleSort => ", arreglo);




