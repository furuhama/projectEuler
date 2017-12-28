/* copyright furuhama 2017 */

// this is a code for Project Euler No.6

#include <iostream>

int getSquareOfNaturalNum(int max) {
    return ((1 + max) * max * (1 + max) * max / 4);
}

int main() {
    int max = 100;
    int sum = 0;
    sum += getSquareOfNaturalNum(max);

    for (int i = 1; i <= max; i++) {
        sum -= i * i;
    }

    std::cout << "the answer is " << sum << std::endl;
    return sum;
}
