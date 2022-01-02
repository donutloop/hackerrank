def group_by_marks(marks, pass_marks)
  h = {}
  groups = marks.group_by {|student, score|  score >= pass_marks }

  if groups.key?(true)
      h["Passed"] = groups[true]
  end

  if groups.key?(false)
      h["Failed"] = groups[false]
  end

  return h
end

