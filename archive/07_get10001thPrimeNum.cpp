/* Copyright 2017 furuhama */

// this is a code for Project Euler No.7

#include <iostream>
#include <cmath>
#define SIZE 10001

int Eratosthenes(int n) {
  int arr[n];
  int count = 1;
  int maxPrime = 2;
  int *maxPrimePointer;
  maxPrimePointer = &maxPrime;

  for (int i = 0; i < n; i++) {
    arr[i] = 1;
  }
  for (int i = 2; i < sqrt(n); i++) {
    if (arr[i]) {
      for (int j = 0; i * (j + 2) < n; j++) {
        arr[i * (j + 2)] = 0;
      }
    }
  }
  for (int i = 2; i < n; i++) {
    if (arr[i]) {
      std::cout << count << ": " << i << std::endl;
      count++;
      *maxPrimePointer = i;
    }
  }
  return maxPrime;
}

int main() {
    Eratosthenes(110000);
    // the answer, the 10001st prime number is 104743
    return 0;
}
