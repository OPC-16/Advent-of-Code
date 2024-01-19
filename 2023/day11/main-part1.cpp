#include <iostream>
#include <string> //for std::string
#include <utility> //for std::pair
#include <vector> //for std::vector
#include <fstream>
#include <cstdlib> //for abs

void insertAdditionalColumn(std::vector<std::string>& vec, int pos) {
    for (int i=0; i<vec.size(); i++) {
        vec[i].insert(pos, ".");
    }
}

int main() {
    std::ifstream file("puzzle_input.txt");

    std::string line;
    std::vector<std::string> vec;

    while (std::getline(file, line)) {
        vec.push_back(line);

        std::size_t found = line.find("#");
        if (found == std::string::npos) {
            // int lengthOfString = line.length();
            // std::string stringToBeAdded(lengthOfString, '.');
            // vec.push_back(stringToBeAdded);

            vec.push_back(line);
        }

    }

    //checking columns for emptiness
    int sizeOfString = vec[0].length();
    for (int i=0; i<=sizeOfString; i++) { // i will go horizontally right
        bool containHashTag = false;
        for (int j=0; j<vec.size(); j++) { // j will go vertically down
            if (vec[j][i] == '#') {
                containHashTag = true;
                break;
            }
        }

        if (containHashTag == false) {
            insertAdditionalColumn(vec, i+1);
            i++;
        }
    }

    std::vector<std::pair<int, int>> posOfHash;  //vector for positions of hashes
    sizeOfString = vec[0].length();
    for (int i=0; i<vec.size(); i++) {
        for (int j=0; j<=sizeOfString; j++) {
            if (vec[i][j] == '#') {
                std::pair<int, int> p = {i, j};
                posOfHash.push_back(p);
            }
        }
    }

    //printing vector of strings
    // for (int i=0; i<vec.size(); i++) {
    //     std::cout << vec[i] << std::endl;
    // }

    //printing the positions of hashes
    // for (int i=0; i<posOfHash.size(); i++) {
    //     std::cout << "Row: "<< posOfHash[i].first << ", Col: " << posOfHash[i].second << std::endl;
    // }

    int finalAns = 0;
    for (int i=0; i<posOfHash.size(); i++) {
        for (int j=i+1; j<posOfHash.size(); j++) {
            std::pair<int, int> p1 = posOfHash[i];
            std::pair<int, int> p2 = posOfHash[j];
            finalAns += abs(p1.first - p2.first) + abs(p1.second - p2.second);
        }
    }

    std::cout << "Final ans: " << finalAns << std::endl;

    file.close();

    return 0;
}
