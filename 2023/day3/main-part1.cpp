#include <iostream>
#include <fstream>
#include <string>

int main() {
    std::ifstream file("test_input.txt");

    std::string line;
    while (std::getline(file, line)) {
        std::cout << line << std::endl;
    }

    file.close();
}
