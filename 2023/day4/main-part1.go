package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    file, _ := os.Open("puzzle_input.txt")
    defer file.Close()

    scn := bufio.NewScanner(file)

    finalScore := 0
    for scn.Scan() {
        score := 0
        line := scn.Text()

        indexOfColon := strings.Index(line, ":")
        line = line[indexOfColon+1:]

        indexOfBar := strings.Index(line, "|")
        numbersWeHave := line[:indexOfBar]
        winningNumbers := line[indexOfBar+1:]

        // fmt.Println("numbers we have: ", numbersWeHave)
        // fmt.Println("winning numbers : ", winningNumbers)

        numsWeHave_slice := strings.Split(numbersWeHave, " ")
        winningNums_slice := strings.Split(winningNumbers, " ")

        fmt.Printf("%q\n", numsWeHave_slice)
        fmt.Printf("%q\n", winningNums_slice)

        for i := 0; i < len(numsWeHave_slice); i++ {
            if numsWeHave_slice[i] == "" {
                continue
            }

            //loop for winning numbers
            for j := 0; j < len(winningNums_slice); j++ {
                if numsWeHave_slice[i] == winningNums_slice[j] {
                    if score == 0 {
                        score = 1
                    } else {
                        score = score * 2
                    }
                    break
                }
            }
        }

        finalScore += score
    }
    
    fmt.Println("Final score: ", finalScore)
}
