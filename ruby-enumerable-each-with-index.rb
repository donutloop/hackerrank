def skip_animals(animals, skip)
  items = []  
  animals.each_with_index do |item, index|
        if index >= skip
            items.push("#{index}:#{item}")
        end    
  end
    
  return items
end