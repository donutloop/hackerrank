def square_of_sum (my_array, proc_square, proc_sum)
    sum = proc_sum.call(my_array)
    square = proc_square.call(sum)
    
    return square
end
proc_square_number = proc {|x| x * x}


proc_sum_array     = Proc.new do |numbers|
  sum = 0
  numbers.each do |n|
    sum = sum + n
  end
  next sum
end

my_array = gets.split().map(&:to_i)

puts square_of_sum(my_array, proc_square_number, proc_sum_array)
