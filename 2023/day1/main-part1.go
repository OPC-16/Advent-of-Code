package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
    total := 0
    file, _ := os.Open("puzzle_input.txt")
    defer file.Close()

    scn := bufio.NewScanner(file)

    for scn.Scan() {
        line := scn.Text()

        //converting string into slice of runes
        runes := []rune{}
        for len(line) > 0 {
            r, size := utf8.DecodeRuneInString(line)
            runes = append(runes, r)
            line = line[size:]
        }

        var num int
        for i := 0; i < len(runes); i++ {
            if runes[i] >= '1' && runes[i] <= '9' {
                num = int(runes[i]) - 48
                break
            }
        }
        for i := len(runes)-1; i >= 0; i-- {
            if runes[i] >= '1' && runes[i] <= '9' {
                num = num*10 + int(runes[i]) - 48
                break
            }
        }

        total += num

    }//scanner loop

    fmt.Printf("Total is: %d\n", total)
}
