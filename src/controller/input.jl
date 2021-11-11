array = [
        [12,9,4,99,56,34,78,22,1,3,10,13,120],
        [32,7*3,7,89,56,909,109,2,9,9874^0,44,3,820*10,11,8*0+8,10]
    ]::Vector{Vector{Int64}};

println("Quick Sort");
println("Valores antes de Quicksort");
for x in 1:length(array[1])
	if x != length(array[1])
		print(array[1][x], "-");
  else
    print(array[1][x]);
  end;
end;
println("");