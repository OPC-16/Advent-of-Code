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
    for i := 0; i < len(line) - 13; i++ {
        chars := make([]rune, 14)
        for j := 0; j < 14; j++ {
            chars[j] = rune(line[i+j])
        }

        areCharsSame := checkSlice2(chars)   //if true then there are duplicates and in case of false there will be no duplicates
        if areCharsSame == false {
            fmt.Printf("the ans is %c on index %d\n", chars[13], i+14)
            break
        }

    }
} //function-main

func checkSlice2(chars []rune) bool {
    for i := 0; i < len(chars); i++ {
        for j := i+1; j < len(chars); j++ {
            if chars[i] == chars[j] {
                return true
            }
        }
    }

    return false
}
