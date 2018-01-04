# this is a code for Project Euler q.17

def convertNumberToLetter(num_int)
  letter_high = getLetterOver99(num_int/100)
  letter_low = getLetterUnder100(num_int%100)

  if letter_high == nil
    return letter_low
  elsif letter_low == nil
    return letter_high
  else
    return letter_high + "and" + letter_low
  end
end

def getLetterOver99(num)
  case num
  when 0
    return nil
  when 1
    return "onehundred"
  when 2
    return "twohundred"
  when 3
    return "threehundred"
  when 4
    return "fourhundred"
  when 5
    return "fivehundred"
  when 6
    return "sixhundred"
  when 7
    return "sevenhundred"
  when 8
    return "eighthundred"
  when 9
    return "ninehundred"
  when 10
    return "onethousand"
  end
end

def getLetterUnder100(num)
  if num == 0
    return nil
  elsif num == 10
    return "ten"
  elsif num == 11
    return "eleven"
  elsif num == 12
    return "twelve"
  elsif num == 13
    return "thirteen"
  elsif num == 14
    return "fourteen"
  elsif num == 15
    return "fifteen"
  elsif num == 16
    return "sixteen"
  elsif num == 17
    return "seventeen"
  elsif num == 18
    return "eighteen"
  elsif num == 19
    return "nineteen"
  end

  result = ""
  
  case num / 10
  when 2
    result += "twenty"
  when 3
    result += "thirty"
  when 4
    result += "forty"
  when 5
    result += "fifty"
  when 6
    result += "sixty"
  when 7
    result += "seventy"
  when 8
    result += "eighty"
  when 9
    result += "ninety"
  end

  case num % 10
  when 0
    result += ""
  when 1
    result += "one"
  when 2
    result += "two"
  when 3
    result += "three"
  when 4
    result += "four"
  when 5
    result += "five"
  when 6
    result += "six"
  when 7
    result += "seven"
  when 8
    result += "eight"
  when 9
    result += "nine"
  end

  return result
end


if __FILE__ == $0
  count = 0
  (1..1000).each {|num|
    count += convertNumberToLetter(num).length
  }
  p count
end