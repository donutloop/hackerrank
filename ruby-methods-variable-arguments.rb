def full_name (first_name, *middle_name)
    middle_name.each do |name|
        first_name = first_name + " " + name
    end
    return first_name
end