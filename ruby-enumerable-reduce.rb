def sum_terms(n)
 (0..n).inject() {|sum, n| sum + (n*n+1)}  
end 
