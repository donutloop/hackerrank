def count_multibyte_char(s)
   c = 0
   s.each_char {|x|
       if x.bytesize >= 2 then c = c + 1 end
   }
   return c
end

