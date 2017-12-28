#!/bin/sh

# return error when it gets too many args...
if [ $# -ge 2 ]
  then
    echo "An error occured. Too many args."
else
  # compile .cpp file as temp.out by clang++
  clang++ $1 -o temp.out -std=c++11

  # run compiled binary file
  ./temp.out

  # delete compiled binary file
  rm temp.out
fi

