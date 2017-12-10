#include <iostream>

int main() {
    int a = 3;
    int b = 5;
    int max = 1000;
    int sum = 0;

    for (int i = 1; i < max; i++) {
        if (i % (a * b) == 0) {
            sum += i;
        }
        else if ((i % a) == 0) {
            sum += i;
        }
        else if ((i % b) == 0) {
            sum += i;
        }
    }
    std::cout << sum << std::endl;
}

