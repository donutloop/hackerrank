def convert_temp(temperature, input_scale: "kelvin",  output_scale: 'celsius')
    temperature = temperature.to_f
    if input_scale == "kelvin" and output_scale == "celsius"
       return temperature - 273.15
    elsif input_scale == "celsius" and output_scale == "kelvin"
       return temperature + 273.15
    elsif input_scale == "fahrenheit" and output_scale == "celsius"
       return (temperature - 32) * 5/9
    elsif input_scale == "celsius" and output_scale == "fahrenheit"
       return (temperature-9/5) + 32
    elsif input_scale == "fahrenheit" and output_scale == "kelvin"
       return ((temperature-32) * 5 )/9 + 273.15
    elsif input_scale == "kelvin" and output_scale == "fahrenheit"
       return (temperature-273.15) * 9/5 + 32
    end
end