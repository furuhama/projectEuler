def count_cycle(numerator, denominator, count, result_array)
  # calculate
  div_result = numerator / denominator
  extra_result = numerator % denominator

  # just divved
  return 0 if extra_result == 0

  # cycled
  return count if extra_result == 1 && div_result != 0

  # cycled
  search_result = result_array.select { |arr| arr[1] == div_result && arr[2] == extra_result }
  return count - search_result[0][0] unless search_result.empty?

  # update array
  result_array << [count, div_result, extra_result]
  # update numerator
  numerator = extra_result*10
  count += 1
  count_cycle(numerator, denominator, count, result_array)
end

def largest_cycle
  largest_den = 1
  largest_cycle = 1

  (2..999).each do |i|
    cycle = count_cycle(1, i, 0, [])
    if cycle > largest_cycle
      largest_den = i
      largest_cycle = cycle
    end
  end

  largest_den
end

if __FILE__ == $0
  p largest_cycle
end
