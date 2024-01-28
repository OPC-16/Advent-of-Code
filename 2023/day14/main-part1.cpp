#include <iostream>
#include <string>
#include <vector>
#include <fstream>

int main() {
    std::vector<std::string> board;
    std::ifstream file("puzzle_input.txt");

    std::string line;
    while (std::getline(file, line)) {
        board.push_back(line);
    }

    int boardLength = board.size();
    int stringLength = board[0].length();
    for (int i = 1; i < boardLength; i++) {  // i will traverse the board string-wise
        for (int j = 0; j < stringLength; j++) { // j will traverse string char-wise
            char c;
            if (board[i][j] == 'O') {
                c = 'O';
            } else {
                continue;
            }

            //check curr char's upper position and move it till it gets obstructed
            int temp_i = i - 1;
            while (temp_i >= 0 && board[temp_i][j] == '.') {
                board[temp_i][j] = c;
                board[temp_i+1][j] = '.';
                temp_i--;
            }
        }
    }

    int finalAns = 0;
    int rowCounter = 1;
    for (int i = boardLength - 1; i >= 0; i--) {
        int stoneCounter = 0;
        for (int j = 0; j < stringLength; j++) {
            if (board[i][j] == 'O') {
                stoneCounter++;
            }
        }
        finalAns += rowCounter * stoneCounter;
        rowCounter++;
    }

    std::cout << finalAns << std::endl;
}
