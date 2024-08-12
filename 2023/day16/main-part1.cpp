#include <iostream>
#include <vector>
#include <fstream>
#include <string>
#include <map>

struct Beam {
    int row;
    int col;
};

enum Direction {
    up = 0, down, left, right
};

int main() {
    std::ifstream file("test_input.txt");

    std::vector<std::string> vectorOfStrings;
    std::string line;
    while(std::getline(file, line)) {
        vectorOfStrings.push_back(line);
    }

    int countOfHashes = 1;
    std::map<Beam, Direction> beamToDirMap;

    Beam firstBeam = { 0, 0};
    beamToDirMap[firstBeam] = right;

    //TODO: iterate thr map and add new beams in it if needed, but newly added elements in map are not considered in current iteration, so do something about it.


    std::cout << "Ans: " << countOfHashes << std::endl;
    file.close();
}
