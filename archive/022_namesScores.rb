# frozen_string_literal: true

ALPHABET_ARRAY = ('A'..'Z').to_a

class String
  # name -> score
  def score
    score = 0
    each_char do |char|
      # score should be index + 1
      score += ALPHABET_ARRAY.find_index { |element| element == char } + 1
    end

    score
  end
end

class Array
  # score -> total_score
  def total_score
    tscore = 0
    each_with_index do |str, idx|
      tscore += str.score * (idx + 1)
    end

    tscore
  end
end

if __FILE__ == $0
  File.open('names.txt') do |file|
    file.each_line do |line|
      names_array = line.gsub(/"/, '').split(',').sort

      p names_array.total_score
    end
  end
end
