/* Copyright 2017 furuhama */

// this is a code for Project Euler q.10

#include <iostream>
#include <cmath>
#define N 2000000
using llong = int64_t;

void Eratosthenes(int n, int *arr) {
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
}

int main() {
    int arr[N];
    Eratosthenes(N, arr);
    // should use llong type (instead of int)
    llong sumOfPrimes = 0;
    for (int i = 2; i < N; i++) {
        if (arr[i] == 1) {
            sumOfPrimes = sumOfPrimes + i;
        }
    }
    std::cout << "the answer is: " << sumOfPrimes << std::endl;
    return 0;
}
