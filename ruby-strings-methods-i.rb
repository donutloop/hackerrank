def process_text(text)
    text = text.map{|textPart| textPart.strip}
    return text.join(" ")
end