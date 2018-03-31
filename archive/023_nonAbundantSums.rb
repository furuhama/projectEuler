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

    if array.sum > self
      return true
    else
      return false
    end
  end

  def is_sum_of_abundant?
    ABUNDANT_NUMBERS.each do |abundant|
      ABUNDANT_NUMBERS.include?(self - abundant)
    end
  end
end

if __FILE__ == $0
  # largest number that cannot be sum of two abundant number is 28123
  # smallest number that can be sum of two abundant number is 12
  # ABUNDANT_NUMBERS = (12..28112).select { |i| i.is_abundant? }
  # NOT_SUM_OF_TWO_ABUNDANT_NUMBERS = (24..28124).select { |i| !i.is_sum_of_abundant? }

  ABUNDANT_NUMBERS = (12..30).select { |i| i.is_abundant? }
  # NOT_SUM_OF_TWO_ABUNDANT_NUMBERS = (24..100).select { |i| !i.is_sum_of_abundant? }
  (15..40).each do |i|
    puts i if i.is_sum_of_abundant?
  end

  result = (1..23).to_a.sum + NOT_SUM_OF_TWO_ABUNDANT_NUMBERS.sum
  puts result
end
