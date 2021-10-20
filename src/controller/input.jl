# a = 1;
# for i in 1:4 
#     a = a + i;
#     println(a); 
# end;


function sumar(a::Int64):: Int64
    return a * sumar(a - 1);
end;

a = sumar(5);
println(a);