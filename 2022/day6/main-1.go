package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    file, _ := os.Open("puzzle_input.txt")
    defer file.Close()

    scn := bufio.NewScanner(file)
    scn.Scan()
    line := scn.Text()

    fmt.Println(line)

    for i := 0; i < len(line) - 3; i++ {
        chars := make([]rune, 4)
        chars[0] = rune(line[i])
        chars[1] = rune(line[i+1])
        chars[2] = rune(line[i+2])
        chars[3] = rune(line[i+3])

        areCharsSame := checkSlice(chars)   //if true then there are duplicates and in case of false there will be no duplicates

        fmt.Printf("%c %c %c %c\n", chars[0], chars[1], chars[2], chars[3])
        if areCharsSame == false {
            fmt.Printf("the ans is %c on index %d\n", chars[3], i+4)
            break
        }
    }
} //function-main

func checkSlice(chars []rune) bool {
    for i := 0; i < len(chars); i++ {
        for j := i+1; j < len(chars); j++ {
            if chars[i] == chars[j] {
                return true
            }
        }
    }

    return false
}
