package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calculateHash(s string) int64 {
    var currValue int64 = 0
    for i := 0; i < len(s); i++ {
        currValue += int64(s[i])
        currValue *= 17
        currValue %= 256
    }

    return currValue
}

func main() {
    file, _ := os.Open("puzzle_input.txt")
    defer file.Close()

    scn := bufio.NewScanner(file)

    scn.Scan()
    line := scn.Text()

    var finalAns int64 = 0
    steps := strings.Split(line, ",")
    stringToIntMap := make(map[string]int64)

    for i := 0; i < len(steps); i++ {
        if _, ok := stringToIntMap[steps[i]]; !ok {
            ans := calculateHash(steps[i])
            stringToIntMap[steps[i]] = ans
            finalAns += ans
        } else {
            finalAns += stringToIntMap[steps[i]]
        }
    }


    fmt.Println("Final Ans:", finalAns)
}
