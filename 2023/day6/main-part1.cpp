#include <iostream>
#include <vector>

int main() {
    std::vector<int> timeVector;
    std::vector<int> distanceVector;

    timeVector.push_back(53);
    timeVector.push_back(91);
    timeVector.push_back(67);
    timeVector.push_back(68);

    distanceVector.push_back(250);
    distanceVector.push_back(1330);
    distanceVector.push_back(1081);
    distanceVector.push_back(1025);

    int numOfWays = 0, finalAns = 1;

    for (int i=0; i<4; i++) {
        for (int time=0; time <= timeVector[i];time++) {
            int product = time * (timeVector[i] - time);
            if (product > distanceVector[i])
                numOfWays++;
        }

        finalAns *= numOfWays;
        numOfWays = 0;
    }

    std::cout << "Final ans: " << finalAns << std::endl;
}
