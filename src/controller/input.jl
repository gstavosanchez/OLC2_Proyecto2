function quicksort(array::Vector{Int64},low::Int64,n::Int64)::Int64
	mid = array[(trunc(((low + n) / 2)))]::Int64;
end;


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