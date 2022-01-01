def single_quote
  return 'Hello World and others!'
end

def double_quote
  return "Hello World and others!"
end

def here_doc
    doc = <<-HERE
    Hello World and others! 
    HERE
    
    return doc
end