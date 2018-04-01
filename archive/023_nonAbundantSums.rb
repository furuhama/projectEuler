class Integer
  def is_abundant?
    array = [1]

    (2..Math.sqrt(self)).each do |i|
      if self%i == 0
        array << i

        if self/i != i
          array << self/i
        end
      end
    end

    return array.sum > self
  end

  def is_sum_of_abundant?
    ABUNDANT_NUMBERS.each do |abundant|
      break if abundant > ABUNDANT_NUMBERS.last / 2
      return true if ABUNDANT_NUMBERS.include?(self - abundant)
    end

    false
  end
end

if __FILE__ == $0
  # largest number that cannot be sum of two abundant number is 28123
  # smallest number that can be sum of two abundant number is 12
  ABUNDANT_NUMBERS = (12..28112).select { |i| i.is_abundant? }
  NOT_SUM_OF_TWO_ABUNDANT_NUMBERS = (24..28124).select { |i| !i.is_sum_of_abundant? }

  result = (1..23).to_a.sum + NOT_SUM_OF_TWO_ABUNDANT_NUMBERS.sum
  puts result
end
