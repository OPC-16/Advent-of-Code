#include <iostream>

int main() {
    long int time = 53916768;
    long long int distance = 250133010811025;
    int numOfWays = 0;

    for (int t=0; t <= time; t++) {
        long long int product = t * (time - t);
        if (product > distance)
            numOfWays++;
    }

    std::cout << "Final ans: " << numOfWays << std::endl;
}
